package Util

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectToPostgreSQL connects to a PostgreSQL database using GORM
func ConnectToPostgreSQL() (*gorm.DB, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Load the connection string from the environment variable
	databaseURL := os.Getenv("CONNECTION_STRING")

	// Use GORM to open a PostgreSQL connection
	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to PostgreSQL using GORM!")
	return db, nil
}
