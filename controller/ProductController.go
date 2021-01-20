package controller

import (
	"encoding/json"
	"github.com/abhishek70/golang-rest-api/service"
	"log"
	"net/http"
)

type ProductController interface {
	GetProducts(response http.ResponseWriter, request *http.Request)
}

type controller struct{
	productService service.ProductService
	logger *log.Logger
}

func NewProductController(logger *log.Logger, productService service.ProductService) ProductController {
	return &controller{productService: productService, logger: logger}
}

func (product *controller) GetProducts(response http.ResponseWriter, request *http.Request) {

	product.logger.Println("GetProducts controller method called ")

	response.Header().Set("Content-Type", "application/json")
	products, err := product.productService.FindAll()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error":"Internal server error"}`))
		//json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error getting the posts"})
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(products)
}
