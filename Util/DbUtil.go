package Util

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
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

func ConnectToRedis() (*redis.Client, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Create a new Redis client
	add, pass, dbStr := os.Getenv("REDIS_URI"), os.Getenv("REDIS_PASSWORD"), os.Getenv("REDIS_DB")

	db, err := strconv.Atoi(dbStr)
	if err != nil {
		return nil, err
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     add,  // Update with your Redis address
		Password: pass, // Add password if Redis requires authentication
		DB:       db,   // Use default DB
	})

	// Test Redis connection
	_, err = redisClient.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	return redisClient, nil
}
