package Services

import (
	"errors"
	"fmt"

	"th3y3m/e-commerce-platform/Repositories"
	"th3y3m/e-commerce-platform/Util"
)

func Login(email, password string) (string, error) {
	if email == "" || password == "" {
		return "", errors.New("email and password are required")
	}

	user, err := Repositories.GetUserByEmail(email)
	if err != nil {
		return "", err
	}

	if Util.CheckPasswordHash(password, user.PasswordHash) {
		return "", errors.New("invalid password")
	}

	token, err := Util.GenerateToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}

// var asynqClient *asynq.Client

// func init() {
// 	// Load environment variables from .env file
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatalf("Error loading .env file")
// 	}

// 	redisAddr := os.Getenv("REDIS_ADDR")
// 	redisPassword := os.Getenv("REDIS_PASSWORD")
// 	redisDBStr := os.Getenv("REDIS_DB")

// 	// Convert redisDB to an integer
// 	redisDB, err := strconv.Atoi(redisDBStr)
// 	if err != nil {
// 		log.Fatalf("Invalid REDIS_DB value: %v", err)
// 	}

// 	asynqClient = asynq.NewClient(asynq.RedisClientOpt{
// 		Addr:         redisAddr,
// 		Password:     redisPassword,
// 		DB:           redisDB,
// 		DialTimeout:  10 * time.Second, // Increase dial timeout
// 		ReadTimeout:  10 * time.Second, // Increase read timeout
// 		WriteTimeout: 10 * time.Second, // Increase write timeout
// 	})
// }

func RegisterCustomer(email, password string) error {
	if email == "" || password == "" {
		return errors.New("email and password are required")
	}

	user, err := Repositories.GetUserByEmail(email)
	if err != nil {
		return fmt.Errorf("error checking user existence: %w", err)
	}

	// If the user already exists, prevent registration
	if user.Email != "" {
		return errors.New("user already exists")
	}

	hash, err := Util.HashPassword(password)
	if err != nil {
		return fmt.Errorf("error hashing password: %w", err)
	}

	newUser, err := CreateUser(email, hash, "customer")
	if err != nil {
		return fmt.Errorf("error creating user: %w", err)
	}

	token, err := Util.GenerateToken(newUser)
	if err != nil {
		return fmt.Errorf("error generating token: %w", err)
	}

	if err := Repositories.StoreToken(&newUser, token); err != nil {
		return fmt.Errorf("error storing token: %w", err)
	}

	if err := SendMail(newUser.Email, newUser.Token); err != nil {
		return fmt.Errorf("error sending mail: %w", err)
	}

	return nil
}
