package mvxxmodel

type HasIDTag interface {
	GetID() int
	GetTag() Tag
}

type Tag struct {
	// 以下はRoughの際に利用する(omitemptyを付与しない)
	ID    int    `gorm:"primaryKey;autoIncrement;" json:"id"`
	Name  string `gorm:"unique" json:"name"`
	Works int    `json:"work"`

	URL string `gorm:"unique" json:"url,omitempty"`
}

type ConvertibleFromTagTo[T any] interface {
	Convert(Tag) T
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

func (a Author) Convert(t Tag) Author {
	return Author{Tag: t}
}

type Character struct {
	Tag
	Metas []MangaMeta `gorm:"many2many:meta_to_character;" json:"-"`
}

func (c Character) Convert(t Tag) Character {
	return Character{Tag: t}
}

type Circle struct {
	Tag
	Metas []MangaMeta `gorm:"many2many:meta_to_circle;" json:"-"`
}

func (c Circle) Convert(t Tag) Circle {
	return Circle{Tag: t}
}

type Genre struct {
	Tag
	Metas []MangaMeta `gorm:"many2many:meta_to_genre;" json:"-"`
}

func (g Genre) Convert(t Tag) Genre {
	return Genre{Tag: t}
}

type Parody struct {
	Tag
	Metas []MangaMeta `gorm:"many2many:meta_to_parody;" json:"-"`
}

func (p Parody) Convert(t Tag) Parody {
	return Parody{Tag: t}
}
