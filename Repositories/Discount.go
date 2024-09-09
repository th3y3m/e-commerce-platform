package Repositories

import (
	"log"
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Util"
)

func GetAllDiscount() ([]BusinessObjects.Discount, error) {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM Discounts")
	if err != nil {
		log.Fatalf("Error querying the database: %v", err)
		return nil, err
	}
	defer rows.Close()

	discounts := []BusinessObjects.Discount{}
	for rows.Next() {
		var discount BusinessObjects.Discount
		err := rows.Scan(&discount.DiscountID, &discount.DiscountType, &discount.DiscountValue, &discount.StartDate, &discount.EndDate)
		if err != nil {
			log.Fatalf("Error scanning row: %v", err)
			return nil, err
		}
		discounts = append(discounts, discount)
	}

	return discounts, nil
}

func GetDiscountById(id string) (BusinessObjects.Discount, error) {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return BusinessObjects.Discount{}, err
	}
	defer db.Close()

	var discount BusinessObjects.Discount

	err = db.QueryRow("SELECT * FROM Discounts WHERE DiscountID = ?", id).Scan(&discount.DiscountID, &discount.DiscountType, &discount.DiscountValue, &discount.StartDate, &discount.EndDate)
	if err != nil {
		log.Fatalf("Error querying the database: %v", err)
		return BusinessObjects.Discount{}, err
	}

	return discount, nil
}

func CreateDiscount(discount BusinessObjects.Discount) error {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return nil
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO Discounts (DiscountType, DiscountValue, StartDate, EndDate) VALUES (?, ?, ?, ?)", discount.DiscountType, discount.DiscountValue, discount.StartDate, discount.EndDate)
	if err != nil {
		log.Fatalf("Error inserting into the database: %v", err)
		return err
	}

	return nil
}

func UpdateDiscount(discount BusinessObjects.Discount) error {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return nil
	}
	defer db.Close()

	_, err = db.Exec("UPDATE Discounts SET DiscountType = ?, DiscountValue = ?, StartDate = ?, EndDate = ? WHERE DiscountID = ?", discount.DiscountType, discount.DiscountValue, discount.StartDate, discount.EndDate, discount.DiscountID)
	if err != nil {
		log.Fatalf("Error updating the database: %v", err)
		return err
	}

	return nil
}

func DeleteDiscount(discount BusinessObjects.Discount) error {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return nil
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM Discounts WHERE DiscountID = ?", discount.DiscountID)
	if err != nil {
		log.Fatalf("Error deleting from the database: %v", err)
		return err
	}

	return nil
}
