package services

import (
	"dennydolok/experimentals-1/config"
	"dennydolok/experimentals-1/domains"
	"dennydolok/experimentals-1/models"
)

type serviceCustomer struct {
	config     config.Config
	repository domains.CustomerRepositoryDomain
}

func (s *serviceCustomer) GetCustomer() []models.Customer {
	return s.repository.GetCustomer()
}

func (s *serviceCustomer) CreateCustomer(customer models.Customer) error {
	return s.repository.CreateCustomer(customer)
}

func (s *serviceCustomer) DeleteCustomer(id int) error {
	err := s.repository.DeleteCustomer(id)
	return err
}

func (s *serviceCustomer) UpdateCustomer(id int, customer models.Customer) error {
	err := s.repository.UpdateCustomer(id, customer)
	return err
}

func NewserviceCustomers(c config.Config, repo domains.CustomerRepositoryDomain) domains.CustomerService {
	return &serviceCustomer{
		config:     c,
		repository: repo,
	}
}
