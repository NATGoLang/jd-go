package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Credentials
	Name     string
	Birthday *time.Time `gorm:"default:null"`
}
