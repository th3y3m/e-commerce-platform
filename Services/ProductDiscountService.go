package Services

import (
	"errors"
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Interface"
	"th3y3m/e-commerce-platform/Util"

	"github.com/thoas/go-funk"
)

type ProductDiscountService struct {
	productDiscountRepository Interface.IProductDiscountRepository
}

func NewProductDiscountService(productDiscountRepository Interface.IProductDiscountRepository) Interface.IProductDiscountService {
	return &ProductDiscountService{productDiscountRepository}
}

func (p *ProductDiscountService) GetPaginatedProductDiscountList(discountID string, sortBy string, productID string, pageIndex int, pageSize int) (Util.PaginatedList[BusinessObjects.ProductDiscount], error) {
	return p.productDiscountRepository.GetPaginatedProductDiscountList(discountID, sortBy, productID, pageIndex, pageSize)
}

func (p *ProductDiscountService) GetAllProductDiscounts() ([]BusinessObjects.ProductDiscount, error) {
	return p.productDiscountRepository.GetAllProductDiscounts()
}

func (p *ProductDiscountService) GetProductDiscountByID(id string) ([]BusinessObjects.ProductDiscount, error) {
	return p.productDiscountRepository.GetProductDiscountByID(id)
}

func (p *ProductDiscountService) CreateProductDiscount(productID string, discountID string) error {
	productDiscount := BusinessObjects.ProductDiscount{
		ProductID:  productID,
		DiscountID: discountID,
	}

	return p.productDiscountRepository.CreateProductDiscount(productDiscount)
}

func (p *ProductDiscountService) UpdateProductDiscount(productDiscount BusinessObjects.ProductDiscount) error {
	return p.productDiscountRepository.UpdateProductDiscount(productDiscount)
}

func (p *ProductDiscountService) DeleteProductDiscount(id string) error {
	return p.productDiscountRepository.DeleteProductDiscount(id)
}

// GetProductsOfDiscount filters products based on the provided discountID using go-funk
func (p *ProductDiscountService) GetProductsOfDiscount(discountID string) ([]BusinessObjects.ProductDiscount, error) {
	// Retrieve all ProductDiscounts
	products, err := p.GetAllProductDiscounts()
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
