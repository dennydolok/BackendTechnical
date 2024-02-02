package repositories

import (
	"dennydolok/experimentals-1/domains"
	"dennydolok/experimentals-1/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type repositoryProduct struct {
	DB *gorm.DB
}

func (r *repositoryProduct) GetProduct() []models.Product {
	product := []models.Product{}
	r.DB.Find(&product)
	return product
}

func (r *repositoryProduct) CreateProduct(product models.Product) error {
	err := r.DB.Create(&product).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repositoryProduct) UpdateProduct(id int, product models.Product) error {
	checkModel := models.Product{}
	check := r.DB.Where("id = ?", id).First(&checkModel)
	if errors.Is(check.Error, gorm.ErrRecordNotFound) {
		err := errors.New("Record not found")
		return err
	}
	err := r.DB.Model(&models.Product{}).Where("id = ?", id).Omit("created_at").Updates(&product).Error
	fmt.Println(product)
	if err != nil {
		return err
	}
	return nil
}

func (r *repositoryProduct) DeleteProduct(id int) error {
	product := models.Product{}
	check := r.DB.Where("id = ?", id).First(&product)
	if errors.Is(check.Error, gorm.ErrRecordNotFound) {
		err := errors.New("Record not found")
		return err
	}
	err := r.DB.Delete(&product).Error
	if err != nil {
		return err
	}
	return nil
}

func NewProductRepository(db *gorm.DB) domains.ProductRepositoryDomain {
	return &repositoryProduct{
		DB: db,
	}
}
