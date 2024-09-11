package BusinessObjects

// ProductDiscount represents the ProductDiscounts table
type ProductDiscount struct {
	ProductID  string `gorm:"primaryKey;column:product_id"`
	DiscountID string `gorm:"primaryKey;column:discount_id"`
}
