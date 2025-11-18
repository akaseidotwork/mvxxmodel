package mvxxmodel

type Tag struct {
	// 以下はRoughの際に利用する(omitemptyを付与しない)
	ID   int    `gorm:"primaryKey;autoIncrement;" json:"id"`
	Name string `gorm:"unique" json:"name"`

	URL  string `gorm:"unique" json:"url,omitempty"`
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
