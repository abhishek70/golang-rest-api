package main

import (
	"github.com/abhishek70/golang-rest-api/controller"
	"github.com/abhishek70/golang-rest-api/database"
	"github.com/abhishek70/golang-rest-api/repository"
	router "github.com/abhishek70/golang-rest-api/router"
	"github.com/abhishek70/golang-rest-api/service"
	"log"
	"net/http"
	"os"
)

var (
	// Creating new logger
	logger 			  = log.New(os.Stdout, "product-service-api : ", log.LstdFlags)

	// Product API
	productRepository = repository.NewProductRepository(logger)
	productService    = service.NewProductService(productRepository)
	productController = controller.NewProductController(logger, productService)
	httpRouter        = router.NewMuxRouter(logger)
)


func main() {

	const PORT string = "8080"

	// Attempt to initialize and connect database
	database.Initialize()

	httpRouter.GET("/", func(response http.ResponseWriter, request *http.Request) {
		logger.Println(response, "Up and running...")
	})

	// Product API
	httpRouter.GET("/products", productController.GetProducts)
	httpRouter.SERVE(PORT)
}
