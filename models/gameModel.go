package models

import (
	"gorm.io/gorm"
)

type Game struct {
	gorm.Model
	GameName        string
	GameDescription string
	GameOwner       string
	InviteCode      string `gorm:"unique"`
	Public          bool
}
