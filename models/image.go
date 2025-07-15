package models

import (
	"gorm.io/gorm"
)

type Image struct {
	gorm.Model
	ProductID uint   `gorm:"index;not null"`
	URL       string `gorm:"size:255;not null"`
	AltText   string `gorm:"size:255"`
}
