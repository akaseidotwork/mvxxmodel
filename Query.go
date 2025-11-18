package mvxxmodel

type OrderQuery struct {
	value string
}

var (
	OrderWorks = OrderQuery{value: "works"}
	OrderTime  = OrderQuery{value: "time"}
)

type QueryBase struct {
	Size  int
	Index int
}

type queryOfRough struct {
	QueryBase
	Order OrderQuery
}

func NewQueryOfRough() queryOfRough {
	return queryOfRough{
		Order: OrderWorks,
	}
}

type queryOfTag struct {
	QueryBase
	Table string
	Query string
}

func NewQueryOfTag() queryOfTag {
	return queryOfTag {}
}

