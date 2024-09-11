package BusinessObjects

import "time"

// ShoppingCart represents the ShoppingCart table
type ShoppingCart struct {
	CartID    string    `gorm:"primaryKey;column:cart_id"`
	UserID    string    `gorm:"column:user_id"`
	CreatedAt time.Time `gorm:"autoCreateTime;column:created_at"`
	Status    bool      `gorm:"column:status"`

	CartItems []CartItem `gorm:"foreignKey:CartID"`
}
