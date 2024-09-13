package Services

import (
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Repositories"
	"th3y3m/e-commerce-platform/Util"
)

func GetPaginatedProductDiscountList(discountID string, sortBy string, productID string, pageIndex int, pageSize int) (Util.PaginatedList[BusinessObjects.ProductDiscount], error) {
	return Repositories.GetPaginatedProductDiscountList(discountID, sortBy, productID, pageIndex, pageSize)
}

func GetAllProductDiscounts() ([]BusinessObjects.ProductDiscount, error) {
	return Repositories.GetAllProductDiscounts()
}

func GetProductDiscountByID(id string) (BusinessObjects.ProductDiscount, error) {
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
