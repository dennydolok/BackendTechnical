package domains

import (
	"dennydolok/experimentals-1/models"
)

type ProductRepositoryDomain interface {
	GetProduct() []models.Product
	CreateProduct(product models.Product) error
	UpdateProduct(id int, product models.Product) error
	DeleteProduct(id int) error
}

type ProductService interface {
	GetProduct() []models.Product
	CreateProduct(product models.Product) error
	UpdateProduct(id int, product models.Product) error
	DeleteProduct(id int) error
}
