package repository

import (
	"log"
	"rest-apis/model"
)

type ProductRepository interface {
	Save(product *model.Product) (*model.Product, error)
	FindAll() ([]model.Product, error)
}

type repo struct {}

var (
	l *log.Logger
)

func NewProductRepository(logger *log.Logger) ProductRepository {
	l = logger
	return &repo{}
}



func (*repo) Save(product *model.Product) (*model.Product, error)  {
	return product, nil
}

func (r *repo) FindAll() ([]model.Product, error)  {

	l.Println("Fetching all products from database")

	var products = []model.Product{{
		Id:   1,
		Sku:  "product sku 1",
		Name: "product name 1",
	}, {
		Id:   2,
		Sku:  "product sku 2",
		Name: "product name 2",
	}}
	return products, nil
}