package repository

import (
	"github.com/kwul0208/go-restapi-gin/models"
	"gorm.io/gorm"
)

type AuthRepository interface {
	Create(user models.Users) (models.Users, error)
	FindByEmail(Email string) (models.Users, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *authRepository {
	return &authRepository{db}
}

func (ar *authRepository) Create(user models.Users) (models.Users, error) {
	err := ar.db.Create(&user).Error

	return user, err
}

func (ar *authRepository) FindByEmail(Email string) (models.Users, error) {
	var user models.Users

	err := ar.db.Where("email = ?", Email).First(&user).Error

	return user, err
}
