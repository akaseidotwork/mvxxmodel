package mvxxmodel

type MangaMeta struct {
	ID     string `gorm:"primaryKey"`
	Title  string
	Length int
	Likes  int
	Time   int64

	Authors    []Author    `gorm:"many2many:meta_to_author"`
	Circles    []Circle    `gorm:"many2many:meta_to_circle"`
	Parodies   []Parody    `gorm:"many2many:meta_to_parody"`
	Genres     []Genre     `gorm:"many2many:meta_to_genre"`
	Characters []Character `gorm:"many2many:meta_to_character"`
}

func (m MangaMeta) RelatedTables(yield func(any) bool) {
	for _, response := range []any{Author{}, Parody{}, Genre{}, Character{}, Circle{}} {
		if !yield(response) {
			return
		}
	}
}

func (m MangaMeta) GetRelations(yield func(string)bool) {
	for _, response := range []string{"Authors", "Circles","Parodies", "Genres", "Characters"} {
		if !yield(response) {
			return
		}
	}
}
