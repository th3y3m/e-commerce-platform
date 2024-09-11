package Repositories

import (
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Util"
)

func GetPaginatedReviewList(sortBy, reviewID, userID, productID string, pageIndex, pageSize int, minRating, maxRating *int, status *bool) (Util.PaginatedList[BusinessObjects.Review], error) {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
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
		return Util.PaginatedList[BusinessObjects.Review]{}, err
	}

	if err := query.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&reviews).Error; err != nil {
		return Util.PaginatedList[BusinessObjects.Review]{}, err
	}

	return Util.NewPaginatedList(reviews, totalItems, pageIndex, pageSize), nil
}

// GetAllReviews retrieves all freight reviews from the database
func GetAllReviews() ([]BusinessObjects.Review, error) {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return nil, err
	}

	var reviews []BusinessObjects.Review
	if err := db.Find(&reviews).Error; err != nil {
		return nil, err
	}

	return reviews, nil
}

// GetReviewByID retrieves a freight review by its ID
func GetReviewByID(reviewID string) (BusinessObjects.Review, error) {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return BusinessObjects.Review{}, err
	}

	var review BusinessObjects.Review
	if err := db.First(&review, "review_id = ?", reviewID).Error; err != nil {
		return BusinessObjects.Review{}, err
	}

	return review, nil
}

// CreateReview adds a new freight review to the database
func CreateReview(review BusinessObjects.Review) error {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return err
	}

	if err := db.Create(&review).Error; err != nil {
		return err
	}

	return nil
}

// UpdateReview updates an existing freight review
func UpdateReview(review BusinessObjects.Review) error {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return err
	}

	if err := db.Save(&review).Error; err != nil {
		return err
	}

	return nil
}

// DeleteReview removes a freight review from the database by its ID
func DeleteReview(reviewID string) error {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return err
	}

	if err := db.Delete(&BusinessObjects.Review{}, "review_id = ?", reviewID).Error; err != nil {
		return err
	}

	return nil
}
