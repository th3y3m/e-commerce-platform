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

	return Util.NewPaginatedList(rates, totalCount, pageIndex, pageSize), nil
}

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

// GetProductDiscountByID retrieves discount rate by its ID
func GetProductDiscountByID(productID string) ([]BusinessObjects.ProductDiscount, error) {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return nil, err
	}

	var rate []BusinessObjects.ProductDiscount
	if err := db.Where("product_id = ?", productID).First(&rate).Error; err != nil {
		return nil, err
	}

	return rate, nil
}

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

func DeleteProductDiscount(productID string) error {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return err
	}

	if err := db.Delete(&BusinessObjects.ProductDiscount{}, "product_id = ?", productID).Error; err != nil {
		return err
	}

	return nil
}
