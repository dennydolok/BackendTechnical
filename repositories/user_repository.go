package repositories

import (
	"dennydolok/experimentals-1/domains"
	"dennydolok/experimentals-1/models"
	"errors"

	"gorm.io/gorm"
)

type repositoryCustomer struct {
	DB *gorm.DB
}

// Get all customers from database
func (r *repositoryCustomer) GetCustomer() []models.Customer {
	customer := []models.Customer{}
	r.DB.Find(&customer)
	return customer
}

// Insert new customer data
func (r *repositoryCustomer) CreateCustomer(customer models.Customer) error {
	err := r.DB.Create(&customer).Error
	if err != nil {
		return err
	}
	return nil
}

// delete customer data by ID
func (r *repositoryCustomer) DeleteCustomer(id int) error {
	customer := models.Customer{}
	check := r.DB.Where("id = ?", id).First(&customer)
	if errors.Is(check.Error, gorm.ErrRecordNotFound) {
		err := errors.New("Record not found")
		return err
	}
	err := r.DB.Delete(&customer).Error
	if err != nil {
		return err
	}
	return nil
}

// update customer data by id
func (r *repositoryCustomer) UpdateCustomer(id int, customer models.Customer) error {
	checkModel := models.Customer{}
	check := r.DB.Where("id = ?", id).First(&checkModel)
	if errors.Is(check.Error, gorm.ErrRecordNotFound) {
		err := errors.New("Record not found")
		return err
	}
	err := r.DB.Model(&models.Customer{}).Where("id = ?", id).Omit("created_at").Updates(customer).Error
	if err != nil {
		return err
	}
	return nil
}
func NewUserRepository(db *gorm.DB) domains.CustomerRepositoryDomain {
	return &repositoryCustomer{
		DB: db,
	}
}
