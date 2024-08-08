package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint   `gorm:"uniqueIndex;not null"`
	Email    string `gorm:"size:255;not null"`
	Password string `gorm:"size:255;not null"`
}
