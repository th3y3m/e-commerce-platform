package Services

import (
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Repositories"
	"th3y3m/e-commerce-platform/Util"
	"time"
)

func GetPaginatedReviewList(sortBy string, reviewID string, userID string, productID string, pageIndex int, pageSize int, minRating *int, maxRating *int, status *bool) (Util.PaginatedList[BusinessObjects.Review], error) {
	return Repositories.GetPaginatedReviewList(sortBy, reviewID, userID, productID, pageIndex, pageSize, minRating, maxRating, status)
}

func GetAllReviews() ([]BusinessObjects.Review, error) {
	return Repositories.GetAllReviews()
}

func GetReviewByID(id string) (BusinessObjects.Review, error) {
	return Repositories.GetReviewByID(id)
}

func CreateReview(review BusinessObjects.NewReview) error {
	reviewObj := BusinessObjects.Review{
		ReviewID:  "REV" + Util.GenerateID(10),
		UserID:    review.UserID,
		ProductID: review.ProductID,
		Rating:    review.Rating,
		Comment:   review.Comment,
		Status:    true,
		CreatedAt: time.Now(),
	}

	return Repositories.CreateReview(reviewObj)
}

func UpdateReview(reviewId, comment string, rating int) error {
	review, err := Repositories.GetReviewByID(reviewId)
	if err != nil {
		return err
	}

	review.Comment = comment
	review.Rating = rating

	return Repositories.UpdateReview(review)
}

func DeleteReview(id string) error {
	return Repositories.DeleteReview(id)
}
