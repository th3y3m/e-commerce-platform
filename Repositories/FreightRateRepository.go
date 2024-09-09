package Repositories

import (
	"log"
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Util"
)

func GetAllFreightRates() ([]BusinessObjects.FreightRate, error) {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM FreightRates")
	if err != nil {
		log.Fatalf("Error querying the database: %v", err)
		return nil, err
	}
	defer rows.Close()

	freightRates := []BusinessObjects.FreightRate{}

	for rows.Next() {
		var freightRate BusinessObjects.FreightRate
		err := rows.Scan(&freightRate.RateID, &freightRate.CourierID, &freightRate.DistanceMinKM, &freightRate.DistanceMaxKM, &freightRate.CostPerKM)
		if err != nil {
			log.Fatalf("Error scanning row: %v", err)
			return nil, err
		}
		freightRates = append(freightRates, freightRate)
	}

	return freightRates, nil
}

func GetFreightRateById(id string) (BusinessObjects.FreightRate, error) {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return BusinessObjects.FreightRate{}, err
	}
	defer db.Close()

	var freightRate BusinessObjects.FreightRate

	err = db.QueryRow("SELECT * FROM FreightRates WHERE RateID = ?", id).Scan(&freightRate.RateID, &freightRate.CourierID, &freightRate.DistanceMinKM, &freightRate.DistanceMaxKM, &freightRate.CostPerKM)
	if err != nil {
		log.Fatalf("Error querying the database: %v", err)
		return BusinessObjects.FreightRate{}, err
	}

	return freightRate, nil
}

func CreateFreightRate(freightRate BusinessObjects.FreightRate) error {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return nil
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO FreightRates (RateID, Courier, ShippingMethod, DistanceMinKM, DistanceMaxKM, CostPerKM) VALUES (?, ?, ?, ?, ?, ?)", freightRate.RateID, freightRate.CourierID, freightRate.DistanceMinKM, freightRate.DistanceMaxKM, freightRate.CostPerKM)
	if err != nil {
		log.Fatalf("Error executing query: %v", err)
		return err
	}

	return nil
}

func UpdateFreightRate(freightRate BusinessObjects.FreightRate) error {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return nil
	}
	defer db.Close()

	_, err = db.Exec("UPDATE FreightRates SET Courier = ?, ShippingMethod = ?, DistanceMinKM = ?, DistanceMaxKM = ?, CostPerKM = ? WHERE RateID = ?", freightRate.CourierID, freightRate.DistanceMinKM, freightRate.DistanceMaxKM, freightRate.CostPerKM, freightRate.RateID)
	if err != nil {
		log.Fatalf("Error executing query: %v", err)
		return err
	}

	return nil
}

func DeleteFreightRate(freightRate BusinessObjects.FreightRate) error {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return nil
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM FreightRates WHERE RateID = ?", freightRate.RateID)
	if err != nil {
		log.Fatalf("Error executing query: %v", err)
		return err
	}

	return nil
}
