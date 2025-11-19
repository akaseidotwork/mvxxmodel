package mvxxmodel

import "fmt"

type OrderQuery struct {
	value string
	Asc   bool
}

func (oq OrderQuery) String() string {
	direc := "asc"
	if !oq.Asc {
		direc = "desc"
	}

	return fmt.Sprintf("%s %s", oq.value, direc)
}

var (
	// for tags
	OrderWorks = OrderQuery{value: "works",
		Asc: true}

	// for works
	OrderTime = OrderQuery{value: "time",
		Asc: true}
	OrderLike = OrderQuery{value: "likes",
		Asc: true}
)

type QueryBase struct {
	Size  int `json:"size"`
	Index int `json:"index"`
}

type QueryOfRough struct {
	QueryBase
	Order OrderQuery
}

func NewQueryOfRough(qb QueryBase) QueryOfRough {
	return QueryOfRough{
		QueryBase: qb,
		Order: OrderWorks,
	}
}

type QueryOfTag struct {
	QueryBase
	Table Table
	Query string
}

func NewQueryOfTag(qb QueryBase) QueryOfTag {
	return QueryOfTag{QueryBase: qb}
}
