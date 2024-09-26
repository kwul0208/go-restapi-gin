package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kwul0208/go-restapi-gin/handler"
	"github.com/kwul0208/go-restapi-gin/initializers"
	"github.com/kwul0208/go-restapi-gin/middleware"

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

	productRoutes := r.Group("api/products", middleware.Auth())
	{
		productRoutes.GET("/", productHandler.Index)
		productRoutes.GET("/:id", productHandler.Show)
		productRoutes.POST("/", productHandler.Create)
		productRoutes.PUT("/:id", productHandler.Update)
		productRoutes.DELETE("/", productHandler.Delete)
	}

	r.POST("api/register", authHandler.Create)
	r.POST("api/login", authHandler.Login)

	r.Run()
}
