package service

import (
	"rest-apis/model"
	"rest-apis/repository"
)

type ProductService interface {
	FindAll() ([]model.Product, error)
}

type service struct{}

var (
	repo repository.ProductRepository
)


func NewProductService(repository repository.ProductRepository) ProductService {
	repo = repository
	return &service{}
}

func (*service) FindAll() ([]model.Product, error) {
	return repo.FindAll()
}
