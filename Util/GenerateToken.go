package Util

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("JFY6rIq/uHpNoAUQ9VSvBExfnqXosicZZIyWALbcaxs=") // Replace with your actual secret key

func GenerateToken(user string) (string, error) {
	// Create the JWT claims, including user ID and expiration time
	claims := jwt.MapClaims{
		"Id":  user,
		"exp": time.Now().Add(time.Hour * 1).Unix(), // Token expires in 1 hour
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
