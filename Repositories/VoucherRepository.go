package Repositories

import (
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Util"
	"time"
)

func GetPaginatedVoucherList(sortBy, voucherID string, pageIndex, pageSize int, status *bool, startDate, endDate time.Time) (Util.PaginatedList[BusinessObjects.Voucher], error) {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return Util.PaginatedList[BusinessObjects.Voucher]{}, err
	}

	var vouchers []BusinessObjects.Voucher
	query := db.Model(&BusinessObjects.Voucher{})

	if voucherID != "" {
		query = query.Where("voucher_id = ?", voucherID)
	}

	if !startDate.IsZero() {
		query = query.Where("start_date >= ?", startDate)
	}

	if !endDate.IsZero() {
		query = query.Where("end_date <= ?", endDate)
	}

	if status != nil {
		query = query.Where("status = ?", *status)
	}

	switch sortBy {
	case "voucher_id_asc":
		query = query.Order("voucher_id ASC")
	case "voucher_id_desc":
		query = query.Order("voucher_id DESC")
	case "end_date_asc":
		query = query.Order("end_date ASC")
	case "end_date_desc":
		query = query.Order("end_date DESC")
	case "discount_amount_asc":
		query = query.Order("discount_amount ASC")
	case "discount_amount_desc":
		query = query.Order("discount_amount DESC")
	case "status_asc":
		query = query.Order("status ASC")
	case "status_desc":
		query = query.Order("status DESC")
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return Util.PaginatedList[BusinessObjects.Voucher]{}, err
	}

	if err := query.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&vouchers).Error; err != nil {
		return Util.PaginatedList[BusinessObjects.Voucher]{}, err
	}

	return Util.NewPaginatedList(vouchers, total, pageIndex, pageSize), nil
}

// GetAllVouchers retrieves all vouchers from the database
func GetAllVouchers() ([]BusinessObjects.Voucher, error) {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return nil, err
	}

	var vouchers []BusinessObjects.Voucher
	if err := db.Find(&vouchers).Error; err != nil {
		return nil, err
	}

	return vouchers, nil
}

// GetVoucherByID retrieves a voucher by its ID
func GetVoucherByID(voucherID string) (BusinessObjects.Voucher, error) {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return BusinessObjects.Voucher{}, err
	}

	var voucher BusinessObjects.Voucher
	if err := db.First(&voucher, "voucher_id = ?", voucherID).Error; err != nil {
		return BusinessObjects.Voucher{}, err
	}

	return voucher, nil
}

// CreateVoucher adds a new voucher to the database
func CreateVoucher(voucher BusinessObjects.Voucher) error {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return err
	}

	if err := db.Create(&voucher).Error; err != nil {
		return err
	}

	return nil
}

// UpdateVoucher updates an existing voucher
func UpdateVoucher(voucher BusinessObjects.Voucher) error {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return err
	}

	if err := db.Save(&voucher).Error; err != nil {
		return err
	}

	return nil
}

// DeleteVoucher removes a voucher from the database by its ID
func DeleteVoucher(voucherID string) error {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return err
	}

	if err := db.Delete(&BusinessObjects.Voucher{}, "voucher_id = ?", voucherID).Error; err != nil {
		return err
	}

	return nil
}
