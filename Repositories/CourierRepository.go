package Repositories

import (
	"log"
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Util"
)

func GetAllCouriers() ([]BusinessObjects.Courier, error) {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM Couriers")
	if err != nil {
		log.Fatalf("Error querying the database: %v", err)
		return nil, err
	}
	defer rows.Close()

	couriers := []BusinessObjects.Courier{}
	for rows.Next() {
		var courier BusinessObjects.Courier
		err := rows.Scan(&courier.CourierID, &courier.Courier)
		if err != nil {
			log.Fatalf("Error scanning row: %v", err)
			return nil, err
		}
		couriers = append(couriers, courier)
	}

	return couriers, nil
}

func GetCourierById(id string) (BusinessObjects.Courier, error) {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return BusinessObjects.Courier{}, err
	}
	defer db.Close()

	var courier BusinessObjects.Courier

	err = db.QueryRow("SELECT * FROM Couriers WHERE CourierID = ?", id).Scan(&courier.CourierID, &courier.Courier)
	if err != nil {
		log.Fatalf("Error querying the database: %v", err)
		return BusinessObjects.Courier{}, err
	}

	return courier, nil
}

func CreateCourier(courier BusinessObjects.Courier) error {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return nil
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO Couriers (CourierID, Courier) VALUES (?, ?)", courier.CourierID, courier.Courier)
	if err != nil {
		log.Fatalf("Error inserting into the database: %v", err)
		return err
	}

	return nil
}

func UpdateCourier(courier BusinessObjects.Courier) error {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return nil
	}
	defer db.Close()

	_, err = db.Exec("UPDATE Couriers SET Courier = ? WHERE CourierID = ?", courier.Courier, courier.CourierID)
	if err != nil {
		log.Fatalf("Error updating the database: %v", err)
		return err
	}

	return nil
}

func DeleteCourier(id string) error {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return nil
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM Couriers WHERE CourierID = ?", id)
	if err != nil {
		log.Fatalf("Error deleting from the database: %v", err)
		return err
	}

	return nil
}
