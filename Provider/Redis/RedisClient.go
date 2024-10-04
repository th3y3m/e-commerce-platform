package Redis

import (
	"context"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

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
