package cqrs

type ProductCreateEvent struct {
	Id    int
	Name  string
	Price float64
}

type ProductNameUpdateEvent struct {
	Id   int
	Name string
}

type ProductDeleteEvent struct {
	Id int
}
type ProductEvent struct {
	Id    int
	Name  string
	Price float64
}
type QueryByProductIdEvent struct {
	Id   int
	Name string
}
type QueryByProductNameEvent struct {
	Id   int
	Name string
}
