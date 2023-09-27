package api

import (
	models "example/model"
	"example/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	AuthRepo *repository.AuthRepository
}

func NewAuthHandler(authRepo *repository.AuthRepository) *AuthHandler {
	return &AuthHandler{AuthRepo: authRepo}
}

func (ah *AuthHandler) SignUp(c *gin.Context) {
	var credentialsDto models.Credentials
	if err := c.BindJSON(&credentialsDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Salt and hash the password using the bcrypt algorithm
	// The second argument is the cost of hashing, which we arbitrarily set as 8 (this value can be more or less, depending on the computing power you wish to utilize)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(credentialsDto.Password), 8)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	credentialsDto.Password = string(hashedPassword)

	if err := ah.AuthRepo.SignUp(&credentialsDto); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func (ah *AuthHandler) SignIn(c *gin.Context) {
	var credentialsInDto models.Credentials
	if err := c.BindJSON(&credentialsInDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, storedPassword, err := ah.AuthRepo.FindStoredPasswordByEmail(credentialsInDto.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(credentialsInDto.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	sessionToken, refreshToken, err := ah.AuthRepo.CreateSession(userId)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	//TBD: should refresh token also have expired_at?
	c.JSON(http.StatusOK, models.SessionTokenOutDto{
		SessionToken:     sessionToken.Value,
		SessionExpiredAt: sessionToken.ExpiredAt,
		RefreshToken:     refreshToken.Value,
		RefreshExpiredAt: refreshToken.ExpiredAt,
	})
}
