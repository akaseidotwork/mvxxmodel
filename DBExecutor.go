package mvxxmodel

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type dbExector struct {
	db *gorm.DB
}

func NewDBExector(file string) (dbExector, error) {
	sqlDB := sqlite.Open(file)
	db, err := gorm.Open(sqlDB, &gorm.Config{})
	if err != nil {
		return dbExector{}, err
	}

	err = db.AutoMigrate(&MangaMeta{}, &Author{}, &Parody{}, &Genre{}, &Character{}, &Circle{}, &MetaToAuthor{}, &MetaToCharacter{}, &MetaToCircle{}, &MetaToGenre{}, &MetaToParody{})
	if err != nil {
		return dbExector{}, err
	}

	return dbExector{db: db}, nil
}

func (db dbExector) GetDetailFromID(id string) (MangaMeta, error) {
	tx := db.db
	var m MangaMeta
	for table := range m.GetRelations {
		tx = tx.Preload(table)
	}

	tx = tx.Where("id = ?", id).Find(&m)
	return m, tx.Error
}

func (db dbExector) GetRoughsFromQuery(q queryOfRough) ([]MangaMeta, error) {
	tx := db.db
	var result []MangaMeta
	tx = tx.Limit(q.Size).Offset(q.Size * q.Index).Find(&result)
	return result, tx.Error
}

func (db dbExector) GetTagsFromQuery(q queryOfTag) ([]Tag, error) {
	nameFinder := fmt.Sprintf("%s%s%s", "%", q.Query, "%")
	var result []Tag
	tx := db.db.Exec("SELECT id, name, works FROM % WHERE name LIKE % ASEC works OFFSET % LIMIT %", q.Table, nameFinder, q.Index*q.Size, q.Size).Scan(&result)
	return result, tx.Error
}

func (db dbExector) Insert(m MangaMeta) error {
	tx := db.db.FirstOrCreate(&m)
	if tx.Error != nil {
		return tx.Error
	}

	tx = tx.Model(&Author{})
	for a := range fromTablesToTag(m.Authors) {
		if err := tx.Where("id = ?", a).Update("works", gorm.Expr("works + 1")).Error; err != nil {
			return err
		}
	}
	tx = tx.Model(&Parody{})
	for a := range fromTablesToTag(m.Parodies) {
		if err := tx.Where("id = ?", a).Update("works", gorm.Expr("works + 1")).Error; err != nil {
			return err
		}
	}

	tx = tx.Model(&Genre{})
	for a := range fromTablesToTag(m.Genres) {
		if err := tx.Where("id = ?", a).Update("works", gorm.Expr("works + 1")).Error; err != nil {
			return err
		}
	}

	tx = tx.Model(&Character{})
	for a := range fromTablesToTag(m.Characters) {
		if err := tx.Where("id = ?", a).Update("works", gorm.Expr("works + 1")).Error; err != nil {
			return err
		}
	}
	tx = tx.Model(&Circle{})
	for a := range fromTablesToTag(m.Circles) {
		if err := tx.Where("id = ?", a).Update("works", gorm.Expr("works + 1")).Error; err != nil {
			return err
		}
	}

	return nil
}

func fromTablesToTag[T HasIDTag](tables []T) func(func(int) bool) {
	return func(yield func(int) bool) {
		for _, t := range tables {
			if !yield(t.GetID()) {
				return
			}
		}
	}
}
