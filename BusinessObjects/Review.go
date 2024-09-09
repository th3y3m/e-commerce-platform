package BusinessObjects

import "time"

// Review represents the Reviews table
type Review struct {
	ReviewID  string    `json:"review_id" db:"review_id"`
	ProductID string    `json:"product_id" db:"product_id"`
	UserID    string    `json:"user_id" db:"user_id"`
	Rating    int       `json:"rating" db:"rating"`
	Comment   string    `json:"comment" db:"comment"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
