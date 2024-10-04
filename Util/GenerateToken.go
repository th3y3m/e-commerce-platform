package Util

import (
	"log"
	"os"
	"th3y3m/e-commerce-platform/BusinessObjects"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func GenerateToken(user BusinessObjects.User) (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Get environment variables
	key := os.Getenv("JWT_SECRET")

	var jwtSecret = []byte(key) // Replace with your actual secret key

	// Create the JWT claims, including user ID and expiration time
	claims := jwt.MapClaims{
		"Id":    user.UserID,
		"Role":  user.UserType,
		"Email": user.Email,
		"exp":   time.Now().Add(time.Hour * 1).Unix(), // Token expires in 1 hour
	}

	// Create the token using the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
