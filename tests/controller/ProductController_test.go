package controller

import (
	"encoding/json"
	"github.com/abhishek70/golang-rest-api/controller"
	"github.com/abhishek70/golang-rest-api/model"
	"github.com/abhishek70/golang-rest-api/repository"
	"github.com/abhishek70/golang-rest-api/service"
	"github.com/stretchr/testify/assert"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var (
	// Creating new logger
	logger 			  = log.New(os.Stdout, "product-service-api : ", log.LstdFlags)
	productRepository = repository.NewProductRepository(logger)
	productService    = service.NewProductService(productRepository)
	productController = controller.NewProductController(logger, productService)
)

func TestGetProducts(t *testing.T) {

	// Create new HTTP Request
	req, _ := http.NewRequest("GET", "/products", nil)

	handler := http.HandlerFunc(productController.GetProducts)

	// Record the HTTP Response
	response := httptest.NewRecorder()

	// Dispatch the HTTP Request
	handler.ServeHTTP(response, req)

	// Assert HTTP status
	status := response.Code
	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Decode HTTP response
	var products []model.Product
	json.NewDecoder(io.Reader(response.Body)).Decode(&products)

	// Assert HTTP response
	assert.Equal(t, 1, products[0].Id)
	assert.Equal(t, "product sku 1", products[0].Sku)
	assert.Equal(t, "product name 1", products[0].Name)

}
