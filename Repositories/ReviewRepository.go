package Repositories

import (
	"log"
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Util"
)

func GetAllReview() ([]BusinessObjects.Review, error) {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM Reviews")
	if err != nil {
		log.Fatalf("Error querying the database: %v", err)
		return nil, err
	}
	defer rows.Close()

	reviews := []BusinessObjects.Review{}
	for rows.Next() {
		var review BusinessObjects.Review
		err := rows.Scan(&review.ReviewID, &review.ProductID, &review.UserID, &review.Rating, &review.Comment, &review.CreatedAt)
		if err != nil {
			log.Fatalf("Error scanning row: %v", err)
			return nil, err
		}
		reviews = append(reviews, review)
	}

	return reviews, nil
}

func GetReviewById(id string) (BusinessObjects.Review, error) {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return BusinessObjects.Review{}, err
	}
	defer db.Close()

	var review BusinessObjects.Review

	err = db.QueryRow("SELECT * FROM Reviews WHERE ReviewID = ?", id).Scan(&review.ReviewID, &review.ProductID, &review.UserID, &review.Rating, &review.Comment, &review.CreatedAt)
	if err != nil {
		log.Fatalf("Error querying the database: %v", err)
		return BusinessObjects.Review{}, err
	}

	return review, nil
}

func CreateReview(review BusinessObjects.Review) error {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO Reviews (ProductID, UserID, Rating, Comment, CreatedAt) VALUES (?, ?, ?, ?, ?)", review.ProductID, review.UserID, review.Rating, review.Comment, review.CreatedAt)

	if err != nil {
		log.Fatalf("Error inserting into the database: %v", err)
		return err
	}

	return nil
}

func GetReviewByProductId(productId string) ([]BusinessObjects.Review, error) {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM Reviews WHERE ProductID = ?", productId)

	if err != nil {
		log.Fatalf("Error querying the database: %v", err)
		return nil, err
	}

	defer rows.Close()

	reviews := []BusinessObjects.Review{}

	for rows.Next() {
		var review BusinessObjects.Review
		err := rows.Scan(&review.ReviewID, &review.ProductID, &review.UserID, &review.Rating, &review.Comment, &review.CreatedAt)
		if err != nil {
			log.Fatalf("Error scanning row: %v", err)
			return nil, err
		}
		reviews = append(reviews, review)
	}

	return reviews, nil
}

func GetReviewByUserId(userId string) ([]BusinessObjects.Review, error) {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM Reviews WHERE UserID = ?", userId)

	if err != nil {
		log.Fatalf("Error querying the database: %v", err)
		return nil, err
	}

	defer rows.Close()

	reviews := []BusinessObjects.Review{}

	for rows.Next() {
		var review BusinessObjects.Review
		err := rows.Scan(&review.ReviewID, &review.ProductID, &review.UserID, &review.Rating, &review.Comment, &review.CreatedAt)
		if err != nil {
			log.Fatalf("Error scanning row: %v", err)
			return nil, err
		}
		reviews = append(reviews, review)
	}

	return reviews, nil
}

func UpdateReview(review BusinessObjects.Review) error {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return err
	}
	defer db.Close()

	_, err = db.Exec("UPDATE Reviews SET ProductID = ?, UserID = ?, Rating = ?, Comment = ?, CreatedAt = ? WHERE ReviewID = ?", review.ProductID, review.UserID, review.Rating, review.Comment, review.CreatedAt, review.ReviewID)

	if err != nil {
		log.Fatalf("Error updating the database: %v", err)
		return err
	}

	return nil
}

func DeleteReview(id string) error {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return err
	}

	defer db.Close()

	_, err = db.Exec("DELETE FROM Reviews WHERE ReviewID = ?", id)

	if err != nil {
		log.Fatalf("Error deleting from the database: %v", err)
		return err
	}

	return nil
}
