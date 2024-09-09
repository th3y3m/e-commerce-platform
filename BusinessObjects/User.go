package BusinessObjects

import "time"

// User represents the Users table
type User struct {
	UserID       string    `json:"user_id" db:"user_id"`
	Username     string    `json:"username" db:"username"`
	PasswordHash string    `json:"password_hash" db:"password_hash"`
	Email        string    `json:"email" db:"email"`
	FullName     string    `json:"full_name" db:"full_name"`
	PhoneNumber  string    `json:"phone_number" db:"phone_number"`
	Address      string    `json:"address" db:"address"`
	UserType     string    `json:"user_type" db:"user_type"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}
