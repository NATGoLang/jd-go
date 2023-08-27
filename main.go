package main

import (
	"example/api"
	"example/model"
	"example/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "<user_name>:<password>@tcp(127.0.0.1:3306)/mysql?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
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
