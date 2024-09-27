package Repositories

import (
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Interface"
	"th3y3m/e-commerce-platform/Util"
	"time"

	"github.com/sirupsen/logrus"
)

type VoucherRepository struct {
	log *logrus.Logger
}

func NewVoucherRepository(log *logrus.Logger) Interface.IVoucherRepository {
	return &VoucherRepository{log}
}

func (v *VoucherRepository) GetPaginatedVoucherList(sortBy, voucherID string, pageIndex, pageSize int, status *bool, startDate, endDate time.Time) (Util.PaginatedList[BusinessObjects.Voucher], error) {
	v.log.Infof("Fetching paginated voucher list with sortBy: %s, voucherID: %s, pageIndex: %d, pageSize: %d, status: %v, startDate: %v, endDate: %v", sortBy, voucherID, pageIndex, pageSize, status, startDate, endDate)
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		v.log.Error("Failed to connect to PostgreSQL:", err)
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
		v.log.Error("Failed to count vouchers:", err)
		return Util.PaginatedList[BusinessObjects.Voucher]{}, err
	}

	if err := query.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&vouchers).Error; err != nil {
		v.log.Error("Failed to fetch paginated vouchers:", err)
		return Util.PaginatedList[BusinessObjects.Voucher]{}, err
	}

	v.log.Infof("Successfully fetched paginated voucher list with total count: %d", total)
	return Util.NewPaginatedList(vouchers, total, pageIndex, pageSize), nil
}

// GetAllVouchers retrieves all vouchers from the database
func (v *VoucherRepository) GetAllVouchers() ([]BusinessObjects.Voucher, error) {
	v.log.Info("Fetching all vouchers")
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		v.log.Error("Failed to connect to PostgreSQL:", err)
		return nil, err
	}

	var vouchers []BusinessObjects.Voucher
	if err := db.Find(&vouchers).Error; err != nil {
		v.log.Error("Failed to fetch all vouchers:", err)
		return nil, err
	}

	v.log.Info("Successfully fetched all vouchers")
	return vouchers, nil
}

// GetVoucherByID retrieves a voucher by its ID
func (v *VoucherRepository) GetVoucherByID(voucherID string) (BusinessObjects.Voucher, error) {
	v.log.Infof("Fetching voucher by ID: %s", voucherID)
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		v.log.Error("Failed to connect to PostgreSQL:", err)
		return BusinessObjects.Voucher{}, err
	}

	var voucher BusinessObjects.Voucher
	if err := db.First(&voucher, "voucher_id = ?", voucherID).Error; err != nil {
		v.log.Error("Failed to fetch voucher by ID:", err)
		return BusinessObjects.Voucher{}, err
	}

	v.log.Infof("Successfully fetched voucher by ID: %s", voucherID)
	return voucher, nil
}

// CreateVoucher adds a new voucher to the database
func (v *VoucherRepository) CreateVoucher(voucher BusinessObjects.Voucher) error {
	v.log.Infof("Creating new voucher with ID: %s", voucher.VoucherID)
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		v.log.Error("Failed to connect to PostgreSQL:", err)
		return err
	}

	if err := db.Create(&voucher).Error; err != nil {
		v.log.Error("Failed to create new voucher:", err)
		return err
	}

	v.log.Infof("Successfully created new voucher with ID: %s", voucher.VoucherID)
	return nil
}

// UpdateVoucher updates an existing voucher
func (v *VoucherRepository) UpdateVoucher(voucher BusinessObjects.Voucher) error {
	v.log.Infof("Updating voucher with ID: %s", voucher.VoucherID)
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		v.log.Error("Failed to connect to PostgreSQL:", err)
		return err
	}

	if err := db.Save(&voucher).Error; err != nil {
		v.log.Error("Failed to update voucher:", err)
		return err
	}

	v.log.Infof("Successfully updated voucher with ID: %s", voucher.VoucherID)
	return nil
}

// DeleteVoucher removes a voucher from the database by its ID
func (v *VoucherRepository) DeleteVoucher(voucherID string) error {
	v.log.Infof("Deleting voucher with ID: %s", voucherID)
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		v.log.Error("Failed to connect to PostgreSQL:", err)
		return err
	}

	// if err := db.Delete(&BusinessObjects.Voucher{}, "voucher_id = ?", voucherID).Error; err != nil {
	// 	return err
	// }

	if err := db.Model(&BusinessObjects.Voucher{}).Where("voucher_id = ?", voucherID).Update("status", false).Error; err != nil {
		v.log.Error("Failed to delete voucher:", err)
		return err
	}

	v.log.Infof("Successfully deleted voucher with ID: %s", voucherID)
	return nil
}
