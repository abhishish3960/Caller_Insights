package controllers

import (
	"myapp/config"
	"myapp/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SearchByName(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name query is required"})
		return
	}

	var users []models.User
	config.DB.Where("name LIKE ?", name+"%").Or("name LIKE ?", "%"+name+"%").Find(&users)

	c.JSON(http.StatusOK, gin.H{"results": users})
}

func SearchByPhoneNumber(c *gin.Context) {
	phone := c.Query("phone_number")
	if phone == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Phone number query is required"})
		return
	}

	var users []models.User
	config.DB.Where("phone_number = ?", phone).Find(&users)

	c.JSON(http.StatusOK, gin.H{"results": users})
}
