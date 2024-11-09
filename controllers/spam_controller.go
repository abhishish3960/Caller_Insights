package controllers

import (
	"myapp/config"
	"myapp/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MarkAsSpam(c *gin.Context) {
	var input struct {
		PhoneNumber string `json:"phone_number"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var spam models.Spam
	result := config.DB.Where("phone_number = ?", input.PhoneNumber).First(&spam)

	if result.RowsAffected == 0 {
		spam.PhoneNumber = input.PhoneNumber
		spam.Count = 1
		config.DB.Create(&spam)
	} else {
		spam.Count += 1
		config.DB.Save(&spam)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Number marked as spam"})
}
