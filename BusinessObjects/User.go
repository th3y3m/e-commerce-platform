package BusinessObjects

import "time"

// User represents a user in the system
type User struct {
	UserID       string `gorm:"primaryKey"`
	Username     string `gorm:"unique;not null"`
	PasswordHash string `gorm:"not null"`
	Email        string `gorm:"unique;not null"`
	FullName     string
	PhoneNumber  string
	Address      string
	UserType     string
	ImageURL     string
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	Token        string    `gorm:"unique"`
	TokenExpires time.Time
	Status       bool

	Orders        []Order        `gorm:"foreignKey:CustomerID"`
	News          []News         `gorm:"foreignKey:AuthorID"`
	Products      []Product      `gorm:"foreignKey:SellerID"`
	ShoppingCarts []ShoppingCart `gorm:"foreignKey:UserID"`
	Reviews       []Review       `gorm:"foreignKey:UserID"`
}
