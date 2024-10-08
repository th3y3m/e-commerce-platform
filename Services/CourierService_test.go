package Services

import (
	"testing"
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Util"
	"th3y3m/e-commerce-platform/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetPaginatedCourierList_Success(t *testing.T) {
	// Arrange
	courierRepository := &mocks.ICourierRepository{}
	courierService := NewCourierService(courierRepository)

	searchValue := ""
	sortBy := ""
	pageIndex := 1
	pageSize := 10
	status := true

	courierRepository.On("GetPaginatedCourierList", searchValue, sortBy, pageIndex, pageSize, &status).Return(Util.PaginatedList[BusinessObjects.Courier]{}, nil)

	// Act
	_, err := courierService.GetPaginatedCourierList(searchValue, sortBy, pageIndex, pageSize, &status)

	// Assert
	assert.NoError(t, err)
	courierRepository.AssertExpectations(t)
}

func TestGetPaginatedCourierList_Error(t *testing.T) {
	// Arrange
	courierRepository := &mocks.ICourierRepository{}
	courierService := NewCourierService(courierRepository)

	searchValue := ""
	sortBy := ""
	pageIndex := 1
	pageSize := 10
	status := true

	error := assert.AnError

	courierRepository.On("GetPaginatedCourierList", searchValue, sortBy, pageIndex, pageSize, &status).Return(Util.PaginatedList[BusinessObjects.Courier]{}, error)

	// Act
	_, err := courierService.GetPaginatedCourierList(searchValue, sortBy, pageIndex, pageSize, &status)

	// Assert
	assert.Error(t, err)
	courierRepository.AssertExpectations(t)
}

func TestGetAllCouriers_Success(t *testing.T) {
	// Arrange
	courierRepository := &mocks.ICourierRepository{}
	courierService := NewCourierService(courierRepository)

	courierRepository.On("GetAllCouriers").Return([]BusinessObjects.Courier{}, nil)

	// Act
	_, err := courierService.GetAllCouriers()

	// Assert
	assert.NoError(t, err)
	courierRepository.AssertExpectations(t)
}

func TestGetAllCouriers_Error(t *testing.T) {
	// Arrange
	courierRepository := &mocks.ICourierRepository{}
	courierService := NewCourierService(courierRepository)

	error := assert.AnError

	courierRepository.On("GetAllCouriers").Return([]BusinessObjects.Courier{}, error)

	// Act
	_, err := courierService.GetAllCouriers()

	// Assert
	assert.Error(t, err)
	courierRepository.AssertExpectations(t)
}

func TestGetCourierByID_Success(t *testing.T) {
	// Arrange
	courierRepository := &mocks.ICourierRepository{}
	courierService := NewCourierService(courierRepository)

	courierID := "COUR1234567890"

	courierRepository.On("GetCourierByID", courierID).Return(BusinessObjects.Courier{}, nil)

	// Act
	_, err := courierService.GetCourierByID(courierID)

	// Assert
	assert.NoError(t, err)
	courierRepository.AssertExpectations(t)
}

func TestGetCourierByID_Error(t *testing.T) {
	// Arrange
	courierRepository := &mocks.ICourierRepository{}
	courierService := NewCourierService(courierRepository)

	courierID := "COUR1234567890"

	error := assert.AnError

	courierRepository.On("GetCourierByID", courierID).Return(BusinessObjects.Courier{}, error)

	// Act
	_, err := courierService.GetCourierByID(courierID)

	// Assert
	assert.Error(t, err)
	courierRepository.AssertExpectations(t)
}

func TestCreateCourier_Success(t *testing.T) {
	// Arrange
	courierRepository := &mocks.ICourierRepository{}
	courierService := NewCourierService(courierRepository)

	Courier := "Courier Name"

	courierRepository.On("CreateCourier", mock.AnythingOfType("BusinessObjects.Courier")).Return(nil)

	// Act
	err := courierService.CreateCourier(Courier)

	// Assert
	assert.NoError(t, err)
	courierRepository.AssertExpectations(t)
}

func TestCreateCourier_Error(t *testing.T) {
	// Arrange
	courierRepository := &mocks.ICourierRepository{}
	courierService := NewCourierService(courierRepository)

	Courier := "Courier Name"

	error := assert.AnError

	courierRepository.On("CreateCourier", mock.AnythingOfType("BusinessObjects.Courier")).Return(error)

	// Act
	err := courierService.CreateCourier(Courier)

	// Assert
	assert.Error(t, err)
	courierRepository.AssertExpectations(t)
}

func TestUpdateCourier_Success(t *testing.T) {
	// Arrange
	courierRepository := &mocks.ICourierRepository{}
	courierService := NewCourierService(courierRepository)

	courierID := "COUR1234567890"
	Courier := "Courier Name"
	status := true

	courierRepository.On("GetCourierByID", courierID).Return(BusinessObjects.Courier{}, nil)
	courierRepository.On("UpdateCourier", mock.AnythingOfType("BusinessObjects.Courier")).Return(nil)

	// Act
	err := courierService.UpdateCourier(courierID, Courier, status)

	// Assert
	assert.NoError(t, err)
	courierRepository.AssertExpectations(t)
}

func TestUpdateCourier_Error(t *testing.T) {
	// Arrange
	courierRepository := &mocks.ICourierRepository{}
	courierService := NewCourierService(courierRepository)

	courierID := "COUR1234567890"
	Courier := "Courier Name"
	status := true

	error := assert.AnError

	courierRepository.On("GetCourierByID", courierID).Return(BusinessObjects.Courier{}, error)

	// Act
	err := courierService.UpdateCourier(courierID, Courier, status)

	// Assert
	assert.Error(t, err)
	courierRepository.AssertExpectations(t)
}

func TestDeleteCourier_Success(t *testing.T) {
	// Arrange
	courierRepository := &mocks.ICourierRepository{}
	courierService := NewCourierService(courierRepository)

	courierID := "COUR1234567890"

	courierRepository.On("DeleteCourier", courierID).Return(nil)

	// Act
	err := courierService.DeleteCourier(courierID)

	// Assert
	assert.NoError(t, err)
	courierRepository.AssertExpectations(t)
}

func TestDeleteCourier_Error(t *testing.T) {
	// Arrange
	courierRepository := &mocks.ICourierRepository{}
	courierService := NewCourierService(courierRepository)

	courierID := "COUR1234567890"

	error := assert.AnError

	courierRepository.On("DeleteCourier", courierID).Return(error)

	// Act
	err := courierService.DeleteCourier(courierID)

	// Assert
	assert.Error(t, err)
	courierRepository.AssertExpectations(t)
}
