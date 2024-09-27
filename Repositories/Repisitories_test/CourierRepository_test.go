package Repositories_test

import (
	"testing"
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Repositories"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestGetAllCouriers(t *testing.T) {
	log := logrus.New()
	mockDB := new(MockDB)

	mocks := []BusinessObjects.Courier{
		{CourierID: "1", Courier: "test1"},
		{CourierID: "2", Courier: "test2"},
	}

	// Mock the Find function to simulate finding Couriers
	mockDB.On("Find", mock.Anything, mock.Anything).Return(&gorm.DB{}).Run(func(args mock.Arguments) {
		// Simulate GORM populating the Couriers
		arg := args.Get(0).(*[]BusinessObjects.Courier)
		*arg = mocks
	})

	// Create CourierRepository with mock DB
	CourierRepository := Repositories.NewCourierRepository(log, mockDB)

	// Call GetAllCouriers
	Couriers, err := CourierRepository.GetAllCouriers()

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, 2, len(Couriers))
	assert.Equal(t, "test1", Couriers[0].Courier)
	assert.Equal(t, "test2", Couriers[1].Courier)

	// Verify that all expectations were met
	mockDB.AssertExpectations(t)
}

func TestGetCourierByID(t *testing.T) {
	// Set up mock logger and DB
	log := logrus.New()
	mockDB := new(MockDB)

	// Create mock return value for the specific Courier
	mockCourier := BusinessObjects.Courier{
		CourierID: "1",
		Courier:   "test1",
	}

	// Mock the First function to simulate finding a Courier by ID
	// We expect the query as a slice []interface{}{"courier_id = ?", "1"}
	mockDB.On("First", mock.Anything, []interface{}{"courier_id = ?", "1"}).Return(&gorm.DB{}).Run(func(args mock.Arguments) {
		// Simulate GORM populating the Courier
		arg := args.Get(0).(*BusinessObjects.Courier)
		*arg = mockCourier
	})

	// Create CourierRepository with mock DB
	repo := Repositories.NewCourierRepository(log, mockDB)

	// Call GetCourierByID
	Courier, err := repo.GetCourierByID("1")

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, "1", Courier.CourierID)
	assert.Equal(t, "test1", Courier.Courier)

	// Verify that all expectations were met
	mockDB.AssertExpectations(t)
}

func TestGetCourierByIDUsingDDT(t *testing.T) {
	// Set up mock logger and DB
	log := logrus.New()
	mockDB := new(MockDB)

	// Define test cases in a table-driven format
	testCases := []struct {
		name        string                  // Name of the test case
		CourierID   string                  // Input: Courier ID
		mockCourier BusinessObjects.Courier // The Courier we expect to be returned
		shouldFail  bool                    // Whether the test should fail (if Courier not found)
	}{
		{
			name:      "Courier Found - ID 1",
			CourierID: "1",
			mockCourier: BusinessObjects.Courier{
				CourierID: "1",
				Courier:   "test1",
			},
			shouldFail: false,
		},
		{
			name:      "Courier Found - ID 2",
			CourierID: "2",
			mockCourier: BusinessObjects.Courier{
				CourierID: "2",
				Courier:   "test2",
			},
			shouldFail: false,
		},
		{
			name:        "Courier Not Found",
			CourierID:   "999",
			mockCourier: BusinessObjects.Courier{},
			shouldFail:  true, // Expect an error when the Courier isn't found
		},
	}

	// Set up mock responses based on the test cases
	for _, tc := range testCases {
		if tc.shouldFail {
			// For the test case where we expect a failure, mock an error response
			mockDB.On("First", mock.Anything, []interface{}{"courier_id = ?", tc.CourierID}).Return(&gorm.DB{Error: gorm.ErrRecordNotFound})
		} else {
			// For test cases where a Courier is found, mock the Courier response
			mockDB.On("First", mock.Anything, []interface{}{"courier_id = ?", tc.CourierID}).Return(&gorm.DB{}).Run(func(args mock.Arguments) {
				arg := args.Get(0).(*BusinessObjects.Courier)
				*arg = tc.mockCourier
			})
		}
	}

	// Iterate over test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create CourierRepository with mock DB
			repo := Repositories.NewCourierRepository(log, mockDB)

			// Call GetCourierByID
			Courier, err := repo.GetCourierByID(tc.CourierID)

			// Assertions based on whether an error was expected
			if tc.shouldFail {
				assert.Error(t, err)
				assert.Equal(t, BusinessObjects.Courier{}, Courier) // Ensure no Courier is returned
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.mockCourier.CourierID, Courier.CourierID)
				assert.Equal(t, tc.mockCourier.Courier, Courier.Courier)
			}
		})
	}

	// Verify that all expectations were met
	mockDB.AssertExpectations(t)
}

