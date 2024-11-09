package models

import (
	"gorm.io/gorm"
)

type Contact struct {
	gorm.Model
	Name    string `json:"name"`
	Phone   string `json:"phone_number"`
	OwnerID uint   `json:"owner_id"` // Foreign key to user
}
