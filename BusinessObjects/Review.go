package BusinessObjects

import "time"

// Review represents the Reviews table
type Review struct {
	ReviewID  string    `gorm:"primaryKey;column:review_id"`
	ProductID string    `gorm:"column:product_id"`
	UserID    string    `gorm:"column:user_id"`
	Rating    int       `gorm:"column:rating"`
	Comment   string    `gorm:"column:comment"`
	CreatedAt time.Time `gorm:"autoCreateTime;column:created_at"`
	Status    bool      `gorm:"column:status"`
}
