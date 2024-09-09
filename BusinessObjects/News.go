package BusinessObjects

import "time"

// News represents the News table
type News struct {
	NewsID        string    `json:"news_id" db:"news_id"`
	Title         string    `json:"title" db:"title"`
	Content       string    `json:"content" db:"content"`
	PublishedDate time.Time `json:"published_date" db:"published_date"`
	AuthorID      string    `json:"author_id" db:"author_id"`
	Status        string    `json:"status" db:"status"`
	ImageURL      string    `json:"image_url" db:"image_url"`
	Category      string    `json:"category" db:"category"`
}
