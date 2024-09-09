package BusinessObjects

// CartItem represents the CartItems table
type CartItem struct {
	CartItemID string `json:"cart_item_id" db:"cart_item_id"`
	CartID     string `json:"cart_id" db:"cart_id"`
	ProductID  string `json:"product_id" db:"product_id"`
	Quantity   int    `json:"quantity" db:"quantity"`
}
