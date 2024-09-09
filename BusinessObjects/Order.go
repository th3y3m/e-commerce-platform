package BusinessObjects

import "time"

// Order represents the Orders table
type Order struct {
	OrderID               string    `json:"order_id" db:"order_id"`
	CustomerID            string    `json:"customer_id" db:"customer_id"`
	OrderDate             time.Time `json:"order_date" db:"order_date"`
	TotalAmount           float64   `json:"total_amount" db:"total_amount"`
	OrderStatus           string    `json:"order_status" db:"order_status"`
	ShippingAddress       string    `json:"shipping_address" db:"shipping_address"`
	RateID                string    `json:"rate_id" db:"rate_id"`
	EstimatedDeliveryDate time.Time `json:"estimated_delivery_date" db:"estimated_delivery_date"`
	ActualDeliveryDate    time.Time `json:"actual_delivery_date" db:"actual_delivery_date"`
	PaymentMethod         string    `json:"payment_method" db:"payment_method"`
	PaymentStatus         string    `json:"payment_status" db:"payment_status"`
	VoucherID             string    `json:"voucher_id" db:"voucher_id"`
}
