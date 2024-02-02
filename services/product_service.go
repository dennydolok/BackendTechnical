package services

import (
	"dennydolok/experimentals-1/config"
	"dennydolok/experimentals-1/domains"
	"dennydolok/experimentals-1/models"
)

type serviceProduct struct {
	config     config.Config
	repository domains.ProductRepositoryDomain
}

func (s *serviceProduct) GetProduct() []models.Product {
	return s.repository.GetProduct()
}

func (s *serviceProduct) CreateProduct(product models.Product) error {
	err := s.repository.CreateProduct(product)
	return err
}

func (s *serviceProduct) UpdateProduct(id int, product models.Product) error {
	err := s.repository.UpdateProduct(id, product)
	return err
}

func (s *serviceProduct) DeleteProduct(id int) error {
	err := s.repository.DeleteProduct(id)
	return err
}

func NewserviceProduct(c config.Config, repo domains.ProductRepositoryDomain) domains.ProductService {
	return &serviceProduct{
		config:     c,
		repository: repo,
	}
}
