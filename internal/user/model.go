package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email        string `gorm:"uniqueIndex"`
	Name         string
	PassHash     string
	RefreshToken string
}
