package Repositories

import (
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Util"
)

func GetPaginatedProductDiscountList(discountID, sortBy, productID string, pageIndex, pageSize int) (Util.PaginatedList[BusinessObjects.ProductDiscount], error) {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return Util.PaginatedList[BusinessObjects.ProductDiscount]{}, err
	}

	var rates []BusinessObjects.ProductDiscount
	query := db.Model(&BusinessObjects.ProductDiscount{})

	if productID != "" {
		query = query.Where("product_id = ?", productID)
	}
	if discountID != "" {
		query = query.Where("discount_id = ?", productID)
	}

	switch sortBy {
	case "product_id_asc":
		query = query.Order("product_id ASC")
	case "product_id_desc":
		query = query.Order("product_id DESC")
	case "discount_id_asc":
		query = query.Order("discount_id ASC")
	case "discount_id_desc":
		query = query.Order("discount_id DESC")
	default:
		query = query.Order("product_id ASC")
	}

	if err := query.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&rates).Error; err != nil {
		return Util.PaginatedList[BusinessObjects.ProductDiscount]{}, err
	}

	var totalCount int64
	if err := query.Count(&totalCount).Error; err != nil {
		return Util.PaginatedList[BusinessObjects.ProductDiscount]{}, err
	}

	return Util.PaginatedList[BusinessObjects.ProductDiscount]{Items: rates, TotalCount: totalCount, PageIndex: pageIndex, PageSize: pageSize}, nil
}

// GetAllProductDiscounts retrieves all freight rates from the database
func GetAllProductDiscounts() ([]BusinessObjects.ProductDiscount, error) {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return nil, err
	}

	var rates []BusinessObjects.ProductDiscount
	if err := db.Find(&rates).Error; err != nil {
		return nil, err
	}

	return rates, nil
}

// GetProductDiscountByID retrieves a freight rate by its ID
func GetProductDiscountByID(rateID string) (BusinessObjects.ProductDiscount, error) {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return BusinessObjects.ProductDiscount{}, err
	}

	var rate BusinessObjects.ProductDiscount
	if err := db.First(&rate, "product_id = ?", rateID).Error; err != nil {
		return BusinessObjects.ProductDiscount{}, err
	}

	return rate, nil
}

// CreateProductDiscount adds a new freight rate to the database
func CreateProductDiscount(rate BusinessObjects.ProductDiscount) error {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return err
	}

	if err := db.Create(&rate).Error; err != nil {
		return err
	}

	return nil
}

// UpdateProductDiscount updates an existing freight rate
func UpdateProductDiscount(rate BusinessObjects.ProductDiscount) error {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return err
	}

	if err := db.Save(&rate).Error; err != nil {
		return err
	}

	return nil
}

// DeleteProductDiscount removes a freight rate from the database by its ID
func DeleteProductDiscount(rateID string) error {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return err
	}

	if err := db.Delete(&BusinessObjects.ProductDiscount{}, "product_id = ?", rateID).Error; err != nil {
		return err
	}

	return nil
}
