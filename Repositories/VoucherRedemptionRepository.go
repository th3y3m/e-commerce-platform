package Repositories

import (
	"log"
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Util"
)

func GetAllVoucherRedemption() ([]BusinessObjects.VoucherRedemption, error) {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM VoucherRedemptions")
	if err != nil {
		log.Fatalf("Error querying the database: %v", err)
		return nil, err
	}
	defer rows.Close()

	voucherRedemptions := []BusinessObjects.VoucherRedemption{}
	for rows.Next() {
		var voucherRedemption BusinessObjects.VoucherRedemption
		err := rows.Scan(&voucherRedemption.RedemptionID, &voucherRedemption.VoucherID, &voucherRedemption.UserID, &voucherRedemption.RedeemedAt)
		if err != nil {
			log.Fatalf("Error scanning row: %v", err)
			return nil, err
		}
		voucherRedemptions = append(voucherRedemptions, voucherRedemption)
	}

	return voucherRedemptions, nil
}

func GetVoucherRedemptionById(id string) (BusinessObjects.VoucherRedemption, error) {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return BusinessObjects.VoucherRedemption{}, err
	}
	defer db.Close()

	var voucherRedemption BusinessObjects.VoucherRedemption

	err = db.QueryRow("SELECT * FROM VoucherRedemptions WHERE RedemptionID = ?", id).Scan(&voucherRedemption.RedemptionID, &voucherRedemption.VoucherID, &voucherRedemption.UserID, &voucherRedemption.RedeemedAt)
	if err != nil {
		log.Fatalf("Error querying the database: %v", err)
		return BusinessObjects.VoucherRedemption{}, err
	}

	return voucherRedemption, nil
}

func CreateVoucherRedemption(voucherRedemption BusinessObjects.VoucherRedemption) error {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO VoucherRedemptions (RedemptionID, VoucherID, UserID, RedeemedAt) VALUES (?, ?, ?, ?)", voucherRedemption.RedemptionID, voucherRedemption.VoucherID, voucherRedemption.UserID, voucherRedemption.RedeemedAt)
	if err != nil {
		log.Fatalf("Error inserting into the database: %v", err)
		return err
	}

	return nil
}

func UpdateVoucherRedemption(voucherRedemption BusinessObjects.VoucherRedemption) error {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return err
	}
	defer db.Close()

	_, err = db.Exec("UPDATE VoucherRedemptions SET VoucherID = ?, UserID = ?, RedeemedAt = ? WHERE RedemptionID = ?", voucherRedemption.VoucherID, voucherRedemption.UserID, voucherRedemption.RedeemedAt, voucherRedemption.RedemptionID)
	if err != nil {
		log.Fatalf("Error updating the database: %v", err)
		return err
	}

	return nil
}

func DeleteVoucherRedemption(id string) error {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM VoucherRedemptions WHERE RedemptionID = ?", id)
	if err != nil {
		log.Fatalf("Error deleting from the database: %v", err)
		return err
	}

	return nil
}
