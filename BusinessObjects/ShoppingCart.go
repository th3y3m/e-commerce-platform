package BusinessObjects

import "time"

// ShoppingCart represents the ShoppingCart table
type ShoppingCart struct {
	CartID    string    `json:"cart_id" db:"cart_id"`
	UserID    string    `json:"user_id" db:"user_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
