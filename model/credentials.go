package model

type Credentials struct {
	Email    string `gorm:"unique"`
	Password string `gorm:"not null"`
}
