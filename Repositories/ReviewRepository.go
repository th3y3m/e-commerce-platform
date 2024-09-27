package Repositories

import (
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Interface"
	"th3y3m/e-commerce-platform/Util"

	"github.com/sirupsen/logrus"
)

type ReviewRepository struct {
	log *logrus.Logger
}

func NewReviewRepository(log *logrus.Logger) Interface.IReviewRepository {
	return &ReviewRepository{log}
}

func (r *ReviewRepository) GetPaginatedReviewList(sortBy, reviewID, userID, productID string, pageIndex, pageSize int, minRating, maxRating *int, status *bool) (Util.PaginatedList[BusinessObjects.Review], error) {
	r.log.Infof("Fetching paginated review list with sortBy: %s, reviewID: %s, userID: %s, productID: %s, pageIndex: %d, pageSize: %d, minRating: %v, maxRating: %v, status: %v", sortBy, reviewID, userID, productID, pageIndex, pageSize, minRating, maxRating, status)
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		r.log.Error("Failed to connect to PostgreSQL:", err)
		return Util.PaginatedList[BusinessObjects.Review]{}, err
	}

	var reviews []BusinessObjects.Review
	query := db.Model(&BusinessObjects.Review{})

	if reviewID != "" {
		query = query.Where("review_id = ?", reviewID)
	}

	if userID != "" {
		query = query.Where("user_id = ?", userID)
	}

	if productID != "" {
		query = query.Where("product_id = ?", productID)
	}

	if minRating != nil {
		query = query.Where("rating >= ?", *minRating)
	}

	if maxRating != nil {
		query = query.Where("rating <= ?", *maxRating)
	}

	if status != nil {
		query = query.Where("status = ?", *status)
	}

	switch sortBy {
	case "review_id_asc":
		query = query.Order("review_id ASC")
	case "review_id_desc":
		query = query.Order("review_id DESC")
	case "user_id_asc":
		query = query.Order("user_id ASC")
	case "user_id_desc":
		query = query.Order("user_id DESC")
	case "product_id_asc":
		query = query.Order("product_id ASC")
	case "product_id_desc":
		query = query.Order("product_id DESC")
	case "rating_asc":
		query = query.Order("rating ASC")
	case "rating_desc":
		query = query.Order("rating DESC")
	case "created_at_asc":
		query = query.Order("created_at ASC")
	case "created_at_desc":
		query = query.Order("created_at DESC")
	default:
		query = query.Order("created_at DESC")
	}

	var totalItems int64
	if err := query.Count(&totalItems).Error; err != nil {
		r.log.Error("Failed to count reviews:", err)
		return Util.PaginatedList[BusinessObjects.Review]{}, err
	}

	if err := query.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&reviews).Error; err != nil {
		r.log.Error("Failed to fetch paginated reviews:", err)
		return Util.PaginatedList[BusinessObjects.Review]{}, err
	}

	r.log.Infof("Successfully fetched paginated review list with total count: %d", totalItems)
	return Util.NewPaginatedList(reviews, totalItems, pageIndex, pageSize), nil
}

// GetAllReviews retrieves all freight reviews from the database
func (r *ReviewRepository) GetAllReviews() ([]BusinessObjects.Review, error) {
	r.log.Info("Fetching all reviews")
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		r.log.Error("Failed to connect to PostgreSQL:", err)
		return nil, err
	}

	var reviews []BusinessObjects.Review
	if err := db.Find(&reviews).Error; err != nil {
		r.log.Error("Failed to fetch all reviews:", err)
		return nil, err
	}

	r.log.Info("Successfully fetched all reviews")
	return reviews, nil
}

// GetReviewByID retrieves a freight review by its ID
func (r *ReviewRepository) GetReviewByID(reviewID string) (BusinessObjects.Review, error) {
	r.log.Infof("Fetching review by ID: %s", reviewID)
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		r.log.Error("Failed to connect to PostgreSQL:", err)
		return BusinessObjects.Review{}, err
	}

	var review BusinessObjects.Review
	if err := db.First(&review, "review_id = ?", reviewID).Error; err != nil {
		r.log.Error("Failed to fetch review by ID:", err)
		return BusinessObjects.Review{}, err
	}

	r.log.Infof("Successfully fetched review by ID: %s", reviewID)
	return review, nil
}

// CreateReview adds a new freight review to the database
func (r *ReviewRepository) CreateReview(review BusinessObjects.Review) error {
	r.log.Infof("Creating new review with ID: %s", review.ReviewID)
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		r.log.Error("Failed to connect to PostgreSQL:", err)
		return err
	}

	if err := db.Create(&review).Error; err != nil {
		r.log.Error("Failed to create new review:", err)
		return err
	}

	r.log.Infof("Successfully created new review with ID: %s", review.ReviewID)
	return nil
}

// UpdateReview updates an existing freight review
func (r *ReviewRepository) UpdateReview(review BusinessObjects.Review) error {
	r.log.Infof("Updating review with ID: %s", review.ReviewID)
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		r.log.Error("Failed to connect to PostgreSQL:", err)
		return err
	}

	if err := db.Save(&review).Error; err != nil {
		r.log.Error("Failed to update review:", err)
		return err
	}

	r.log.Infof("Successfully updated review with ID: %s", review.ReviewID)
	return nil
}

// DeleteReview removes a freight review from the database by its ID
func (r *ReviewRepository) DeleteReview(reviewID string) error {
	r.log.Infof("Deleting review with ID: %s", reviewID)
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		r.log.Error("Failed to connect to PostgreSQL:", err)
		return err
	}

	// if err := db.Delete(&BusinessObjects.Review{}, "review_id = ?", reviewID).Error; err != nil {
	// 	return err
	// }

	if err := db.Model(&BusinessObjects.Review{}).Where("review_id = ?", reviewID).Update("status", false).Error; err != nil {
		r.log.Error("Failed to delete review:", err)
		return err
	}

	r.log.Infof("Successfully deleted review with ID: %s", reviewID)
	return nil
}
