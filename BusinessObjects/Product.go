package BusinessObjects

import "time"

// Product represents the Products table
type Product struct {
	ProductID   string    `gorm:"primaryKey;column:product_id"`
	SellerID    string    `gorm:"column:seller_id"`
	ProductName string    `gorm:"column:product_name"`
	Description string    `gorm:"column:description"`
	Price       float64   `gorm:"column:price"`
	Quantity    int       `gorm:"column:quantity"`
	CategoryID  string    `gorm:"column:category_id"`
	ImageURL    string    `gorm:"column:image_url"`
	CreatedAt   time.Time `gorm:"autoCreateTime;column:created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime;column:updated_at"`
	Status      bool      `gorm:"column:status"`

	OrderDetails     []OrderDetail     `gorm:"foreignKey:ProductID"`
	ProductDiscounts []ProductDiscount `gorm:"foreignKey:ProductID"`
	Reviews          []Review          `gorm:"foreignKey:ProductID"`
	CartItems        []CartItem        `gorm:"foreignKey:ProductID"`
}
