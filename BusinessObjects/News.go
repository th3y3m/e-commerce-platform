package BusinessObjects

import "time"

// News represents the News table
type News struct {
	NewsID        string    `gorm:"primaryKey;column:news_id"`
	Title         string    `gorm:"column:title"`
	Content       string    `gorm:"column:content"`
	PublishedDate time.Time `gorm:"autoCreateTime;column:published_date"`
	AuthorID      string    `gorm:"column:author_id"`
	Status        string    `gorm:"column:status"`
	ImageURL      string    `gorm:"column:image_url"`
	Category      string    `gorm:"column:category"`
}
