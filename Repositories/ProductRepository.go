package Repositories

import (
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Util"
)

func GetPaginatedProductList(searchValue, sortBy, productID, sellerID, categoryID string, pageIndex, pageSize int, status *bool) (Util.PaginatedList[BusinessObjects.Product], error) {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return Util.PaginatedList[BusinessObjects.Product]{}, err
	}

	var products []BusinessObjects.Product
	query := db.Model(&BusinessObjects.Product{})

	if productID != "" {
		query = query.Where("product_id = ?", productID)
	}

	if searchValue != "" {
		query = query.Where("product_name LIKE ?", "%"+searchValue+"%")
	}

	if status != nil {
		query = query.Where("status = ?", *status)
	}

	switch sortBy {
	case "product_id_asc":
		query = query.Order("product_id ASC")
	case "product_id_desc":
		query = query.Order("product_id DESC")
	case "product_name_asc":
		query = query.Order("product_name ASC")
	case "product_name_desc":
		query = query.Order("product_name DESC")
	case "product_description_asc":
		query = query.Order("product_description ASC")
	case "product_description_desc":
		query = query.Order("product_description DESC")
	case "product_price_asc":
		query = query.Order("product_price ASC")
	case "product_price_desc":
		query = query.Order("product_price DESC")
	case "product_stock_asc":
		query = query.Order("product_stock ASC")
	case "product_stock_desc":
		query = query.Order("product_stock DESC")
	case "status_asc":
		query = query.Order("status ASC")
	case "status_desc":
		query = query.Order("status DESC")
	}

	if err := query.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&products).Error; err != nil {
		return Util.PaginatedList[BusinessObjects.Product]{}, err
	}

	var totalCount int64
	if err := query.Count(&totalCount).Error; err != nil {
		return Util.PaginatedList[BusinessObjects.Product]{}, err
	}

	return Util.NewPaginatedList(products, totalCount, pageIndex, pageSize), nil
}

// GetAllProducts retrieves all products from the database
func GetAllProducts() ([]BusinessObjects.Product, error) {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return nil, err
	}

	var products []BusinessObjects.Product
	if err := db.Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

// GetProductByID retrieves a product by its ID
func GetProductByID(productID string) (BusinessObjects.Product, error) {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return BusinessObjects.Product{}, err
	}

	var product BusinessObjects.Product
	if err := db.First(&product, "product_id = ?", productID).Error; err != nil {
		return BusinessObjects.Product{}, err
	}

	return product, nil
}

// CreateProduct adds a new product to the database
func CreateProduct(product BusinessObjects.Product) error {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return err
	}

	if err := db.Create(&product).Error; err != nil {
		return err
	}

	return nil
}

// UpdateProduct updates an existing product
func UpdateProduct(product BusinessObjects.Product) error {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return err
	}

	if err := db.Save(&product).Error; err != nil {
		return err
	}

	return nil
}

// DeleteProduct removes a product from the database by its ID
func DeleteProduct(productID string) error {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return err
	}
	// Set product.Status to false
	if err := db.Model(&BusinessObjects.Product{}).Where("product_id = ?", productID).Update("status", false).Error; err != nil {
		return err
	}

	return nil
}
