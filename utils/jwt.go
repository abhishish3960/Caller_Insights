package utils

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtSecret = []byte("2Y1CAlLNRlIYFBYHq+3UuVDNGK3YwdIwhCv4p3782z4=") // Replace with your secret key

// GenerateJWT generates a new JWT token for the given user ID
func GenerateJWT(userID uint) (string, error) {
	// Create a new token object, specifying signing method and claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,                                // Store user ID in the token
		"exp":     time.Now().Add(time.Hour * 72).Unix(), // Expire in 72 hours
	})

	// Sign and get the complete encoded token as a string
	return token.SignedString(jwtSecret)
}

// JWTMiddleware validates the JWT token in the Authorization header
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the token from the Authorization header
		tokenString := c.GetHeader("Authorization")

		// Token should be in the format `Bearer <token>`
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
			c.Abort()
			return
		}

		// Remove "Bearer " from the token string
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		// Parse and validate the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		// Check for token validity
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Extract claims and set user information in the context
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			userID, err := getUserIDFromClaims(claims)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
				c.Abort()
				return
			}
			c.Set("user_id", userID) // Set user_id in the request context
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		// Continue to the next handler
		c.Next()
	}
}

// Helper function to extract user_id from JWT claims
func getUserIDFromClaims(claims jwt.MapClaims) (uint, error) {
	userIDFloat, ok := claims["user_id"].(float64)
	if !ok {
		return 0, errors.New("invalid user_id in token claims")
	}
	return uint(userIDFloat), nil
}

// GetUserID extracts the user_id from the context set by the JWT middleware
func GetUserID(c *gin.Context) (uint, bool) {
	userID, exists := c.Get("user_id")
	if !exists {
		return 0, false
	}

	// Convert the userID to uint
	userIDUint, ok := userID.(uint)
	if !ok {
		return 0, false
	}

	return userIDUint, true
}
