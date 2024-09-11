package Services

import (
	"log"
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Repositories"
	"th3y3m/e-commerce-platform/Util"

	"github.com/markbates/goth"
	"gorm.io/gorm"
)

func HandleGoogleUser(googleUser goth.User) (string, error) {
	// Log the Google user details for debugging purposes
	log.Printf("Handling Google user: %+v", googleUser)

	// Check if the user exists in your database by their email or Google UserID
	existingUser, err := Repositories.GetUserByEmail(googleUser.Email)
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Printf("Error retrieving user by email: %v", err)
		return "", err
	}

	// Check if the existingUser is empty
	if existingUser.UserID == "" {
		// Create a new user account using the Google user data
		newUser := BusinessObjects.User{
			UserID:   googleUser.UserID, // Storing the Google ID
			Email:    googleUser.Email,
			Username: googleUser.Name,
			// You can set additional user properties like avatar URL, etc.
		}

		// Create the new user in the database
		err := Repositories.CreateUser(newUser)
		if err != nil {
			log.Printf("Error creating new user: %v", err)
			return "", err
		}

		// Generate a JWT token for the new user
		token, err := Util.GenerateToken(newUser)
		if err != nil {
			log.Printf("Error generating token for new user: %v", err)
			return "", err
		}

		return token, nil
	} else {
		// User exists, generate a JWT token for the user
		token, err := Util.GenerateToken(existingUser)
		if err != nil {
			log.Printf("Error generating token for existing user: %v", err)
			return "", err
		}

		return token, nil
	}
}
