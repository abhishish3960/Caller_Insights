package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name        string    `json:"name"`
	PhoneNumber string    `json:"phone_number" gorm:"unique"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Contacts    []Contact `gorm:"foreignKey:OwnerID"`
}
