package Services

import (
	"testing"
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Util"
	"th3y3m/e-commerce-platform/mocks"

	"github.com/stretchr/testify/assert"
)

func TestGetPaginatedProductDiscountList_Success(t *testing.T) {
	productDiscountRepository := &mocks.IProductDiscountRepository{}
	productDiscountService := NewProductDiscountService(productDiscountRepository)

	productDiscountRepository.On("GetPaginatedProductDiscountList", "discountID", "sortBy", "productID", 1, 10).Return(Util.PaginatedList[BusinessObjects.ProductDiscount]{}, nil)

	_, err := productDiscountService.GetPaginatedProductDiscountList("discountID", "sortBy", "productID", 1, 10)

	assert.NoError(t, err)
	productDiscountRepository.AssertExpectations(t)
}

func TestGetAllProductDiscounts_Success(t *testing.T) {
	productDiscountRepository := &mocks.IProductDiscountRepository{}
	productDiscountService := NewProductDiscountService(productDiscountRepository)

	productDiscountRepository.On("GetAllProductDiscounts").Return([]BusinessObjects.ProductDiscount{}, nil)

	_, err := productDiscountService.GetAllProductDiscounts()

	assert.NoError(t, err)
	productDiscountRepository.AssertExpectations(t)
}

func TestGetProductDiscountByID_Success(t *testing.T) {
	productDiscountRepository := &mocks.IProductDiscountRepository{}
	productDiscountService := NewProductDiscountService(productDiscountRepository)

	productDiscountRepository.On("GetProductDiscountByID", "id").Return([]BusinessObjects.ProductDiscount{}, nil)

	_, err := productDiscountService.GetProductDiscountByID("id")

	assert.NoError(t, err)
	productDiscountRepository.AssertExpectations(t)
}

func TestCreateProductDiscount_Success(t *testing.T) {
	productDiscountRepository := &mocks.IProductDiscountRepository{}
	productDiscountService := NewProductDiscountService(productDiscountRepository)

	productDiscount := BusinessObjects.ProductDiscount{
		ProductID:  "productID",
		DiscountID: "discountID",
	}

	productDiscountRepository.On("CreateProductDiscount", productDiscount).Return(nil)

	err := productDiscountService.CreateProductDiscount("productID", "discountID")

	assert.NoError(t, err)
	productDiscountRepository.AssertExpectations(t)
}

func TestUpdateProductDiscount_Success(t *testing.T) {
	productDiscountRepository := &mocks.IProductDiscountRepository{}
	productDiscountService := NewProductDiscountService(productDiscountRepository)

	productDiscount := BusinessObjects.ProductDiscount{
		ProductID:  "productID",
		DiscountID: "discountID",
	}

	productDiscountRepository.On("UpdateProductDiscount", productDiscount).Return(nil)

	err := productDiscountService.UpdateProductDiscount(productDiscount)

	assert.NoError(t, err)
	productDiscountRepository.AssertExpectations(t)
}

func TestDeleteProductDiscount_Success(t *testing.T) {
	productDiscountRepository := &mocks.IProductDiscountRepository{}
	productDiscountService := NewProductDiscountService(productDiscountRepository)

	productDiscountRepository.On("DeleteProductDiscount", "id").Return(nil)

	err := productDiscountService.DeleteProductDiscount("id")

	assert.NoError(t, err)
	productDiscountRepository.AssertExpectations(t)
}

func TestGetProductsOfDiscount_Success(t *testing.T) {
	productDiscountRepository := &mocks.IProductDiscountRepository{}
	productDiscountService := NewProductDiscountService(productDiscountRepository)

	productDiscountRepository.On("GetAllProductDiscounts").Return([]BusinessObjects.ProductDiscount{
		{ProductID: "productID", DiscountID: "discountID"},
		{ProductID: "productID", DiscountID: "discountID2"},
		{ProductID: "productID2", DiscountID: "discountID"},
	}, nil)

	products, err := productDiscountService.GetProductsOfDiscount("discountID")

	assert.NoError(t, err)
	assert.Len(t, products, 2)
	productDiscountRepository.AssertExpectations(t)
}

func TestGetProductsOfDiscount_Error(t *testing.T) {
	productDiscountRepository := &mocks.IProductDiscountRepository{}
	productDiscountService := NewProductDiscountService(productDiscountRepository)

	productDiscountRepository.On("GetAllProductDiscounts").Return(nil, assert.AnError)

	products, err := productDiscountService.GetProductsOfDiscount("discountID")

	assert.Error(t, err)
	assert.Nil(t, products)
	productDiscountRepository.AssertExpectations(t)
}

func TestGetProductsOfDiscount_NoProducts(t *testing.T) {
	productDiscountRepository := &mocks.IProductDiscountRepository{}
	productDiscountService := NewProductDiscountService(productDiscountRepository)

	productDiscountRepository.On("GetAllProductDiscounts").Return([]BusinessObjects.ProductDiscount{}, nil)

	products, err := productDiscountService.GetProductsOfDiscount("discountID")

	assert.Error(t, err)
	assert.Nil(t, products)
	productDiscountRepository.AssertExpectations(t)
}
