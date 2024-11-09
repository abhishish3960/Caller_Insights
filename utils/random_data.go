package utils

import (
	"math/rand"
	"myapp/config"
	"myapp/models"
	"time"
)

func PopulateRandomData() {
	// Create a new random source and generator
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < 100; i++ {
		user := models.User{
			Name:        randomString(8, rng),
			PhoneNumber: randomPhoneNumber(rng),
			Password:    "password",
		}
		HashPassword(&user.Password)
		config.DB.Create(&user)
	}
}

func randomString(n int, rng *rand.Rand) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rng.Intn(len(letters))]
	}
	return string(s)
}

func randomPhoneNumber(rng *rand.Rand) string {
	return "07" + randomString(8, rng)
}
