package repositories

import (
	"github.com/Heroin-lab/nixProject/models"
)

type UserRepositoryInterface interface {
	Create(u *models.User) error
	GetByEmail(email string) (*models.User, error)
	UpdatePassword(u *models.User)
}
type ProductsRepositoryInterface interface {
	GetByCategory(category string) ([]*models.ForSelectProducts, error)
	Insert(product *models.Products) models.Products
	Delete(product *models.Products) models.Products
	UpdatePrise(product *models.Products) models.Products
}

type SuppliersRepositoryInterface interface {
	GetByName(name string) models.Suppliers
	AddSupplier(supplier *models.Suppliers) models.Suppliers
	DeleteSupplier(name string) models.Suppliers
	Update(supplier *models.Suppliers) models.Suppliers
}

type Order interface {
	InsertProduct(order *models.Order) error
	DeleteOrderById(id string) models.Order
	ClearAllByUser(userid string) models.Order
	ChangeNumber(id string) models.Order
}

type OrderProducts interface {
	InsertProduct(order *models.Order) error
	DeleteOrderById(id string) error
	ClearAllByUser(userid string) error
	ChangeNumber(id string) error
}
