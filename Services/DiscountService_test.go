package Services

import (
	"testing"
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Util"
	"th3y3m/e-commerce-platform/mocks"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetPaginatedDiscountList_Success(t *testing.T) {
	discountRepository := &mocks.IDiscountRepository{}

	discountService := NewDiscountService(discountRepository)

	searchValue := ""
	sortBy := ""
	pageIndex := 1
	pageSize := 10
	status := true

	discountRepository.On("GetPaginatedDiscountList", searchValue, sortBy, pageIndex, pageSize, &status).Return(Util.PaginatedList[BusinessObjects.Discount]{}, nil)

	_, err := discountService.GetPaginatedDiscountList(searchValue, sortBy, pageIndex, pageSize, &status)

	assert.NoError(t, err)
	discountRepository.AssertExpectations(t)
}

func TestGetPaginatedDiscountList_Error(t *testing.T) {
	discountRepository := &mocks.IDiscountRepository{}

	discountService := NewDiscountService(discountRepository)

	searchValue := ""
	sortBy := ""
	pageIndex := 1
	pageSize := 10
	status := true

	error := assert.AnError

	discountRepository.On("GetPaginatedDiscountList", searchValue, sortBy, pageIndex, pageSize, &status).Return(Util.PaginatedList[BusinessObjects.Discount]{}, error)

	_, err := discountService.GetPaginatedDiscountList(searchValue, sortBy, pageIndex, pageSize, &status)

	assert.Error(t, err)
	discountRepository.AssertExpectations(t)
}

func TestGetAllDiscounts_Success(t *testing.T) {
	discountRepository := &mocks.IDiscountRepository{}

	discountService := NewDiscountService(discountRepository)

	discountRepository.On("GetAllDiscounts").Return([]BusinessObjects.Discount{}, nil)

	_, err := discountService.GetAllDiscounts()

	assert.NoError(t, err)
	discountRepository.AssertExpectations(t)
}

func TestGetDiscountByID_Success(t *testing.T) {
	discountRepository := &mocks.IDiscountRepository{}

	discountService := NewDiscountService(discountRepository)

	id := "DISC123"

	discountRepository.On("GetDiscountByID", id).Return(BusinessObjects.Discount{}, nil)

	_, err := discountService.GetDiscountByID(id)

	assert.NoError(t, err)
	discountRepository.AssertExpectations(t)
}

func TestCreateDiscount_Success(t *testing.T) {
	discountRepository := &mocks.IDiscountRepository{}

	discountService := NewDiscountService(discountRepository)

	DiscountType := "Percentage"
	DiscountValue := 10.0
	startDate := time.Now()
	endDate := time.Now()

	discountRepository.On("CreateDiscount", mock.AnythingOfType("BusinessObjects.Discount")).Return(nil)

	err := discountService.CreateDiscount(DiscountType, DiscountValue, startDate, endDate)

	assert.NoError(t, err)
	discountRepository.AssertExpectations(t)
}
func TestCreateDiscount_Error(t *testing.T) {
	discountRepository := &mocks.IDiscountRepository{}

	discountService := NewDiscountService(discountRepository)

	DiscountType := "Percentage"
	DiscountValue := 10.0
	startDate := time.Now()
	endDate := time.Now()

	error := assert.AnError

	discountRepository.On("CreateDiscount", mock.AnythingOfType("BusinessObjects.Discount")).Return(error)

	err := discountService.CreateDiscount(DiscountType, DiscountValue, startDate, endDate)

	assert.Error(t, err)
	discountRepository.AssertExpectations(t)
}

func TestUpdateDiscount_Success(t *testing.T) {
	discountRepository := &mocks.IDiscountRepository{}

	discountService := NewDiscountService(discountRepository)

	discountID := "DISC123"
	DiscountType := "Percentage"
	DiscountValue := 10.0
	startDate := time.Now()
	endDate := time.Now()

	discountRepository.On("GetDiscountByID", discountID).Return(BusinessObjects.Discount{}, nil)

	discountRepository.On("UpdateDiscount", mock.AnythingOfType("BusinessObjects.Discount")).Return(nil)

	err := discountService.UpdateDiscount(discountID, DiscountType, DiscountValue, startDate, endDate)

	assert.NoError(t, err)
	discountRepository.AssertExpectations(t)
}

func TestUpdateDiscount_Error(t *testing.T) {
	discountRepository := &mocks.IDiscountRepository{}

	discountService := NewDiscountService(discountRepository)

	discountID := "DISC123"
	DiscountType := "Percentage"
	DiscountValue := 10.0
	startDate := time.Now()
	endDate := time.Now()

	error := assert.AnError

	discountRepository.On("GetDiscountByID", discountID).Return(BusinessObjects.Discount{}, error)

	err := discountService.UpdateDiscount(discountID, DiscountType, DiscountValue, startDate, endDate)

	assert.Error(t, err)
	discountRepository.AssertExpectations(t)
}

func TestDeleteDiscount_Success(t *testing.T) {
	discountRepository := &mocks.IDiscountRepository{}

	discountService := NewDiscountService(discountRepository)

	discountID := "DISC123"

	discountRepository.On("DeleteDiscount", discountID).Return(nil)

	err := discountService.DeleteDiscount(discountID)

	assert.NoError(t, err)
	discountRepository.AssertExpectations(t)
}
