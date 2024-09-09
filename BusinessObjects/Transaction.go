package BusinessObjects

import "time"

// Transaction represents the Transactions table
type Transaction struct {
	TransactionID   string    `json:"transaction_id" db:"transaction_id"`
	OrderID         string    `json:"order_id" db:"order_id"`
	PaymentAmount   float64   `json:"payment_amount" db:"payment_amount"`
	TransactionDate time.Time `json:"transaction_date" db:"transaction_date"`
	PaymentMethod   string    `json:"payment_method" db:"payment_method"`
	PaymentStatus   string    `json:"payment_status" db:"payment_status"`
}
