package Repositories

import (
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Interface"
	"th3y3m/e-commerce-platform/Util"

	"github.com/sirupsen/logrus"
)

type ProductDiscountRepository struct {
	log *logrus.Logger
}

func NewProductDiscountRepository(log *logrus.Logger) Interface.IProductDiscountRepository {
	return &ProductDiscountRepository{log: log}
}

func (p *ProductDiscountRepository) GetPaginatedProductDiscountList(discountID, sortBy, productID string, pageIndex, pageSize int) (Util.PaginatedList[BusinessObjects.ProductDiscount], error) {
	p.log.Infof("Fetching paginated product discount list with discountID: %s, sortBy: %s, productID: %s, pageIndex: %d, pageSize: %d", discountID, sortBy, productID, pageIndex, pageSize)
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		p.log.Error("Failed to connect to PostgreSQL:", err)
		return Util.PaginatedList[BusinessObjects.ProductDiscount]{}, err
	}

	var rates []BusinessObjects.ProductDiscount
	query := db.Model(&BusinessObjects.ProductDiscount{})

	if productID != "" {
		query = query.Where("product_id = ?", productID)
	}
	if discountID != "" {
		query = query.Where("discount_id = ?", discountID)
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
		p.log.Error("Failed to fetch paginated product discounts:", err)
		return Util.PaginatedList[BusinessObjects.ProductDiscount]{}, err
	}

	var totalCount int64
	if err := query.Count(&totalCount).Error; err != nil {
		p.log.Error("Failed to count product discounts:", err)
		return Util.PaginatedList[BusinessObjects.ProductDiscount]{}, err
	}

	p.log.Infof("Successfully fetched paginated product discount list with total count: %d", totalCount)
	return Util.NewPaginatedList(rates, totalCount, pageIndex, pageSize), nil
}

func (p *ProductDiscountRepository) GetAllProductDiscounts() ([]BusinessObjects.ProductDiscount, error) {
	p.log.Info("Fetching all product discounts")
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		p.log.Error("Failed to connect to PostgreSQL:", err)
		return nil, err
	}

	var rates []BusinessObjects.ProductDiscount
	if err := db.Find(&rates).Error; err != nil {
		p.log.Error("Failed to fetch all product discounts:", err)
		return nil, err
	}

	p.log.Info("Successfully fetched all product discounts")
	return rates, nil
}

// GetProductDiscountByID retrieves discount rate by its ID
func (p *ProductDiscountRepository) GetProductDiscountByID(productID string) ([]BusinessObjects.ProductDiscount, error) {
	p.log.Infof("Fetching product discount by product ID: %s", productID)
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		p.log.Error("Failed to connect to PostgreSQL:", err)
		return nil, err
	}

	var rate []BusinessObjects.ProductDiscount
	if err := db.Where("product_id = ?", productID).First(&rate).Error; err != nil {
		p.log.Error("Failed to fetch product discount by product ID:", err)
		return nil, err
	}

	p.log.Infof("Successfully fetched product discount by product ID: %s", productID)
	return rate, nil
}

func (p *ProductDiscountRepository) CreateProductDiscount(rate BusinessObjects.ProductDiscount) error {
	p.log.Infof("Creating new product discount with product ID: %s", rate.ProductID)
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		p.log.Error("Failed to connect to PostgreSQL:", err)
		return err
	}

	if err := db.Create(&rate).Error; err != nil {
		p.log.Error("Failed to create new product discount:", err)
		return err
	}

	p.log.Infof("Successfully created new product discount with product ID: %s", rate.ProductID)
	return nil
}

func (p *ProductDiscountRepository) UpdateProductDiscount(rate BusinessObjects.ProductDiscount) error {
	p.log.Infof("Updating product discount with product ID: %s", rate.ProductID)
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		p.log.Error("Failed to connect to PostgreSQL:", err)
		return err
	}

	if err := db.Save(&rate).Error; err != nil {
		p.log.Error("Failed to update product discount:", err)
		return err
	}

	p.log.Infof("Successfully updated product discount with product ID: %s", rate.ProductID)
	return nil
}

func (p *ProductDiscountRepository) DeleteProductDiscount(productID string) error {
	p.log.Infof("Deleting product discount with product ID: %s", productID)
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		p.log.Error("Failed to connect to PostgreSQL:", err)
		return err
	}

	if err := db.Delete(&BusinessObjects.ProductDiscount{}, "product_id = ?", productID).Error; err != nil {
		p.log.Error("Failed to delete product discount:", err)
		return err
	}

	p.log.Infof("Successfully deleted product discount with product ID: %s", productID)
	return nil
}
