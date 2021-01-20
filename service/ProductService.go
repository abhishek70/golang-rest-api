package service

import (
	"github.com/abhishek70/golang-rest-api/model"
	"github.com/abhishek70/golang-rest-api/repository"
)

type ProductService interface {
	FindAll() ([]model.Product, error)
}

type service struct{
	repository repository.ProductRepository
}


func NewProductService(repository repository.ProductRepository) ProductService {
	return &service{repository: repository}
}

func (service *service) FindAll() ([]model.Product, error) {
	return service.repository.FindAll()
}
