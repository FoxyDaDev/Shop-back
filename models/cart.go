package models

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserID    *uint
	CartItems []CartItem `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
