package mvxxmodel

type MangaMeta struct {
	ID     string `gorm:"primaryKey"`
	Title  string
	Length int
	Likes  int
	Time   int64

	Authors    []Author    `gorm:"many2many:meta_to_author"`
	Circles    []Circle    `gorm:"many2many:meta_to_author"`
	Parodies   []Parody    `gorm:"many2many:meta_to_author"`
	Genres     []Genre     `gorm:"many2many:meta_to_author"`
	Characters []Character `gorm:"many2many:meta_to_author"`
}

func (m MangaMeta) RelatedTables(yield func(any) bool) {
	for _, response := range []any{Author{}, Parody{}, Genre{}, Character{}, Circle{}} {
		if !yield(response) {
			return
		}
	}
}
