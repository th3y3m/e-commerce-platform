package Repositories

import (
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Interface"
	"th3y3m/e-commerce-platform/Util"

	"github.com/sirupsen/logrus"
)

type DiscountRepository struct {
	log *logrus.Logger
}

func NewDiscountRepository(log *logrus.Logger) Interface.IDiscountRepository {
	return &DiscountRepository{log: log}
}

func (d *DiscountRepository) GetPaginatedDiscountList(searchValue, sortBy string, pageIndex, pageSize int, status *bool) (Util.PaginatedList[BusinessObjects.Discount], error) {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return Util.PaginatedList[BusinessObjects.Discount]{}, err
	}

	var discounts []BusinessObjects.Discount
	query := db.Model(&discounts)

	if searchValue != "" {
		query = query.Where("discount_type LIKE ?", "%"+searchValue+"%")
	}

	if status != nil {
		query = query.Where("status = ?", *status)
	}

	switch sortBy {
	case "discount_type_asc":
		query = query.Order("discount_type ASC")
	case "discount_type_desc":
		query = query.Order("discount_type DESC")
	case "discount_value_asc":
		query = query.Order("discount_value ASC")
	case "discount_value_desc":
		query = query.Order("discount_value DESC")
	case "start_date_asc":
		query = query.Order("start_date ASC")
	case "start_date_desc":
		query = query.Order("start_date DESC")
	case "end_date_asc":
		query = query.Order("end_date ASC")
	case "end_date_desc":
		query = query.Order("end_date DESC")
	default:
		query = query.Order("discount_value DESC")
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return Util.PaginatedList[BusinessObjects.Discount]{}, err
	}

	if err := query.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&discounts).Error; err != nil {
		return Util.PaginatedList[BusinessObjects.Discount]{}, err
	}

	return Util.NewPaginatedList(discounts, total, pageIndex, pageSize), nil
}

// GetAllDiscounts retrieves all discounts from the database
func (d *DiscountRepository) GetAllDiscounts() ([]BusinessObjects.Discount, error) {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return nil, err
	}

	var discounts []BusinessObjects.Discount
	if err := db.Find(&discounts).Error; err != nil {
		return nil, err
	}

	return discounts, nil
}

// GetDiscountByID retrieves a discount by its ID
func (d *DiscountRepository) GetDiscountByID(discountID string) (BusinessObjects.Discount, error) {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return BusinessObjects.Discount{}, err
	}

	var discount BusinessObjects.Discount
	if err := db.First(&discount, "discount_id = ?", discountID).Error; err != nil {
		return BusinessObjects.Discount{}, err
	}

	return discount, nil
}

// CreateDiscount adds a new discount to the database
func (d *DiscountRepository) CreateDiscount(discount BusinessObjects.Discount) error {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return err
	}

	if err := db.Create(&discount).Error; err != nil {
		return err
	}

	return nil
}

// UpdateDiscount updates an existing discount
func (d *DiscountRepository) UpdateDiscount(discount BusinessObjects.Discount) error {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return err
	}

	if err := db.Save(&discount).Error; err != nil {
		return err
	}

	return nil
}

// DeleteDiscount removes a discount from the database by its ID
func (d *DiscountRepository) DeleteDiscount(discountID string) error {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return err
	}

	if err := db.Delete(&BusinessObjects.Discount{}, "discount_id = ?", discountID).Error; err != nil {
		return err
	}

	return nil
}
