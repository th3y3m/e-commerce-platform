package BusinessObjects

// Add a comment or any valid Go code here

// Category represents the Categories table
type Category struct {
	CategoryID   string `json:"category_id" db:"category_id"`
	CategoryName string `json:"category_name" db:"category_name"`
}
