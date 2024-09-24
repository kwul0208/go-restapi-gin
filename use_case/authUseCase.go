package use_case

import (
	"log"

	"github.com/kwul0208/go-restapi-gin/models"
	"github.com/kwul0208/go-restapi-gin/repository"
	"github.com/kwul0208/go-restapi-gin/request"
	"golang.org/x/crypto/bcrypt"
)

type AuthUseCase interface {
	Create(request request.RegisterRequest) (models.Users, error)
}

type authUseCase struct {
	authRepository repository.AuthRepository
}

func NewAuthUseCase(authRepository repository.AuthRepository) *authUseCase {
	return &authUseCase{authRepository}
}

func (au *authUseCase) Create(registerRequest request.RegisterRequest) (models.Users, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(registerRequest.Password), 10)
	if err != nil {
		log.Panic(err.Error())
	}

	user := models.Users{
		Name:     registerRequest.Name,
		Email:    registerRequest.Email,
		Password: string(hash),
	}

	newUser, err := au.authRepository.Create(user)

	return newUser, err
}
