package models

import (
	"gorm.io/gorm"
)

type Variant struct {
	gorm.Model
	ProductID uint   `gorm:"index;not null"`
	Color     string `gorm:"size:50;not null"`
	Size      string `gorm:"size:10;not null"`
	Price     float64
	Stock     int `gorm:"not null;default:0"`
}
