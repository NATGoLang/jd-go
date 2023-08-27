package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string     `gorm:"unique"`
	Birthday *time.Time `gorm:"default:null"`
}
