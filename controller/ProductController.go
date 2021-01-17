package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"rest-apis/service"
)

type ProductController interface {
	GetProducts(response http.ResponseWriter, request *http.Request)
}

type controller struct{}

var (
	productService service.ProductService
	l *log.Logger
)

func NewProductController(logger *log.Logger, service service.ProductService) ProductController {
	productService = service
	l = logger
	return &controller{}
}

func (*controller) GetProducts(response http.ResponseWriter, request *http.Request) {

	l.Println("GetProducts controller method called ")

	response.Header().Set("Content-Type", "application/json")
	products, err := productService.FindAll()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error":"Internal server error"}`))
		//json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error getting the posts"})
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(products)
}
