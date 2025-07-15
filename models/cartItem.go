package models

import (
	"gorm.io/gorm"
)

type CartItem struct {
	gorm.Model
	CartID    uint `gorm:"index;not null"`
	VariantID uint `gorm:"index;not null"`
	Quantity  int  `gorm:"not null;default:1"`

	Variant Variant `gorm:"foreignKey:VariantID"`
}
