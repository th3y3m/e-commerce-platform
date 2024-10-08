package Services

import (
	"testing"
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Util"
	"th3y3m/e-commerce-platform/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetPaginatedFreightRateList_Success(t *testing.T) {
	// Arrange
	freightRateRepository := &mocks.IFreightRateRepository{}
	freightRateService := NewFreightRateService(freightRateRepository)

	searchValue := ""
	sortBy := ""
	courierId := ""
	pageIndex := 1
	pageSize := 10
	status := true

	freightRateRepository.On("GetPaginatedFreightRateList", searchValue, sortBy, courierId, pageIndex, pageSize, &status).Return(Util.PaginatedList[BusinessObjects.FreightRate]{}, nil)

	// Act
	_, err := freightRateService.GetPaginatedFreightRateList(searchValue, sortBy, courierId, pageIndex, pageSize, &status)

	// Assert
	assert.NoError(t, err)
	freightRateRepository.AssertExpectations(t)
}
func TestGetAllFreightRates_Success(t *testing.T) {
	// Arrange
	freightRateRepository := &mocks.IFreightRateRepository{}
	freightRateService := NewFreightRateService(freightRateRepository)

	freightRateRepository.On("GetAllFreightRates").Return([]BusinessObjects.FreightRate{}, nil)

	// Act
	_, err := freightRateService.GetAllFreightRates()

	// Assert
	assert.NoError(t, err)
	freightRateRepository.AssertExpectations(t)
}

func TestGetFreightRateByID_Success(t *testing.T) {
	// Arrange
	freightRateRepository := &mocks.IFreightRateRepository{}
	freightRateService := NewFreightRateService(freightRateRepository)

	freightRateId := "1"

	freightRateRepository.On("GetFreightRateByID", freightRateId).Return(BusinessObjects.FreightRate{}, nil)

	// Act
	_, err := freightRateService.GetFreightRateByID(freightRateId)

	// Assert
	assert.NoError(t, err)
	freightRateRepository.AssertExpectations(t)
}

func TestCreateFreightRate_Success(t *testing.T) {
	// Arrange
	freightRateRepository := &mocks.IFreightRateRepository{}
	freightRateService := NewFreightRateService(freightRateRepository)

	freightRate := BusinessObjects.FreightRate{
		CourierID:     "1",
		DistanceMinKM: 1,
		DistanceMaxKM: 10,
		CostPerKM:     10.0,
	}

	freightRateRepository.On("CreateFreightRate", mock.AnythingOfType("BusinessObjects.FreightRate")).Return(nil)

	// Act
	err := freightRateService.CreateFreightRate(freightRate.CourierID, freightRate.DistanceMinKM, freightRate.DistanceMaxKM, freightRate.CostPerKM)

	// Assert
	assert.NoError(t, err)
	freightRateRepository.AssertExpectations(t)
}
func TestCreateFreightRate_Error(t *testing.T) {
	// Arrange
	freightRateRepository := &mocks.IFreightRateRepository{}
	freightRateService := NewFreightRateService(freightRateRepository)

	freightRate := BusinessObjects.FreightRate{
		CourierID:     "1",
		DistanceMinKM: 1,
		DistanceMaxKM: 10,
		CostPerKM:     10.0,
	}

	error := assert.AnError

	freightRateRepository.On("CreateFreightRate", mock.AnythingOfType("BusinessObjects.FreightRate")).Return(error)

	// Act
	err := freightRateService.CreateFreightRate(freightRate.CourierID, freightRate.DistanceMinKM, freightRate.DistanceMaxKM, freightRate.CostPerKM)

	// Assert
	assert.Error(t, err)
	freightRateRepository.AssertExpectations(t)
}

func TestUpdateFreightRate_Success(t *testing.T) {
	// Arrange
	freightRateRepository := &mocks.IFreightRateRepository{}
	freightRateService := NewFreightRateService(freightRateRepository)

	freightRateID := "1"
	courierID := "1"
	distanceMinKM := 1
	distanceMaxKM := 10
	costPerKM := 10.0
	status := true

	freightRateRepository.On("GetFreightRateByID", freightRateID).Return(BusinessObjects.FreightRate{}, nil)
	freightRateRepository.On("UpdateFreightRate", mock.AnythingOfType("BusinessObjects.FreightRate")).Return(nil)

	// Act
	err := freightRateService.UpdateFreightRate(freightRateID, courierID, distanceMinKM, distanceMaxKM, costPerKM, status)

	// Assert
	assert.NoError(t, err)
	freightRateRepository.AssertExpectations(t)
}

