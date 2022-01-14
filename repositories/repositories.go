package repositories

import (
	models2 "github.com/Heroin-lab/nixProject/repositories/models"
)

type UserRepositoryInterface interface {
	Create(u *models2.User) (*models2.User, error)
	GetByEmailGetByEmail(email string) (*models2.User, error)
	UpdateByEmail(u *models2.User)
}
type SuppliersRepositoryInterface interface {
	GetByName(name string) models2.Suppliers
	Insert(supplier *models2.Suppliers) models2.Suppliers
	DeleteSupplier(name string) models2.Suppliers
	Update(supplier *models2.Suppliers) models2.Suppliers
}

type ProductsRepositoryInterface interface {
	GetByCategory(category string) models2.Products
	Insert(product *models2.Products) models2.Products
	Delete(product *models2.Products) models2.Products
	UpdatePrise(product *models2.Products) models2.Products
}

type Order interface {
	InsertProduct(order *models2.Order) error
	DeleteOrderById(id string) models2.Order
	ClearAllByUser(userid string) models2.Order
	ChangeNumber(id string) models2.Order
}

type OrderProducts interface {
	InsertProduct(order *models2.Order) error
	DeleteOrderById(id string) error
	ClearAllByUser(userid string) error
	ChangeNumber(id string) error
}
