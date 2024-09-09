package BusinessObjects

import "time"

// ProductDiscount represents the ProductDiscounts table
type ProductDiscount struct {
	ProductDiscountID string    `json:"product_discount_id" db:"product_discount_id"`
	ProductID         string    `json:"product_id" db:"product_id"`
	DiscountID        string    `json:"discount_id" db:"discount_id"`
	StartDate         time.Time `json:"start_date" db:"start_date"`
	EndDate           time.Time `json:"end_date" db:"end_date"`
}
