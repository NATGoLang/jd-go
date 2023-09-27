package repository

import (
	models "example/model"
	"time"

	"github.com/google/uuid"
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

func (ar *AuthRepository) FindStoredPasswordByEmail(email string) (uint, string, error) {
	var user models.User
	if err := ar.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return 0, "", err
	}

	return user.ID, user.Credentials.Password, nil
}

func (ar *AuthRepository) CreateSession(userId uint) (*models.SessionToken, *models.RefreshToken, error) {
	sessionToken := models.SessionToken{
		Token: models.Token{
			UserID:    userId,
			Value:     uuid.NewString(),
			ExpiredAt: time.Now().Add(24 * time.Hour)},
	}
	refreshToken := models.RefreshToken{
		Token: models.Token{
			UserID:    userId,
			Value:     uuid.NewString(),
			ExpiredAt: time.Now().Add(90 * 24 * time.Hour)},
	}
	if err := ar.DB.Transaction(func(tx *gorm.DB) error {
		if err1 := tx.Create(&sessionToken).Error; err1 != nil {
			return err1
		}

		if err2 := tx.Create(&refreshToken).Error; err2 != nil {
			return err2
		}

		return nil
	}); err != nil {
		return nil, nil, err
	}

	return &sessionToken, &refreshToken, nil
}
