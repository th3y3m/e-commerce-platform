package BusinessObjects

// ProductDiscount represents the ProductDiscounts table
type ProductDiscount struct {
	ProductID  string `json:"product_id" db:"product_id"`
	DiscountID string `json:"discount_id" db:"discount_id"`
}
