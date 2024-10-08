package Services

import (
	"testing"
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Util"
	"th3y3m/e-commerce-platform/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetPaginatedProductList_Success(t *testing.T) {
	productRepository := &mocks.IProductRepository{}
	discountService := &mocks.IDiscountService{}
	productDiscountService := &mocks.IProductDiscountService{}
	productService := NewProductService(productRepository, discountService, productDiscountService)

	var status *bool

	productRepository.On("GetPaginatedProductList", "searchValue", "sortBy", "productID", "sellerID", "categoryID", 1, 10, status).Return(Util.PaginatedList[BusinessObjects.Product]{}, nil)

	_, err := productService.GetPaginatedProductList("searchValue", "sortBy", "productID", "sellerID", "categoryID", 1, 10, status)

	assert.NoError(t, err)
	productRepository.AssertExpectations(t)
}

func TestGetAllProducts_Success(t *testing.T) {
	productRepository := &mocks.IProductRepository{}
	discountService := &mocks.IDiscountService{}
	productDiscountService := &mocks.IProductDiscountService{}
	productService := NewProductService(productRepository, discountService, productDiscountService)

	productRepository.On("GetAllProducts").Return([]BusinessObjects.Product{}, nil)

	_, err := productService.GetAllProducts()

	assert.NoError(t, err)
	productRepository.AssertExpectations(t)
}

func TestGetProductByID_Success(t *testing.T) {
	productRepository := &mocks.IProductRepository{}
	discountService := &mocks.IDiscountService{}
	productDiscountService := &mocks.IProductDiscountService{}
	productService := NewProductService(productRepository, discountService, productDiscountService)

	productRepository.On("GetProductByID", "id").Return(BusinessObjects.Product{}, nil)

	_, err := productService.GetProductByID("id")

	assert.NoError(t, err)
	productRepository.AssertExpectations(t)
}

func TestCreateProduct_Success(t *testing.T) {
	productRepository := &mocks.IProductRepository{}
	discountService := &mocks.IDiscountService{}
	productDiscountService := &mocks.IProductDiscountService{}
	productService := NewProductService(productRepository, discountService, productDiscountService)

	productRepository.On("CreateProduct", mock.AnythingOfType("BusinessObjects.Product")).Return(nil)

	err := productService.CreateProduct("sellerID", "ProductName", "Description", "CategoryID", "ImageURL", 10.0, 10)

	assert.NoError(t, err)
	productRepository.AssertExpectations(t)
}
func TestCreateProduct_Error(t *testing.T) {
	productRepository := &mocks.IProductRepository{}
	discountService := &mocks.IDiscountService{}
	productDiscountService := &mocks.IProductDiscountService{}
	productService := NewProductService(productRepository, discountService, productDiscountService)

	productRepository.On("CreateProduct", mock.AnythingOfType("BusinessObjects.Product")).Return(assert.AnError)

	err := productService.CreateProduct("sellerID", "ProductName", "Description", "CategoryID", "ImageURL", 10.0, 10)

	assert.Error(t, err)
	productRepository.AssertExpectations(t)
}

func TestUpdateProduct_Success(t *testing.T) {
	productRepository := &mocks.IProductRepository{}
	discountService := &mocks.IDiscountService{}
	productDiscountService := &mocks.IProductDiscountService{}
	productService := NewProductService(productRepository, discountService, productDiscountService)

	productRepository.On("GetProductByID", "productID").Return(BusinessObjects.Product{}, nil)
	productRepository.On("UpdateProduct", mock.AnythingOfType("BusinessObjects.Product")).Return(nil)

	err := productService.UpdateProduct("productID", "sellerID", "ProductName", "Description", "CategoryID", "ImageURL", 10.0, 10)

	assert.NoError(t, err)
	productRepository.AssertExpectations(t)
}

func TestUpdateProduct_Error_GetProductByID(t *testing.T) {
	productRepository := &mocks.IProductRepository{}
	discountService := &mocks.IDiscountService{}
	productDiscountService := &mocks.IProductDiscountService{}
	productService := NewProductService(productRepository, discountService, productDiscountService)

	productRepository.On("GetProductByID", "productID").Return(BusinessObjects.Product{}, assert.AnError)

	err := productService.UpdateProduct("productID", "sellerID", "ProductName", "Description", "CategoryID", "ImageURL", 10.0, 10)

	assert.Error(t, err)
	productRepository.AssertExpectations(t)
}
func TestUpdateProduct_Error_Update(t *testing.T) {
	productRepository := &mocks.IProductRepository{}
	discountService := &mocks.IDiscountService{}
	productDiscountService := &mocks.IProductDiscountService{}
	productService := NewProductService(productRepository, discountService, productDiscountService)

	productRepository.On("GetProductByID", "productID").Return(BusinessObjects.Product{}, nil)

	productRepository.On("UpdateProduct", mock.AnythingOfType("BusinessObjects.Product")).Return(assert.AnError)

	err := productService.UpdateProduct("productID", "sellerID", "ProductName", "Description", "CategoryID", "ImageURL", 10.0, 10)

	assert.Error(t, err)
	assert.Equal(t, assert.AnError, err)
	productRepository.AssertExpectations(t)
}

func TestUpdateAllProduct_Success(t *testing.T) {
	productRepository := &mocks.IProductRepository{}
	discountService := &mocks.IDiscountService{}
	productDiscountService := &mocks.IProductDiscountService{}
	productService := NewProductService(productRepository, discountService, productDiscountService)

	product := BusinessObjects.Product{ProductID: "productID", SellerID: "sellerID", ProductName: "ProductName", Description: "Description", CategoryID: "CategoryID", ImageURL: "ImageURL", Price: 10.0, Quantity: 10}

	productRepository.On("UpdateProduct", product).Return(nil)

	err := productService.UpdateAllProduct(product)

	assert.NoError(t, err)
	productRepository.AssertExpectations(t)
}

