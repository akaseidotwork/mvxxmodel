package mvxxmodel

type MetaToAuthor struct {
	MetaID   string `gorm:"uniqueindex:idx_meta_author"`
	AuthorID int    `gorm:"uniqueindex:idx_meta_author"`
}

type MetaToCircle struct {
	MetaID   string `gorm:"uniqueindex:idx_meta_circle"`
	CircleID int    `gorm:"uniqueindex:idx_meta_circle"`
}

type MetaToParody struct {
	MetaID   string `gorm:"uniqueindex:idx_meta_parody"`
	CircleID int    `gorm:"uniqueindex:idx_meta_parody"`
}

type MetaToGenre struct {
	MetaID  string `gorm:"uniqueindex:idx_meta_genre"`
	GenreID int    `gorm:"uniqueindex:idx_meta_genre"`
}

type MetaToCharacter struct {
	MetaID      string `gorm:"uniqueindex:idx_meta_character"`
	CharacterID int    `gorm:"uniqueindex:idx_meta_character"`
}
