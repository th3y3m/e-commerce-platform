package BusinessObjects

// CartItem represents the CartItems table
type CartItem struct {
	CartID    string `gorm:"primaryKey;column:cart_id"`
	ProductID string `gorm:"primaryKey;column:product_id"`
	Quantity  int    `gorm:"column:quantity"`
}
