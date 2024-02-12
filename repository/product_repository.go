package repository

import "mock_go/entity"

type ProductRepository interface {
	FindById(id string) *entity.Product
	FindAll() []*entity.Product
}
