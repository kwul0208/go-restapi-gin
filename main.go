package main

import (
	"github.com/gin-gonic/gin"
	productController "github.com/kwul0208/go-restapi-gin/controllers/product"
	"github.com/kwul0208/go-restapi-gin/models"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("api/products", productController.Index)
	r.GET("api/products/:id", productController.Show)
	r.POST("api/products", productController.Create)
	r.PUT("api/products/:id", productController.Update)
	r.DELETE("api/products", productController.Delete)

	r.Run()
}
