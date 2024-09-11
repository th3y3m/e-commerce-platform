package BusinessObjects

import "time"

// Order represents an order in the system
type Order struct {
	OrderID               string    `gorm:"primaryKey;column:order_id"`
	CustomerID            string    `gorm:"column:customer_id"`
	OrderDate             time.Time `gorm:"autoCreateTime;column:order_date"`
	TotalAmount           float64   `gorm:"column:total_amount"`
	OrderStatus           string    `gorm:"column:order_status"`
	ShippingAddress       string    `gorm:"column:shipping_address"`
	CourierID             string    `gorm:"column:courier_id"`
	FreightPrice          float64   `gorm:"column:freight_price"`
	EstimatedDeliveryDate time.Time `gorm:"column:estimated_delivery_date"`
	ActualDeliveryDate    time.Time `gorm:"column:actual_delivery_date"`
	PaymentMethod         string    `gorm:"column:payment_method"`
	PaymentStatus         string    `gorm:"column:payment_status"`
	VoucherID             string    `gorm:"column:voucher_id"`

	OrderDetails []OrderDetail `gorm:"foreignKey:OrderID"`
	Transactions []Transaction `gorm:"foreignKey:OrderID"`
}
