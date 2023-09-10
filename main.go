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

	db, err := gorm.Open(mysql.Open(mConfig.Database.Dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&model.User{})

	userRepo := repository.NewUserRepository(db)
	userHandler := api.NewUserHandler(userRepo)
	router := gin.Default()
	router.GET("/users/:email", userHandler.GetUser)
	router.POST("/users", userHandler.PostUser)

	router.Run("localhost:8080")
}