func TestDeleteProduct_Success(t *testing.T) {
	productRepository := &mocks.IProductRepository{}
	discountService := &mocks.IDiscountService{}
	productDiscountService := &mocks.IProductDiscountService{}
	productService := NewProductService(productRepository, discountService, productDiscountService)

	productRepository.On("DeleteProduct", "id").Return(nil)

	err := productService.DeleteProduct("id")

	assert.NoError(t, err)
	productRepository.AssertExpectations(t)
}

func TestGetProductPriceAfterDiscount_Success(t *testing.T) {
	productRepository := &mocks.IProductRepository{}
	discountService := &mocks.IDiscountService{}
	productDiscountService := &mocks.IProductDiscountService{}
	productService := NewProductService(productRepository, discountService, productDiscountService)

	productRepository.On("GetProductByID", "productID").Return(BusinessObjects.Product{}, nil)
	productDiscountService.On("GetProductDiscountByID", "productID").Return([]BusinessObjects.ProductDiscount{}, nil)

	_, err := productService.GetProductPriceAfterDiscount("productID")

	assert.NoError(t, err)
	productRepository.AssertExpectations(t)
	productDiscountService.AssertExpectations(t)
}

func TestGetProductPriceAfterDiscount_Error_GetProductByID(t *testing.T) {
	productRepository := &mocks.IProductRepository{}
	discountService := &mocks.IDiscountService{}
	productDiscountService := &mocks.IProductDiscountService{}
	productService := NewProductService(productRepository, discountService, productDiscountService)

	productRepository.On("GetProductByID", "productID").Return(BusinessObjects.Product{}, assert.AnError)

	_, err := productService.GetProductPriceAfterDiscount("productID")

	assert.Error(t, err)
	productRepository.AssertExpectations(t)
}

func TestGetProductPriceAfterDiscount_Error_GetProductDiscountByID(t *testing.T) {
	productRepository := &mocks.IProductRepository{}
	discountService := &mocks.IDiscountService{}
	productDiscountService := &mocks.IProductDiscountService{}
	productService := NewProductService(productRepository, discountService, productDiscountService)

	productRepository.On("GetProductByID", "productID").Return(BusinessObjects.Product{}, nil)
	productDiscountService.On("GetProductDiscountByID", "productID").Return([]BusinessObjects.ProductDiscount{}, assert.AnError)

	_, err := productService.GetProductPriceAfterDiscount("productID")

	assert.Error(t, err)
	productRepository.AssertExpectations(t)
	productDiscountService.AssertExpectations(t)
}

func TestGetProductPriceAfterDiscount_Error_GetDiscountByID(t *testing.T) {
	productRepository := &mocks.IProductRepository{}
	discountService := &mocks.IDiscountService{}
	productDiscountService := &mocks.IProductDiscountService{}
	productService := NewProductService(productRepository, discountService, productDiscountService)

	productRepository.On("GetProductByID", "productID").Return(BusinessObjects.Product{}, nil)
	productDiscountService.On("GetProductDiscountByID", "productID").Return([]BusinessObjects.ProductDiscount{{DiscountID: "discountID"}}, nil)
	discountService.On("GetDiscountByID", "discountID").Return(BusinessObjects.Discount{}, assert.AnError)

	_, err := productService.GetProductPriceAfterDiscount("productID")

	assert.Error(t, err)
	productRepository.AssertExpectations(t)
	productDiscountService.AssertExpectations(t)
	discountService.AssertExpectations(t)
}

func TestGetProductPriceAfterDiscount_Success_PercentageDiscount(t *testing.T) {
	productRepository := &mocks.IProductRepository{}
	discountService := &mocks.IDiscountService{}
	productDiscountService := &mocks.IProductDiscountService{}
	productService := NewProductService(productRepository, discountService, productDiscountService)

	productRepository.On("GetProductByID", "productID").Return(BusinessObjects.Product{Price: 10.0}, nil)
	productDiscountService.On("GetProductDiscountByID", "productID").Return([]BusinessObjects.ProductDiscount{{DiscountID: "discountID"}}, nil)
	discountService.On("GetDiscountByID", "discountID").Return(BusinessObjects.Discount{DiscountType: "Percentage", DiscountValue: 0.1}, nil)

	price, err := productService.GetProductPriceAfterDiscount("productID")

	assert.NoError(t, err)
	assert.Equal(t, 9.0, price)
	productRepository.AssertExpectations(t)
	productDiscountService.AssertExpectations(t)
	discountService.AssertExpectations(t)
}

func TestGetProductPriceAfterDiscount_Success_FixedDiscount(t *testing.T) {
	productRepository := &mocks.IProductRepository{}
	discountService := &mocks.IDiscountService{}
	productDiscountService := &mocks.IProductDiscountService{}
	productService := NewProductService(productRepository, discountService, productDiscountService)

	productRepository.On("GetProductByID", "productID").Return(BusinessObjects.Product{Price: 10.0}, nil)
	productDiscountService.On("GetProductDiscountByID", "productID").Return([]BusinessObjects.ProductDiscount{{DiscountID: "discountID"}}, nil)
	discountService.On("GetDiscountByID", "discountID").Return(BusinessObjects.Discount{DiscountType: "Fixed", DiscountValue: 1.0}, nil)

	price, err := productService.GetProductPriceAfterDiscount("productID")

	assert.NoError(t, err)
	assert.Equal(t, 9.0, price)
	productRepository.AssertExpectations(t)
	productDiscountService.AssertExpectations(t)
	discountService.AssertExpectations(t)
}
