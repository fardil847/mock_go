package service

import (
	"errors"
	"mock_go/entity"
	"mock_go/repository"
)

type ProductService struct {
	Repository repository.ProductRepository
}

func (service ProductService) GetOneProduct(id string) (*entity.Product, error) {
	product := service.Repository.FindById(id)
	if product == nil {
		return nil, errors.New("Product not found")

	}
	return product, nil
}

func (service ProductService) GetAllProduct() ([]*entity.Product, error) {
	products := service.Repository.FindAll()
	if len(products) == 0 {
		return nil, errors.New("No product found")
	}
	return products, nil
}
