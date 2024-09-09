package BusinessObjects

// OrderDetail represents the OrderDetails table
type OrderDetail struct {
	OrderID   string  `json:"order_id" db:"order_id"`
	ProductID string  `json:"product_id" db:"product_id"`
	Quantity  int     `json:"quantity" db:"quantity"`
	UnitPrice float64 `json:"unit_price" db:"unit_price"`
}
