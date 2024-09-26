package Services

import (
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Interface"
	"th3y3m/e-commerce-platform/Util"
	"time"
)

type ReviewService struct {
	reviewRepository Interface.IReviewRepository
}

func NewReviewService(reviewRepository Interface.IReviewRepository) Interface.IReviewService {
	return &ReviewService{reviewRepository}
}

func (r *ReviewService) GetPaginatedReviewList(sortBy string, reviewID string, userID string, productID string, pageIndex int, pageSize int, minRating *int, maxRating *int, status *bool) (Util.PaginatedList[BusinessObjects.Review], error) {
	return r.reviewRepository.GetPaginatedReviewList(sortBy, reviewID, userID, productID, pageIndex, pageSize, minRating, maxRating, status)
}

func (r *ReviewService) GetAllReviews() ([]BusinessObjects.Review, error) {
	return r.reviewRepository.GetAllReviews()
}

func (r *ReviewService) GetReviewByID(id string) (BusinessObjects.Review, error) {
	return r.reviewRepository.GetReviewByID(id)
}

func (r *ReviewService) CreateReview(review BusinessObjects.NewReview) error {
	reviewObj := BusinessObjects.Review{
		ReviewID:  "REV" + Util.GenerateID(10),
		UserID:    review.UserID,
		ProductID: review.ProductID,
		Rating:    review.Rating,
		Comment:   review.Comment,
		Status:    true,
		CreatedAt: time.Now(),
	}

	return r.reviewRepository.CreateReview(reviewObj)
}

func (r *ReviewService) UpdateReview(reviewId, comment string, rating int) error {
	review, err := r.reviewRepository.GetReviewByID(reviewId)
	if err != nil {
		return err
	}

	review.Comment = comment
	review.Rating = rating

	return r.reviewRepository.UpdateReview(review)
}

func (r *ReviewService) DeleteReview(id string) error {
	return r.reviewRepository.DeleteReview(id)
}
