package Repositories

import (
	"log"
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Util"
)

func CheckProductDiscount(productId string) ([]BusinessObjects.ProductDiscount, error) {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM ProductDiscounts WHERE ProductID = ?", productId)

	if err != nil {
		log.Fatalf("Error querying the database: %v", err)
		return nil, err
	}

	defer rows.Close()

	productDiscounts := []BusinessObjects.ProductDiscount{}
	for rows.Next() {
		var productDiscount BusinessObjects.ProductDiscount
		err := rows.Scan(&productDiscount.ProductID, &productDiscount.DiscountID)
		if err != nil {
			log.Fatalf("Error scanning row: %v", err)
			return nil, err
		}
		productDiscounts = append(productDiscounts, productDiscount)
	}

	return productDiscounts, nil
}

func GetProductByDiscountId(discountId string) (BusinessObjects.ProductDiscount, error) {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return BusinessObjects.ProductDiscount{}, err
	}
	defer db.Close()

	var productDiscount BusinessObjects.ProductDiscount

	err = db.QueryRow("SELECT * FROM ProductDiscounts WHERE DiscountID = ?", discountId).Scan(&productDiscount.ProductID, &productDiscount.DiscountID)
	if err != nil {
		log.Fatalf("Error querying the database: %v", err)
		return BusinessObjects.ProductDiscount{}, err
	}

	return productDiscount, nil
}

func CreateProductDiscount(productDiscount BusinessObjects.ProductDiscount) error {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO ProductDiscounts (ProductID, DiscountID) VALUES (?, ?)", productDiscount.ProductID, productDiscount.DiscountID)

	if err != nil {
		log.Fatalf("Error inserting into the database: %v", err)
		return err
	}

	return nil
}

func DeleteProductDiscount(productId string, discountId string) error {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return err
	}

	defer db.Close()

	_, err = db.Exec("DELETE FROM ProductDiscounts WHERE ProductID = ? AND DiscountID = ?", productId, discountId)
	if err != nil {
		log.Fatalf("Error deleting from the database: %v", err)
		return err
	}

	return nil
}