func TestCreateCourier(t *testing.T) {
	// Set up mock logger and DB
	log := logrus.New()
	mockDB := new(MockDB)

	// Define test cases
	testCases := []struct {
		name       string                  // Name of the test case
		Courier    BusinessObjects.Courier // Input: Courier to create
		shouldFail bool                    // Whether the test should fail
	}{
		{
			name: "Create Courier Success",
			Courier: BusinessObjects.Courier{
				CourierID: "1",
				Courier:   "test1",
			},
			shouldFail: false,
		},
	}

	// Set up mock responses based on the test cases
	for _, tc := range testCases {
		if tc.shouldFail {
			mockDB.On("Create", &tc.Courier).Return(&gorm.DB{Error: gorm.ErrInvalidData})
		} else {
			mockDB.On("Create", &tc.Courier).Return(&gorm.DB{})
			mockDB.On("First", mock.Anything, []interface{}{"courier_id = ?", tc.Courier.CourierID}).Return(&gorm.DB{}).Run(func(args mock.Arguments) {
				arg := args.Get(0).(*BusinessObjects.Courier)
				*arg = tc.Courier
			})
		}
	}

	// Iterate over test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create CourierRepository with mock DB
			repo := Repositories.NewCourierRepository(log, mockDB)

			// Call CreateCourier
			err := repo.CreateCourier(tc.Courier)

			if err != nil {
				t.Log(err)
			}

			Courier, err := repo.GetCourierByID(tc.Courier.CourierID)

			// Assertions based on whether an error was expected
			if tc.shouldFail {
				assert.Error(t, err)
				assert.Equal(t, BusinessObjects.Courier{}, Courier)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.Courier.Courier, Courier.Courier)
			}
		})
	}

	// Verify that all expectations were met
	mockDB.AssertExpectations(t)
}

func TestUpdateCourier(t *testing.T) {
	// Set up mock logger and DB
	log := logrus.New()
	mockDB := new(MockDB)

	// Define test cases
	testCases := []struct {
		name       string                  // Name of the test case
		Courier    BusinessObjects.Courier // Input: Courier to update
		shouldFail bool                    // Whether the test should fail
	}{
		{
			name: "Update Courier Success",
			Courier: BusinessObjects.Courier{
				CourierID: "1",
				Courier:   "test1",
			},
			shouldFail: false,
		},
		{
			name: "Update Courier Failure",
			Courier: BusinessObjects.Courier{
				CourierID: "2",
				Courier:   "test2",
			},
			shouldFail: true, // Expect failure
		},
	}

	// Set up mock responses based on the test cases
	for _, tc := range testCases {
		if tc.shouldFail {
			mockDB.On("Save", &tc.Courier).Return(&gorm.DB{Error: gorm.ErrInvalidData})
		} else {
			mockDB.On("Save", &tc.Courier).Return(&gorm.DB{})
		}
	}

	// Iterate over test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create CourierRepository with mock DB
			repo := Repositories.NewCourierRepository(log, mockDB)

			// Call UpdateCourier
			err := repo.UpdateCourier(tc.Courier)

			// Assertions based on whether an error was expected
			if tc.shouldFail {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}

	// Verify that all expectations were met
	mockDB.AssertExpectations(t)
}

// func TestDeleteCourier(t *testing.T) {
// 	// Set up mock logger and DB
// 	log := logrus.New()
// 	mockDB := new(MockDB)

// 	// Define test case
// 	testCase := BusinessObjects.Courier{
// 		CourierID: "1",
// 		Status:    true,
// 	}

// 	// Create a separate Courier object for the Model method call
// 	expectedCourier := BusinessObjects.Courier{
// 		CourierID: "1",
// 		Status:    true,
// 	}

// 	// Mock the Model method to return the mock DB
// 	mockDB.On("Model", &expectedCourier).Return(mockDB)

// 	// Mock the Save method to simulate deleting a Courier
// 	mockDB.On("Save", mock.Anything).Return(&gorm.DB{}).Run(func(args mock.Arguments) {
// 		arg := args.Get(0).(*BusinessObjects.Courier)
// 		arg.Status = false
// 	})

// 	// Mock the First method to simulate fetching the Courier by ID
// 	mockDB.On("First", mock.Anything, "courier_id = ?", testCase.CourierID).Return(&gorm.DB{}).Run(func(args mock.Arguments) {
// 		arg := args.Get(0).(*BusinessObjects.Courier)
// 		*arg = testCase
// 	})

// 	// Create CourierRepository with mock DB
// 	repo := Repositories.NewCourierRepository(log, mockDB)

// 	// Call DeleteCourier
// 	err := repo.DeleteCourier(testCase.CourierID)
// 	if err != nil {
// 		t.Log(err)
// 	}

// 	// Fetch the courier to verify the status
// 	courier, err := repo.GetCourierByID(testCase.CourierID)
// 	if err != nil {
// 		t.Log(err)
// 	}

// 	// Assertions
// 	assert.NoError(t, err)
// 	assert.Equal(t, false, courier.Status)
// 	mockDB.AssertExpectations(t)
// }
