package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model

	ID       uint          `json:"id" gorm:"primaryKey"`
	UserID   uint          `json:"user_id"`
	Products []CartProduct `json:"products" gorm:"foreignKey:CartID"`
}

type CartProduct struct {
	gorm.Model

	ID        uint `gorm:"primaryKey"`
	CartID    uint
	ProductID uint
	Quantity  int
	Product   Product `gorm:"foreignKey:ProductID"`
}
