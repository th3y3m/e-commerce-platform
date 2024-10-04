package Util

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func DecodeJWT(tokenString string) (string, error) {
	// Load .env file to get the secret
	err := godotenv.Load(".env")
	if err != nil {
		return "", errors.New("error loading .env file")
	}

	// Get the JWT secret from the environment variables
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))

	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the algorithm used to sign the token
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return jwtSecret, nil
	})

	if err != nil {
		return "", errors.New("invalid token")
	}

	// Extract claims and verify them
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Extract the user ID from the claims
		userID, ok := claims["Id"].(string)
		if !ok {
			return "", errors.New("user ID not found in token")
		}
		return userID, nil
	}

	return "", errors.New("invalid token")
}
