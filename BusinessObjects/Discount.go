package BusinessObjects

import "time"

// Discount represents the Discounts table
type Discount struct {
	DiscountID    string    `json:"discount_id" db:"discount_id"`
	DiscountType  string    `json:"discount_type" db:"discount_type"`
	DiscountValue float64   `json:"discount_value" db:"discount_value"`
	StartDate     time.Time `json:"start_date" db:"start_date"`
	EndDate       time.Time `json:"end_date" db:"end_date"`
}
