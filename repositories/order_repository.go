package repositories

import (
	"dennydolok/experimentals-1/domains"
	"dennydolok/experimentals-1/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type repositoryOrder struct {
	DB *gorm.DB
}

func (r *repositoryOrder) CreateOrder(order models.Order) error {
	product := models.Product{}
	r.DB.Where("id = ?", order.Product_id).Find(&product)
	fmt.Println(product)
	order.Total = float64(order.Quantity) * product.Price
	err := r.DB.Create(&order).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repositoryOrder) GetOrderById(id int) models.Order {
	order := models.Order{}
	check := r.DB.Where("id = ?", id).Find(&order)
	if errors.Is(check.Error, gorm.ErrRecordNotFound) {
		return models.Order{}
	}
	return order
}

func (r *repositoryOrder) GetCustomerById(id int) models.Customer {
	customer := models.Customer{}
	check := r.DB.Where("id = ?", id).Find(&customer)
	if errors.Is(check.Error, gorm.ErrRecordNotFound) {
		return models.Customer{}
	}
	return customer
}

func (r *repositoryOrder) GetProductById(id int) models.Product {
	product := models.Product{}
	check := r.DB.Where("id = ?", id).Find(&product)
	if errors.Is(check.Error, gorm.ErrRecordNotFound) {
		return models.Product{}
	}
	return product
}

func NewOrderRepository(db *gorm.DB) domains.OrderRepositoryDomain {
	return &repositoryOrder{
		DB: db,
	}
}
