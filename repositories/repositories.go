package repositories

import "github.com/Heroin-lab/nixProject/internal/app/models"

type UserRepositoryInterface interface {
	Create(u *models.User) (*models.User, error)
	GetByEmailGetByEmail(email string) (*models.User, error)
	UpdateByEmail(u *models.User)
}
type SuppliersRepositoryInterface interface {
	GetByName(name string) models.Suppliers
	Insert(supplier *models.Suppliers) models.Suppliers
	DeleteSupplier(name string) models.Suppliers
	Update(supplier *models.Suppliers) models.Suppliers
}

type ProductsRepositoryInterface interface {
	GetByCategory(category string) models.Products
	Insert(product *models.Products) models.Products
	Delete(product *models.Products) models.Products
	UpdatePrise(product *models.Products) models.Products
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
