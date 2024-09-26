package Interface

import (
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Util"
)

type IReviewRepository interface {
	GetPaginatedReviewList(sortBy, reviewID, userID, productID string, pageIndex, pageSize int, minRating, maxRating *int, status *bool) (Util.PaginatedList[BusinessObjects.Review], error)
	GetAllReviews() ([]BusinessObjects.Review, error)
	GetReviewByID(reviewID string) (BusinessObjects.Review, error)
	CreateReview(review BusinessObjects.Review) error
	UpdateReview(review BusinessObjects.Review) error
	DeleteReview(reviewID string) error
}
type IReviewService interface {
	GetPaginatedReviewList(sortBy, reviewID, userID, productID string, pageIndex, pageSize int, minRating, maxRating *int, status *bool) (Util.PaginatedList[BusinessObjects.Review], error)
	GetAllReviews() ([]BusinessObjects.Review, error)
	GetReviewByID(reviewID string) (BusinessObjects.Review, error)
	CreateReview(review BusinessObjects.NewReview) error
	UpdateReview(reviewId, comment string, rating int) error
	DeleteReview(reviewID string) error
}
