package BusinessObjects

import "time"

// Discount represents the Discounts table
type Discount struct {
	DiscountID    string    `gorm:"primaryKey;column:discount_id"`
	DiscountType  string    `gorm:"column:discount_type"`
	DiscountValue float64   `gorm:"column:discount_value"`
	StartDate     time.Time `gorm:"column:start_date"`
	EndDate       time.Time `gorm:"column:end_date"`

	ProductDiscounts []ProductDiscount `gorm:"foreignKey:DiscountID"`
}
