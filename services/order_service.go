package services

import (
	"dennydolok/experimentals-1/config"
	"dennydolok/experimentals-1/domains"
	"dennydolok/experimentals-1/models"
)

type serviceOrder struct {
	config     config.Config
	repository domains.OrderRepositoryDomain
}

func (so *serviceOrder) GetOrderById(id int) models.Order {
	order := so.repository.GetOrderById(id)
	order.Customer = so.repository.GetCustomerById(order.Customer_id)
	order.Product = so.repository.GetProductById(order.Product_id)
	return order
}

func (so *serviceOrder) CreateOrder(order models.Order) error {
	err := so.repository.CreateOrder(order)
	return err
}

func NewserviceOrder(c config.Config, repo domains.OrderRepositoryDomain) domains.OrderService {
	return &serviceOrder{
		config:     c,
		repository: repo,
	}
}
