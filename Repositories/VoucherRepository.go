package Repositories

import (
	"log"
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Util"
)

func GetAllVoucher() ([]BusinessObjects.Voucher, error) {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return nil, err
	}

	defer db.Close()

	rows, err := db.Query("SELECT * FROM Vouchers")
	if err != nil {
		log.Fatalf("Error querying the database: %v", err)
		return nil, err
	}

	defer rows.Close()

	vouchers := []BusinessObjects.Voucher{}
	for rows.Next() {
		var voucher BusinessObjects.Voucher
		err := rows.Scan(&voucher.VoucherID, &voucher.VoucherCode, &voucher.DiscountType, &voucher.DiscountValue, &voucher.MinimumOrderAmount, &voucher.MaxDiscountAmount, &voucher.StartDate, &voucher.EndDate, &voucher.UsageLimit, &voucher.UsageCount, &voucher.IsActive)
		if err != nil {
			log.Fatalf("Error scanning row: %v", err)
			return nil, err
		}
		vouchers = append(vouchers, voucher)
	}

	return vouchers, nil
}

func GetVoucherById(id string) (BusinessObjects.Voucher, error) {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return BusinessObjects.Voucher{}, err
	}
	defer db.Close()

	var voucher BusinessObjects.Voucher

	err = db.QueryRow("SELECT * FROM Vouchers WHERE VoucherID = ?", id).Scan(&voucher.VoucherID, &voucher.VoucherCode, &voucher.DiscountType, &voucher.DiscountValue, &voucher.MinimumOrderAmount, &voucher.MaxDiscountAmount, &voucher.StartDate, &voucher.EndDate, &voucher.UsageLimit, &voucher.UsageCount, &voucher.IsActive)
	if err != nil {
		log.Fatalf("Error querying the database: %v", err)
		return BusinessObjects.Voucher{}, err
	}

	return voucher, nil
}

func CreateVoucher(voucher BusinessObjects.Voucher) error {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO Vouchers (VoucherCode, DiscountType, DiscountValue, MinimumOrderAmount, MaxDiscountAmount, StartDate, EndDate, UsageLimit, UsageCount, IsActive) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", voucher.VoucherCode, voucher.DiscountType, voucher.DiscountValue, voucher.MinimumOrderAmount, voucher.MaxDiscountAmount, voucher.StartDate, voucher.EndDate, voucher.UsageLimit, voucher.UsageCount, voucher.IsActive)

	if err != nil {
		log.Fatalf("Error inserting into the database: %v", err)
		return err
	}

	return nil
}

func UpdateVoucher(voucher BusinessObjects.Voucher) error {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return err
	}
	defer db.Close()

	_, err = db.Exec("UPDATE Vouchers SET VoucherCode = ?, DiscountType = ?, DiscountValue = ?, MinimumOrderAmount = ?, MaxDiscountAmount = ?, StartDate = ?, EndDate = ?, UsageLimit = ?, UsageCount = ?, IsActive = ? WHERE VoucherID = ?", voucher.VoucherCode, voucher.DiscountType, voucher.DiscountValue, voucher.MinimumOrderAmount, voucher.MaxDiscountAmount, voucher.StartDate, voucher.EndDate, voucher.UsageLimit, voucher.UsageCount, voucher.IsActive, voucher.VoucherID)
	if err != nil {
		log.Fatalf("Error updating the database: %v", err)
		return err
	}

	return nil
}

func DeleteVoucher(id string) error {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM Vouchers WHERE VoucherID = ?", id)
	if err != nil {
		log.Fatalf("Error deleting from the database: %v", err)
		return err
	}

	return nil
}
