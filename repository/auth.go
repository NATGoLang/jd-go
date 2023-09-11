package repository

import (
	models "example/model"

	"gorm.io/gorm"
)

type AuthRepository struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{DB: db}
}

func (ar *AuthRepository) SignUp(credentials *models.Credentials) error {
	user := &models.User{Credentials: *credentials}
	return ar.DB.Create(user).Error
}
