package BusinessObjects

import "time"

// Transaction represents the Transactions table
type Transaction struct {
	TransactionID   string    `gorm:"primaryKey;column:transaction_id"`
	OrderID         string    `gorm:"column:order_id"`
	PaymentAmount   float64   `gorm:"column:payment_amount"`
	TransactionDate time.Time `gorm:"autoCreateTime;column:transaction_date"`
	PaymentMethod   string    `gorm:"column:payment_method"`
	PaymentStatus   string    `gorm:"column:payment_status"`
}

type NewTransaction struct {
	OrderID       string
	PaymentAmount float64
	PaymentMethod string
	PaymentStatus string
}
