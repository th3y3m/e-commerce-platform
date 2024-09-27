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

func TestGetAllProducts(t *testing.T) {
	log := logrus.New()
	mockDB := new(MockDB)

	mocks := []BusinessObjects.Product{
		{ProductID: "1", ProductName: "test1"},
		{ProductID: "2", ProductName: "test2"},
	}

	// Mock the Find function to simulate finding Products
	mockDB.On("Find", mock.Anything, mock.Anything).Return(&gorm.DB{}).Run(func(args mock.Arguments) {
		// Simulate GORM populating the Products
		arg := args.Get(0).(*[]BusinessObjects.Product)
		*arg = mocks
	})

	// Create ProductRepository with mock DB
	ProductRepository := Repositories.NewProductRepository(log, mockDB)

	// Call GetAllProducts
	Products, err := ProductRepository.GetAllProducts()

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, 2, len(Products))
	assert.Equal(t, "test1", Products[0].ProductName)
	assert.Equal(t, "test2", Products[1].ProductName)

	// Verify that all expectations were met
	mockDB.AssertExpectations(t)
}

func TestGetProductByID(t *testing.T) {
	// Set up mock logger and DB
	log := logrus.New()
	mockDB := new(MockDB)

	// Create mock return value for the specific Product
	mockProduct := BusinessObjects.Product{
		ProductID:   "1",
		ProductName: "test1",
	}

	// Mock the First function to simulate finding a Product by ID
	// We expect the query as a slice []interface{}{"product_id = ?", "1"}
	mockDB.On("First", mock.Anything, []interface{}{"product_id = ?", "1"}).Return(&gorm.DB{}).Run(func(args mock.Arguments) {
		// Simulate GORM populating the Product
		arg := args.Get(0).(*BusinessObjects.Product)
		*arg = mockProduct
	})

	// Create ProductRepository with mock DB
	repo := Repositories.NewProductRepository(log, mockDB)

	// Call GetProductByID
	Product, err := repo.GetProductByID("1")

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, "1", Product.ProductID)
	assert.Equal(t, "test1", Product.ProductName)

	// Verify that all expectations were met
	mockDB.AssertExpectations(t)
}

func TestGetProductByIDUsingDDT(t *testing.T) {
	// Set up mock logger and DB
	log := logrus.New()
	mockDB := new(MockDB)

	// Define test cases in a table-driven format
	testCases := []struct {
		name        string                  // Name of the test case
		ProductID   string                  // Input: Product ID
		mockProduct BusinessObjects.Product // The Product we expect to be returned
		shouldFail  bool                    // Whether the test should fail (if Product not found)
	}{
		{
			name:      "Product Found - ID 1",
			ProductID: "1",
			mockProduct: BusinessObjects.Product{
				ProductID:   "1",
				ProductName: "test1",
			},
			shouldFail: false,
		},
		{
			name:      "Product Found - ID 2",
			ProductID: "2",
			mockProduct: BusinessObjects.Product{
				ProductID:   "2",
				ProductName: "test2",
			},
			shouldFail: false,
		},
		{
			name:        "Product Not Found",
			ProductID:   "999",
			mockProduct: BusinessObjects.Product{},
			shouldFail:  true, // Expect an error when the Product isn't found
		},
	}

	// Set up mock responses based on the test cases
	for _, tc := range testCases {
		if tc.shouldFail {
			// For the test case where we expect a failure, mock an error response
			mockDB.On("First", mock.Anything, []interface{}{"product_id = ?", tc.ProductID}).Return(&gorm.DB{Error: gorm.ErrRecordNotFound})
		} else {
			// For test cases where a Product is found, mock the Product response
			mockDB.On("First", mock.Anything, []interface{}{"product_id = ?", tc.ProductID}).Return(&gorm.DB{}).Run(func(args mock.Arguments) {
				arg := args.Get(0).(*BusinessObjects.Product)
				*arg = tc.mockProduct
			})
		}
	}

	// Iterate over test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create ProductRepository with mock DB
			repo := Repositories.NewProductRepository(log, mockDB)

			// Call GetProductByID
			Product, err := repo.GetProductByID(tc.ProductID)

			// Assertions based on whether an error was expected
			if tc.shouldFail {
				assert.Error(t, err)
				assert.Equal(t, BusinessObjects.Product{}, Product) // Ensure no Product is returned
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.mockProduct.ProductID, Product.ProductID)
				assert.Equal(t, tc.mockProduct.ProductName, Product.ProductName)
			}
		})
	}

	// Verify that all expectations were met
	mockDB.AssertExpectations(t)
}

