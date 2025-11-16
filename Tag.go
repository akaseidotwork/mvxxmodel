package mvxxmodel

type Tag struct {
	ID   int    `gorm:"primaryKey;autoIncrement;"`
	URL  string `gorm:"unique"`
	Name string `gorm:"unique"`
}

type Author struct {
	Tag,
	Metas []MangaMeta `gorm:"many2many:meta_to_author;"`
}

type Character struct {
	Tag,
	Metas []MangaMeta `gorm:"many2many:meta_to_character;"`
}

type Circle struct {
	Tag,
	Metas []MangaMeta `gorm:"many2many:meta_to_circle;"`
}

type Genre struct {
	Tag,
	Metas []MangaMeta `gorm:"many2many:meta_to_genre;"`
}

type Parody struct {
	Tag,
	Metas []MangaMeta `gorm:"many2many:meta_to_parody;"`
}
