package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password *string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	*password = string(hashedPassword)
	return nil
}

// It also logs the process to help debug any issues.
func CheckPassword(hashedPassword, password string) error {
	log.Println("Starting password check...")

	// Log the provided hashed and plain password for debugging (avoid logging in production for security reasons)
	log.Printf("Hashed Password (from DB): %s\n", hashedPassword)
	log.Printf("Plain Password (provided): %s\n", password)

	// Compare the hashed password with the provided plain password
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		// Log the error if the comparison fails
		log.Println("Error comparing password:", err)
		return err
	}

	// Log success if the passwords match
	log.Println("Password comparison successful!")
	return nil
}
