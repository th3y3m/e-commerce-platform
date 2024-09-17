package Services

import (
	"errors"
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Repositories"
	"th3y3m/e-commerce-platform/Util"

	"github.com/thoas/go-funk"
)

func GetPaginatedProductDiscountList(discountID string, sortBy string, productID string, pageIndex int, pageSize int) (Util.PaginatedList[BusinessObjects.ProductDiscount], error) {
	return Repositories.GetPaginatedProductDiscountList(discountID, sortBy, productID, pageIndex, pageSize)
}

func GetAllProductDiscounts() ([]BusinessObjects.ProductDiscount, error) {
	return Repositories.GetAllProductDiscounts()
}

func GetProductDiscountByID(id string) ([]BusinessObjects.ProductDiscount, error) {
	return Repositories.GetProductDiscountByID(id)
}

func CreateProductDiscount(productID string, discountID string, discount float64) error {
	productDiscount := BusinessObjects.ProductDiscount{
		ProductID:  productID,
		DiscountID: discountID,
	}

	return Repositories.CreateProductDiscount(productDiscount)
}

func UpdateProductDiscount(productDiscount BusinessObjects.ProductDiscount) error {
	return Repositories.UpdateProductDiscount(productDiscount)
}

func DeleteProductDiscount(id string) error {
	return Repositories.DeleteProductDiscount(id)
}

// GetProductsOfDiscount filters products based on the provided discountID using go-funk
func GetProductsOfDiscount(discountID string) ([]BusinessObjects.ProductDiscount, error) {
	// Retrieve all ProductDiscounts
	products, err := GetAllProductDiscounts()
	if err != nil {
		return nil, err
	}

	// Use funk to filter the products with the matching discountID
	filteredProducts := funk.Filter(products, func(p BusinessObjects.ProductDiscount) bool {
		return p.DiscountID == discountID
	}).([]BusinessObjects.ProductDiscount)

	if len(filteredProducts) == 0 {
		return nil, errors.New("no products found for the given discountID")
	}

	return filteredProducts, nil
}