func TestUpdateFreightRate_GetFreightRateByIDError(t *testing.T) {
	// Arrange
	freightRateRepository := &mocks.IFreightRateRepository{}
	freightRateService := NewFreightRateService(freightRateRepository)

	freightRateID := "1"
	courierID := "1"
	distanceMinKM := 1
	distanceMaxKM := 10
	costPerKM := 10.0
	status := true

	expectedError := assert.AnError

	freightRateRepository.On("GetFreightRateByID", freightRateID).Return(BusinessObjects.FreightRate{}, expectedError)

	// Act
	err := freightRateService.UpdateFreightRate(freightRateID, courierID, distanceMinKM, distanceMaxKM, costPerKM, status)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	freightRateRepository.AssertExpectations(t)
}

// Test case for UpdateFreightRate when UpdateFreightRate returns an error
func TestUpdateFreightRate_UpdateFreightRateError(t *testing.T) {
	// Arrange
	freightRateRepository := &mocks.IFreightRateRepository{}
	freightRateService := NewFreightRateService(freightRateRepository)

	freightRateID := "1"
	courierID := "1"
	distanceMinKM := 1
	distanceMaxKM := 10
	costPerKM := 10.0
	status := true

	freightRate := BusinessObjects.FreightRate{
		RateID:        freightRateID,
		CourierID:     courierID,
		DistanceMinKM: distanceMinKM,
		DistanceMaxKM: distanceMaxKM,
		CostPerKM:     costPerKM,
		Status:        status,
	}

	freightRateRepository.On("GetFreightRateByID", freightRateID).Return(freightRate, nil)
	expectedError := assert.AnError
	freightRateRepository.On("UpdateFreightRate", mock.AnythingOfType("BusinessObjects.FreightRate")).Return(expectedError)

	// Act
	err := freightRateService.UpdateFreightRate(freightRateID, courierID, distanceMinKM, distanceMaxKM, costPerKM, status)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	freightRateRepository.AssertExpectations(t)
}

func TestDeleteFreightRate_Success(t *testing.T) {
	// Arrange
	freightRateRepository := &mocks.IFreightRateRepository{}
	freightRateService := NewFreightRateService(freightRateRepository)

	freightRateID := "1"

	freightRateRepository.On("DeleteFreightRate", freightRateID).Return(nil)

	// Act
	err := freightRateService.DeleteFreightRate(freightRateID)

	// Assert
	assert.NoError(t, err)
	freightRateRepository.AssertExpectations(t)
}

func TestGetFreightRateByCourierID_Success(t *testing.T) {
	// Arrange
	freightRateRepository := &mocks.IFreightRateRepository{}
	freightRateService := NewFreightRateService(freightRateRepository)

	courierID := "1"

	freightRateRepository.On("GetFreightRateByCourierID", courierID).Return([]BusinessObjects.FreightRate{}, nil)

	// Act
	_, err := freightRateService.GetFreightRateByCourierID(courierID)

	// Assert
	assert.NoError(t, err)
	freightRateRepository.AssertExpectations(t)
}

func TestCalculateFreightRate_Success(t *testing.T) {
	// Arrange
	freightRateRepository := &mocks.IFreightRateRepository{}
	freightRateService := NewFreightRateService(freightRateRepository)

	courierID := "1"
	distance := 10.0

	freightRateRepository.On("GetFreightRateByCourierID", courierID).Return([]BusinessObjects.FreightRate{
		{
			CourierID:     courierID,
			DistanceMinKM: 1,
			DistanceMaxKM: 10,
			CostPerKM:     10.0,
		},
	}, nil)

	// Act
	_, err := freightRateService.CalculateFreightRate(courierID, distance)

	// Assert
	assert.NoError(t, err)
	freightRateRepository.AssertExpectations(t)
}
func TestCalculateFreightRate_Error_GetByCourtID(t *testing.T) {
	// Arrange
	freightRateRepository := &mocks.IFreightRateRepository{}
	freightRateService := NewFreightRateService(freightRateRepository)

	courierID := "1"
	distance := 10.0

	error := assert.AnError

	freightRateRepository.On("GetFreightRateByCourierID", courierID).Return([]BusinessObjects.FreightRate{
		{
			CourierID:     courierID,
			DistanceMinKM: 1,
			DistanceMaxKM: 10,
			CostPerKM:     10.0,
		},
	}, error)

	// Act
	_, err := freightRateService.CalculateFreightRate(courierID, distance)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, error, err)
	freightRateRepository.AssertExpectations(t)
}
func TestCalculateFreightRate_Success_OurOfRange(t *testing.T) {
	// Arrange
	freightRateRepository := &mocks.IFreightRateRepository{}
	freightRateService := NewFreightRateService(freightRateRepository)

	courierID := "1"
	distance := 15.0

	freightRateRepository.On("GetFreightRateByCourierID", courierID).Return([]BusinessObjects.FreightRate{
		{
			CourierID:     courierID,
			DistanceMinKM: 1,
			DistanceMaxKM: 10,
			CostPerKM:     10.0,
		},
	}, nil)

	// Act
	_, err := freightRateService.CalculateFreightRate(courierID, distance)

	// Assert
	assert.NoError(t, err)
	freightRateRepository.AssertExpectations(t)
}
