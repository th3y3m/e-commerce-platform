package Interface

import (
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Util"
)

type IProductDiscountRepository interface {
	GetPaginatedProductDiscountList(discountID, sortBy, productID string, pageIndex, pageSize int) (Util.PaginatedList[BusinessObjects.ProductDiscount], error)
	GetAllProductDiscounts() ([]BusinessObjects.ProductDiscount, error)
	GetProductDiscountByID(productID string) ([]BusinessObjects.ProductDiscount, error)
	CreateProductDiscount(rate BusinessObjects.ProductDiscount) error
	UpdateProductDiscount(rate BusinessObjects.ProductDiscount) error
	DeleteProductDiscount(productID string) error
}

type IProductDiscountService interface {
	GetPaginatedProductDiscountList(discountID, sortBy, productID string, pageIndex, pageSize int) (Util.PaginatedList[BusinessObjects.ProductDiscount], error)
	GetAllProductDiscounts() ([]BusinessObjects.ProductDiscount, error)
	GetProductDiscountByID(productID string) ([]BusinessObjects.ProductDiscount, error)
	CreateProductDiscount(productID string, discountID string) error
	UpdateProductDiscount(rate BusinessObjects.ProductDiscount) error
	DeleteProductDiscount(productID string) error
	GetProductsOfDiscount(discountID string) ([]BusinessObjects.ProductDiscount, error)
}
