package models

type Products struct {
	Id         string
	SupplierId string
	Category   string
	Name       string
	Price      float32
	Discount   int
}
