package repository

import (
	"github.com/kwul0208/go-restapi-gin/models"
	"gorm.io/gorm"
)

type AuthRepository interface {
	Create(user models.Users) (models.Users, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *authRepository {
	return &authRepository{db}
}

func (pr *authRepository) Create(user models.Users) (models.Users, error) {
	err := pr.db.Create(&user).Error

	return user, err
}
