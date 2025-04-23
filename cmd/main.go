package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	// dbConnection
	dbConnection, err := db.ConnectDB()

	if err != nil {
		panic(err)
	}

	// repository
	ProductRepository := repository.NewProductRepository(dbConnection)

	// use case
	ProductUseCase := usecase.NewProductUseCase(ProductRepository)

	// controller
	productController := controller.NewProductController(ProductUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello, world!",
		})
	})

	server.GET("/products", productController.GetProduct)
	server.GET("/products/:id", productController.GetProductById)
	server.POST("/products", productController.CreateProduct)

	server.Run(":8000")
}
