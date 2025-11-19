package mvxxmodel

type HasIDTag interface {
	GetID() int
	GetTag() Tag
}

type Tag struct {
	// 以下はRoughの際に利用する(omitemptyを付与しない)
	ID   int    `gorm:"primaryKey;autoIncrement;" json:"id"`
	Name string `gorm:"unique" json:"name"`

	URL   string `gorm:"unique" json:"url,omitempty"`
	Works int    `json:"work"`
}

func (t Tag) GetID() int {
	return t.ID
}

func (t Tag) GetTag() Tag {
	return t
}

type Author struct {
	Tag
	Metas []MangaMeta `gorm:"many2many:meta_to_author;" json:"-"`
}

type Character struct {
	Tag
	Metas []MangaMeta `gorm:"many2many:meta_to_character;" json:"-"`
}

type Circle struct {
	Tag
	Metas []MangaMeta `gorm:"many2many:meta_to_circle;" json:"-"`
}

type Genre struct {
	Tag
	Metas []MangaMeta `gorm:"many2many:meta_to_genre;" json:"-"`
}

type Parody struct {
	Tag
	Metas []MangaMeta `gorm:"many2many:meta_to_parody;" json:"-"`
}
