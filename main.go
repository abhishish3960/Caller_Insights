package main

import (
	"log"
	"myapp/config"
	"myapp/models"
	"myapp/routes"
	"os"
)

func main() {
	// Initialize database connection
	config.ConnectDatabase()

	// Auto-migrate models
	config.AutoMigrate(&models.User{}, &models.Contact{}, &models.Spam{})

	// Populate random data for testing
	//utils.PopulateRandomData()

	// Set up the router
	r := routes.SetupRouter()

	// Start the server
	goPort := os.Getenv("GOLANG_PORT")

	if goPort == "" {
		goPort = "8080" // Default port
	}

	if err := r.Run(":" + goPort); err != nil {
		log.Fatal(err)
	}
}
