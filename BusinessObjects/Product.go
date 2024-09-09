package BusinessObjects

import "time"

// Product represents the Products table
type Product struct {
	ProductID   string    `json:"product_id" db:"product_id"`
	SellerID    string    `json:"seller_id" db:"seller_id"`
	ProductName string    `json:"product_name" db:"product_name"`
	Description string    `json:"description" db:"description"`
	Price       float64   `json:"price" db:"price"`
	Quantity    int       `json:"quantity" db:"quantity"`
	CategoryID  string    `json:"category_id" db:"category_id"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}
