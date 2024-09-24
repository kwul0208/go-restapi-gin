package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kwul0208/go-restapi-gin/request"
	useCase "github.com/kwul0208/go-restapi-gin/use_case"
)

type authHandler struct {
	authUseCase useCase.AuthUseCase
}

func NewAuthHandler(authUseCase useCase.AuthUseCase) *authHandler {
	return &authHandler{authUseCase}
}

func (h *authHandler) Create(c *gin.Context) {
	var registerRequest request.RegisterRequest

	err := c.ShouldBindJSON(&registerRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	user, err := h.authUseCase.Create(registerRequest)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "success",
		"data":    user,
	})
}
