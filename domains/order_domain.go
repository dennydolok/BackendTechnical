package domains

import "dennydolok/experimentals-1/models"

type OrderRepositoryDomain interface {
	CreateOrder(order models.Order) error
	GetOrderById(id int) models.Order
	GetProductById(id int) models.Product
	GetCustomerById(id int) models.Customer
}

type OrderService interface{
	CreateOrder(order models.Order) error
	GetOrderById(id int) models.Order
}