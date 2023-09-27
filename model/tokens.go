package model

import (
	"time"

	"gorm.io/gorm"
)

type Token struct {
	gorm.Model
	UserID    uint      `gorm:"<-:create;not null"`
	Value     string    `gorm:"<-:create;unique;not null"`
	ExpiredAt time.Time `gorm:"<-:create;not null"`
}

type SessionToken struct {
	Token
}

type RefreshToken struct {
	Token
}

type SessionTokenOutDto struct {
	SessionToken     string
	SessionExpiredAt time.Time
	RefreshToken     string
	RefreshExpiredAt time.Time
}
