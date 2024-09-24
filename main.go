package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kwul0208/go-restapi-gin/handler"
	"github.com/kwul0208/go-restapi-gin/initializers"

	// "github.com/kwul0208/go-restapi-gin/models"
	"github.com/kwul0208/go-restapi-gin/repository"
	"github.com/kwul0208/go-restapi-gin/use_case"
)

func init() {
	// initializers.LoadEnvVariables()
}

func main() {
	r := gin.Default()

	initializers.LoadEnvVariables()
	db := initializers.ConnectDatabase()

	productRepository := repository.NewProductRepository(db)
	productUseCase := use_case.NewProductUseCase(productRepository)
	productHandler := handler.NewProductHandler(productUseCase)

	authRepository := repository.NewAuthRepository(db)
	authUseCase := use_case.NewAuthUseCase(authRepository)
	authHandler := handler.NewAuthHandler(authUseCase)

	r.GET("api/products", productHandler.Index)
	r.GET("api/products/:id", productHandler.Show)
	r.POST("api/products", productHandler.Create)
	r.PUT("api/products/:id", productHandler.Update)
	r.DELETE("api/products", productHandler.Delete)

	r.POST("api/register", authHandler.Create)

	r.Run()
}
