package mvxxmodel

type Table struct {
	value string
}

func (t Table) String() string {
	return t.value
}

var (
	AuthorTable    = Table{value: "Authors"}
	CircleTable    = Table{value: "Circles"}
	ParodyTable    = Table{value: "Parodys"}
	GenreTable     = Table{value: "Genres"}
	CharacterTable = Table{value: "Characters"}

	AssumedTables = []Table{AuthorTable, CircleTable, ParodyTable, GenreTable, CharacterTable}
)