func TestCreateProduct(t *testing.T) {
	// Set up mock logger and DB
	log := logrus.New()
	mockDB := new(MockDB)

	// Define test cases
	testCases := []struct {
		name       string                  // Name of the test case
		Product    BusinessObjects.Product // Input: Product to create
		shouldFail bool                    // Whether the test should fail
	}{
		{
			name: "Create Product Success",
			Product: BusinessObjects.Product{
				ProductID:   "1",
				ProductName: "test1",
			},
			shouldFail: false,
		},
	}

	// Set up mock responses based on the test cases
	for _, tc := range testCases {
		if tc.shouldFail {
			mockDB.On("Create", &tc.Product).Return(&gorm.DB{Error: gorm.ErrInvalidData})
		} else {
			mockDB.On("Create", &tc.Product).Return(&gorm.DB{})
			mockDB.On("First", mock.Anything, []interface{}{"product_id = ?", tc.Product.ProductID}).Return(&gorm.DB{}).Run(func(args mock.Arguments) {
				arg := args.Get(0).(*BusinessObjects.Product)
				*arg = tc.Product
			})
		}
	}

	// Iterate over test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create ProductRepository with mock DB
			repo := Repositories.NewProductRepository(log, mockDB)

			// Call CreateProduct
			err := repo.CreateProduct(tc.Product)

			if err != nil {
				t.Log(err)
			}

			Product, err := repo.GetProductByID(tc.Product.ProductID)

			// Assertions based on whether an error was expected
			if tc.shouldFail {
				assert.Error(t, err)
				assert.Equal(t, BusinessObjects.Product{}, Product)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.Product.ProductName, Product.ProductName)
			}
		})
	}

	// Verify that all expectations were met
	mockDB.AssertExpectations(t)
}

func TestUpdateProduct(t *testing.T) {
	// Set up mock logger and DB
	log := logrus.New()
	mockDB := new(MockDB)

	// Define test cases
	testCases := []struct {
		name       string                  // Name of the test case
		Product    BusinessObjects.Product // Input: Product to update
		shouldFail bool                    // Whether the test should fail
	}{
		{
			name: "Update Product Success",
			Product: BusinessObjects.Product{
				ProductID:   "1",
				ProductName: "test1",
			},
			shouldFail: false,
		},
		{
			name: "Update Product Failure",
			Product: BusinessObjects.Product{
				ProductID:   "2",
				ProductName: "test2",
			},
			shouldFail: true, // Expect failure
		},
	}

	// Set up mock responses based on the test cases
	for _, tc := range testCases {
		if tc.shouldFail {
			mockDB.On("Save", &tc.Product).Return(&gorm.DB{Error: gorm.ErrInvalidData})
		} else {
			mockDB.On("Save", &tc.Product).Return(&gorm.DB{})
		}
	}

	// Iterate over test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create ProductRepository with mock DB
			repo := Repositories.NewProductRepository(log, mockDB)

			// Call UpdateProduct
			err := repo.UpdateProduct(tc.Product)

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

// func TestDeleteProduct(t *testing.T) {
// 	// Set up mock logger and DB
// 	log := logrus.New()
// 	mockDB := new(MockDB)
// 	mockGormDB := &gorm.DB{} // Properly initialize the gorm.DB object

// 	// Define test case
// 	testCase := BusinessObjects.Product{
// 		ProductID: "1",
// 		Status:    true,
// 	}

// 	// Mock the Model method to return the mock Gorm DB
// 	mockDB.On("Model", &BusinessObjects.Product{}).Return(mockGormDB).Run(func(args mock.Arguments) {
// 		arg := args.Get(0).(*BusinessObjects.Product)
// 		arg.ProductID = testCase.ProductID
// 		arg.Status = testCase.Status
// 	})

// 	// Mock the Where method to return the mock Gorm DB
// 	mockDB.On("Where", "product_id = ?", testCase.ProductID).Return(mockGormDB)

// 	// Mock the Update method to simulate setting the status to false
// 	mockDB.On("Update", "status", false).Return(mockGormDB)

// 	// Create ProductRepository with mock DB
// 	repo := Repositories.NewProductRepository(log, mockDB)

// 	// Call DeleteProduct
// 	err := repo.DeleteProduct(testCase.ProductID)

// 	// Assertions
// 	assert.NoError(t, err)
// 	mockDB.AssertExpectations(t)
// }
