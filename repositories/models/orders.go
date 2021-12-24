package models

type Order struct {
	Id         string
	UserId     string
	ProductId  string
	SupplierId string
	Quantity   int
}
