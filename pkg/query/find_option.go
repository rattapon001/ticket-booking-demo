package query

type FindOneOptions[T any] struct {
	Where T
}

type FindManyOptions[T any] struct {
	Where   T
	Limit   int
	Offset  int
	OrderBy string
}
