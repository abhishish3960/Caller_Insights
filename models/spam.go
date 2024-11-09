package models

import (
	"gorm.io/gorm"
)

type Spam struct {
	gorm.Model
	PhoneNumber string `json:"phone_number"`
	Count       int    `json:"count"`
}
