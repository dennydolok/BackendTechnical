package domains

import (
	"dennydolok/experimentals-1/models"
)

type CustomerRepositoryDomain interface {
	GetCustomer() []models.Customer
	CreateCustomer(customer models.Customer) error
	UpdateCustomer(id int, customer models.Customer) error
	DeleteCustomer(id int) error
}

type CustomerService interface {
	GetCustomer() []models.Customer
	CreateCustomer(user models.Customer) error
	UpdateCustomer(id int, customer models.Customer) error
	DeleteCustomer(id int) error
}
