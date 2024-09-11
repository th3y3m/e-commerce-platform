package BusinessObjects

// Category represents the Categories table
type Category struct {
	CategoryID   string `gorm:"primaryKey;column:category_id"`
	CategoryName string `gorm:"column:category_name"`

	Products []Product `gorm:"foreignKey:CategoryID"`
}
