package mvxxmodel

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DBExecutor struct {
	db *gorm.DB
}

func NewDBExector(file string) (DBExecutor, error) {
	sqlDB := sqlite.Open(file)
	db, err := gorm.Open(sqlDB, &gorm.Config{})
	if err != nil {
		return DBExecutor{}, err
	}

	m := MangaMeta{}
	err = db.AutoMigrate(append(append(m.GetIntermediateStructs(), m.GetRelatedStructs()...), &m)...)
	if err != nil {
		return DBExecutor{}, err
	}

	return DBExecutor{db: db}, nil
}

func (db DBExecutor) GetDetailFromID(id string) (MangaMeta, error) {
	tx := db.db
	var m MangaMeta
	for table := range m.GetRelations {
		tx = tx.Preload(table)
	}

	tx = tx.Where("id = ?", id).Find(&m)
	return m, tx.Error
}

func (db DBExecutor) GetRoughsFromQuery(q QueryOfRough) ([]MangaMeta, error) {
	tx := db.db
	var result []MangaMeta
	tx = tx.Order(q.Order.String()).Limit(q.Size).Offset(q.Size * q.Index).Find(&result)
	return result, tx.Error
}

func (db DBExecutor) GetTagsFromQuery(q QueryOfTag) ([]Tag, error) {
	switch q.Table {
	case AuthorTable:
		{
			return getTagsAccordingToT[Author](db.db, q)
		}
	case GenreTable:
		{
			return getTagsAccordingToT[Genre](db.db, q)
		}
	case CharacterTable:
		{
			return getTagsAccordingToT[Character](db.db, q)
		}
	case ParodyTable:
		{
			return getTagsAccordingToT[Parody](db.db, q)
		}
	case CircleTable:
		{
			return getTagsAccordingToT[Circle](db.db, q)
		}
	}

	return []Tag{}, ErrNoSuchTable

}

func (db DBExecutor) Insert(m MangaMeta) error {
	tx := db.db.Statement.FirstOrCreate(&m)
	if tx.Error != nil {
		return tx.Error
	}
	if err := afterCreate(db.db, m.Authors); err != nil {
		return err
	}
	if err := afterCreate(db.db, m.Circles); err != nil {
		return err
	}
	if err := afterCreate(db.db, m.Genres); err != nil {
		return err
	}
	if err := afterCreate(db.db, m.Characters); err != nil {
		return err
	}
	if err := afterCreate(db.db, m.Parodies); err != nil {
		return err
	}

	return nil
}

func getTagsAccordingToT[T HasIDTag](db *gorm.DB, q QueryOfTag) ([]Tag, error) {
	tx := db.Model(new(T))
	precursors := new([]T)
	tx = tx.Order("works asc").Offset(q.Index * q.Size).Limit(q.Size).Find(precursors)
	if tx.Error != nil {
		return []Tag{}, tx.Error
	}

	result := []Tag{}
	for _, precursor := range *precursors {
		result = append(result, precursor.GetTag())
	}

	return result, nil
}

func afterCreate[T HasIDTag](db *gorm.DB, list []T) error {
	tx := db.Model(new(T))
	for a := range fromTablesToTag(list) {
		err := tx.Where("id = ?", a).Update("works", gorm.Expr("works + 1")).Error
		if err != nil {
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
