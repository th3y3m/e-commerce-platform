package Interface

import (
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Util"
)

type IProductRepository interface {
	GetPaginatedProductList(searchValue, sortBy, productID, sellerID, categoryID string, pageIndex, pageSize int, status *bool) (Util.PaginatedList[BusinessObjects.Product], error)
	GetAllProducts() ([]BusinessObjects.Product, error)
	GetProductByID(productID string) (BusinessObjects.Product, error)
	CreateProduct(product BusinessObjects.Product) error
	UpdateProduct(product BusinessObjects.Product) error
	DeleteProduct(productID string) error
}

type IProductService interface {
	GetPaginatedProductList(searchValue, sortBy, productID, sellerID, categoryID string, pageIndex, pageSize int, status *bool) (Util.PaginatedList[BusinessObjects.Product], error)
	GetAllProducts() ([]BusinessObjects.Product, error)
	GetProductByID(productID string) (BusinessObjects.Product, error)
	CreateProduct(SellerID, ProductName, Description, CategoryID, ImageURL string, Price float64, Quantity int) error
	UpdateProduct(productID, SellerID, ProductName, Description, CategoryID, ImageURL string, Price float64, Quantity int) error
	UpdateAllProduct(product BusinessObjects.Product) error
	DeleteProduct(productID string) error
	GetProductPriceAfterDiscount(productID string) (float64, error)
}
