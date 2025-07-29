package query

type FindOptions[T any] struct {
	Where   T
	Limit   int
	Offset  int
	OrderBy string
}
