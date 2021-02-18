package repository

import (
	"github.com/abhishek70/golang-rest-api/database"
	"github.com/abhishek70/golang-rest-api/model"
	"log"
)

type ProductRepository interface {
	//Save(product *model.Product) (*model.Product, error)
	FindAll() ([]model.Product, error)
}

type repo struct {
	logger *log.Logger
}

func NewProductRepository(logger *log.Logger) ProductRepository {
	return &repo{logger: logger}
}

//func (*repo) Save(product *model.Product) (*model.Product, error)  {
//	return product, nil
//}

func (repo *repo) FindAll() ([]model.Product, error)  {

	repo.logger.Println("Fetching all products from database")

	//var products = []model.Product{{
	//	ID:   1,
	//	Sku:  "product sku 1",
	//	Name: "product name 1",
	//}, {
	//	ID:   2,
	//	Sku:  "product sku 2",
	//	Name: "product name 2",
	//}}

	var products []model.Product

	database.DB.Find(&products)

	return products, nil
}