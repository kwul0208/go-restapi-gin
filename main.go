package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kwul0208/go-restapi-gin/handler"
	"github.com/kwul0208/go-restapi-gin/models"
	"github.com/kwul0208/go-restapi-gin/repository"
	"github.com/kwul0208/go-restapi-gin/use_case"
)

func main() {
	r := gin.Default()
	db := models.ConnectDatabase()

	productRepository := repository.NewProductRepository(db)
	productUseCase := use_case.NewProductUseCase(productRepository)
	productHandler := handler.NewProductHandler(productUseCase)

	r.GET("api/products", productHandler.Index)
	r.GET("api/products/:id", productHandler.Show)
	r.POST("api/products", productHandler.Create)
	r.PUT("api/products/:id", productHandler.Update)
	r.DELETE("api/products", productHandler.Delete)

	r.Run()
}
