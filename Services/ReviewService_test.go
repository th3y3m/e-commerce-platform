package Services

import (
	"testing"
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Util"
	"th3y3m/e-commerce-platform/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetPaginatedReviewList_Success(t *testing.T) {
	reviewRepository := &mocks.IReviewRepository{}
	reviewService := NewReviewService(reviewRepository)
	var minRating *int
	var maxRating *int
	var status *bool
	reviewRepository.On("GetPaginatedReviewList", "sortBy", "reviewID", "userID", "productID", 1, 10, minRating, maxRating, status).Return(Util.PaginatedList[BusinessObjects.Review]{}, nil)

	_, err := reviewService.GetPaginatedReviewList("sortBy", "reviewID", "userID", "productID", 1, 10, minRating, maxRating, status)

	assert.NoError(t, err)
	reviewRepository.AssertExpectations(t)
}

func TestGetAllReviews_Success(t *testing.T) {
	reviewRepository := &mocks.IReviewRepository{}
	reviewService := NewReviewService(reviewRepository)
	reviewRepository.On("GetAllReviews").Return([]BusinessObjects.Review{}, nil)

	_, err := reviewService.GetAllReviews()

	assert.NoError(t, err)
	reviewRepository.AssertExpectations(t)
}

func TestGetReviewByID_Success(t *testing.T) {
	reviewRepository := &mocks.IReviewRepository{}
	reviewService := NewReviewService(reviewRepository)
	reviewRepository.On("GetReviewByID", "id").Return(BusinessObjects.Review{}, nil)

	_, err := reviewService.GetReviewByID("id")

	assert.NoError(t, err)
	reviewRepository.AssertExpectations(t)
}

func TestCreateReview_Success(t *testing.T) {
	reviewRepository := &mocks.IReviewRepository{}
	reviewService := NewReviewService(reviewRepository)
	review := BusinessObjects.NewReview{
		UserID:    "userID",
		ProductID: "productID",
		Rating:    5,
		Comment:   "comment",
	}
	reviewRepository.On("CreateReview", mock.AnythingOfType("BusinessObjects.Review")).Return(nil)

	err := reviewService.CreateReview(review)

	assert.NoError(t, err)
	reviewRepository.AssertExpectations(t)
}

func TestUpdateReview_Success(t *testing.T) {
	reviewRepository := &mocks.IReviewRepository{}
	reviewService := NewReviewService(reviewRepository)
	review := BusinessObjects.Review{
		ReviewID:  "reviewID",
		UserID:    "userID",
		ProductID: "productID",
		Rating:    5,
		Comment:   "comment",
		Status:    true,
	}
	reviewRepository.On("GetReviewByID", "reviewID").Return(review, nil)
	reviewRepository.On("UpdateReview", review).Return(nil)

	err := reviewService.UpdateReview("reviewID", "comment", 5)

	assert.NoError(t, err)
	reviewRepository.AssertExpectations(t)
}
func TestUpdateReview_Error(t *testing.T) {
	reviewRepository := &mocks.IReviewRepository{}
	reviewService := NewReviewService(reviewRepository)
	review := BusinessObjects.Review{
		ReviewID:  "reviewID",
		UserID:    "userID",
		ProductID: "productID",
		Rating:    5,
		Comment:   "comment",
		Status:    true,
	}
	reviewRepository.On("GetReviewByID", "reviewID").Return(review, assert.AnError)

	err := reviewService.UpdateReview("reviewID", "comment", 5)

	assert.Error(t, err)
	reviewRepository.AssertExpectations(t)
}

func TestDeleteReview_Success(t *testing.T) {
	reviewRepository := &mocks.IReviewRepository{}
	reviewService := NewReviewService(reviewRepository)
	reviewRepository.On("DeleteReview", "id").Return(nil)

	err := reviewService.DeleteReview("id")

	assert.NoError(t, err)
	reviewRepository.AssertExpectations(t)
}
