package main

import (
	"example/api"
	"example/configs"
	"example/model"
	"example/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	mConfig, err := configs.InitConfig()
	if err != nil {
		panic("failed to read config")
	}

	db, err := gorm.Open(mysql.Open(mConfig.GetDSNString()), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&model.User{})

	authRepo := repository.NewAuthRepository(db)
	authHandler := api.NewAuthHandler(authRepo)

	userRepo := repository.NewUserRepository(db)
	userHandler := api.NewUserHandler(userRepo)

	router := gin.Default()
	router.POST("/signup", authHandler.SignUp)

	router.GET("/users/:email", userHandler.GetUser)
	// commented for now, should be modified to put users
	// router.POST("/users", userHandler.PostUser)

	router.Run("localhost:8080")
}
