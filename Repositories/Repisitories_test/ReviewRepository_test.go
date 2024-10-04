package Repositories_test

// import (
// 	"testing"
// 	"th3y3m/e-commerce-platform/BusinessObjects"
// 	"th3y3m/e-commerce-platform/Provider"
// 	"th3y3m/e-commerce-platform/Repositories"

// 	"github.com/sirupsen/logrus"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// func TestGetReviewByID(t *testing.T) {
// 	// Initialize the logger
// 	logger := logrus.New()

// 	// Initialize the DbProvider
// 	dbProvider := Provider.NewDbProvider(logger)

// 	// Get the mock database
// 	mockDb, err := dbProvider.GetMockDb()
// 	if err != nil {
// 		t.Fatalf("Failed to get mock database: %v", err)
// 	}

// 	// Set up expectations
// 	reviewID := "12345"
// 	expectedReview := BusinessObjects.Review{
// 		ReviewID: reviewID,
// 		// Add other fields as necessary
// 	}
// 	mockDb.On("First", mock.Anything, "review_id = ?", reviewID).Return(func(review *BusinessObjects.Review, query string, args ...interface{}) error {
// 		*review = expectedReview
// 		return nil
// 	})

// 	// Initialize the ReviewRepository
// 	repo := Repositories.NewReviewRepository(logger, mockDb)

// 	// Call the function under test
// 	result, err := repo.GetReviewByID(reviewID)

// 	// Assert the results
// 	assert.NoError(t, err)
// 	assert.Equal(t, expectedReview, result)

// 	// Assert that the expectations were met
// 	mockDb.AssertExpectations(t)
// }
