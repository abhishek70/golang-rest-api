package main

import (
	"log"
	"net/http"
	"os"
	"rest-apis/controller"
	"rest-apis/repository"
	router "rest-apis/router"
	"rest-apis/service"
)

var (
	// Creating new logger
	logger 			  = log.New(os.Stdout, "product-service-api : ", log.LstdFlags)

	productRepository = repository.NewProductRepository(logger)
	productService    = service.NewProductService(productRepository)
	productController = controller.NewProductController(logger, productService)
	httpRouter        = router.NewMuxRouter(logger)
)


func main() {

	const PORT string = "9090"

	httpRouter.GET("/", func(response http.ResponseWriter, request *http.Request) {
		logger.Println(response, "Up and running...")
	})

	httpRouter.GET("/products", productController.GetProducts)
	httpRouter.SERVE(PORT)
}
