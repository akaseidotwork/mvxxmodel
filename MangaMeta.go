package mvxxmodel

type MangaMeta struct {
	// 以下はRoughの際も必要(omitemptyを付与しない)
	ID    string `gorm:"primaryKey" json:"id"`
	Title string `json:"title"`
	Likes int    `json:"likes"`

	Length int   `json:"length,omitempty"`
	Time   int64 `json:"time,omitempty"`

	Authors    []Author    `gorm:"many2many:meta_to_author" json:"authors,omitempty"`
	Circles    []Circle    `gorm:"many2many:meta_to_circle" json:"circles,omitempty"`
	Parodies   []Parody    `gorm:"many2many:meta_to_parody" json:"parodies,omitempty"`
	Genres     []Genre     `gorm:"many2many:meta_to_genre" json:"genres,omitempty"`
	Characters []Character `gorm:"many2many:meta_to_character" json:"characters,omitempty"`
}

func (m MangaMeta) RelatedTables(yield func(any) bool) {
	for _, response := range []any{Author{}, Parody{}, Genre{}, Character{}, Circle{}} {
		if !yield(response) {
			return
		}
	}
}

func (m MangaMeta) GetRelations(yield func(string) bool) {
	for _, response := range []string{"Authors", "Circles", "Parodies", "Genres", "Characters"} {
		if !yield(response) {
			return
		}
	}
}
