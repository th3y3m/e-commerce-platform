package Repositories

import (
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Interface"
	"th3y3m/e-commerce-platform/Util"

	"github.com/sirupsen/logrus"
)

type ProductRepository struct {
	log *logrus.Logger
	db  Interface.IDatabase
}

func NewProductRepository(log *logrus.Logger, db Interface.IDatabase) Interface.IProductRepository {
	return &ProductRepository{
		log: log,
		db:  db,
	}
}

func (p *ProductRepository) GetPaginatedProductList(searchValue, sortBy, productID, sellerID, categoryID string, pageIndex, pageSize int, status *bool) (Util.PaginatedList[BusinessObjects.Product], error) {
	p.log.Infof("Fetching paginated product list with searchValue: %s, sortBy: %s, productID: %s, sellerID: %s, categoryID: %s, pageIndex: %d, pageSize: %d, status: %v", searchValue, sortBy, productID, sellerID, categoryID, pageIndex, pageSize, status)

	var products []BusinessObjects.Product
	query := p.db.Model(&BusinessObjects.Product{})

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
		p.log.Error("Failed to fetch paginated products:", err)
		return Util.PaginatedList[BusinessObjects.Product]{}, err
	}

	var totalCount int64
	if err := query.Count(&totalCount).Error; err != nil {
		p.log.Error("Failed to count products:", err)
		return Util.PaginatedList[BusinessObjects.Product]{}, err
	}

	p.log.Infof("Successfully fetched paginated product list with total count: %d", totalCount)
	return Util.NewPaginatedList(products, totalCount, pageIndex, pageSize), nil
}

// GetAllProducts retrieves all products from the database
func (p *ProductRepository) GetAllProducts() ([]BusinessObjects.Product, error) {
	p.log.Info("Fetching all products")

	var products []BusinessObjects.Product
	if err := p.db.Find(&products).Error; err != nil {
		p.log.Error("Failed to fetch all products:", err)
		return nil, err
	}

	p.log.Info("Successfully fetched all products")
	return products, nil
}

// GetProductByID retrieves a product by its ID
func (p *ProductRepository) GetProductByID(productID string) (BusinessObjects.Product, error) {
	p.log.Infof("Fetching product by ID: %s", productID)

	var product BusinessObjects.Product
	if err := p.db.First(&product, "product_id = ?", productID).Error; err != nil {
		p.log.Error("Failed to fetch product by ID:", err)
		return BusinessObjects.Product{}, err
	}

	p.log.Infof("Successfully fetched product by ID: %s", productID)
	return product, nil
}

// CreateProduct adds a new product to the database
func (p *ProductRepository) CreateProduct(product BusinessObjects.Product) error {
	p.log.Infof("Creating new product with ID: %s", product.ProductID)

	if err := p.db.Create(&product).Error; err != nil {
		p.log.Error("Failed to create new product:", err)
		return err
	}

	p.log.Infof("Successfully created new product with ID: %s", product.ProductID)
	return nil
}

// UpdateProduct updates an existing product
func (p *ProductRepository) UpdateProduct(product BusinessObjects.Product) error {
	p.log.Infof("Updating product with ID: %s", product.ProductID)

	if err := p.db.Save(&product).Error; err != nil {
		p.log.Error("Failed to update product:", err)
		return err
	}

	p.log.Infof("Successfully updated product with ID: %s", product.ProductID)
	return nil
}

// DeleteProduct removes a product from the database by its ID
func (p *ProductRepository) DeleteProduct(productID string) error {
	p.log.Infof("Deleting product with ID: %s", productID)

	// Set product.Status to false
	if err := p.db.Model(&BusinessObjects.Product{}).Where("product_id = ?", productID).Update("status", false).Error; err != nil {
		p.log.Error("Failed to delete product:", err)
		return err
	}

	p.log.Infof("Successfully deleted product with ID: %s", productID)
	return nil
}
