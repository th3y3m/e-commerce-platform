package BusinessObjects

// OrderDetail represents the OrderDetails table
type OrderDetail struct {
	OrderID   string  `gorm:"primaryKey;column:order_id"`
	ProductID string  `gorm:"primaryKey;column:product_id"`
	Quantity  int     `gorm:"column:quantity"`
	UnitPrice float64 `gorm:"column:unit_price"`
}
