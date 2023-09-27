package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
	Role     string `gorm:"default:USER"`
	Avatar   string `gorm:"default:avatar.png"`
}
