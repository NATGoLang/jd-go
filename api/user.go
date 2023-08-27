package api

import (
	models "example/model"
	"example/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserRepo *repository.UserRepository
}

func NewUserHandler(userRepo *repository.UserRepository) *UserHandler {
	return &UserHandler{UserRepo: userRepo}
}

func (uh *UserHandler) GetUser(c *gin.Context) {
	email := c.Param("email")

	user, err := uh.UserRepo.FindByEmail(email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (uh *UserHandler) PostUser(c *gin.Context) {
	var postDto models.User

	if err := c.BindJSON(&postDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := uh.UserRepo.Create(&postDto); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, postDto)
}
