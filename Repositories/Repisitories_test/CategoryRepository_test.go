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

func TestGetAllCategories(t *testing.T) {
	log := logrus.New()
	mockDB := new(MockDB)

	mocks := []BusinessObjects.Category{
		{CategoryID: "1", CategoryName: "test1"},
		{CategoryID: "2", CategoryName: "test2"},
	}

	// Mock the Find function to simulate finding Categorys
	mockDB.On("Find", mock.Anything, mock.Anything).Return(&gorm.DB{}).Run(func(args mock.Arguments) {
		// Simulate GORM populating the Categorys
		arg := args.Get(0).(*[]BusinessObjects.Category)
		*arg = mocks
	})

	// Create CategoryRepository with mock DB
	categoryRepository := Repositories.NewCategoryRepository(log, mockDB)

	// Call GetAllCategorys
	categories, err := categoryRepository.GetAllCategories()

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, 2, len(categories))
	assert.Equal(t, "test1", categories[0].CategoryName)
	assert.Equal(t, "test2", categories[1].CategoryName)

	// Verify that all expectations were met
	mockDB.AssertExpectations(t)
}

func TestGetCategoryByID(t *testing.T) {
	// Set up mock logger and DB
	log := logrus.New()
	mockDB := new(MockDB)

	// Create mock return value for the specific Category
	mockCategory := BusinessObjects.Category{
		CategoryID:   "1",
		CategoryName: "test1",
	}

	// Mock the First function to simulate finding a Category by ID
	// We expect the query as a slice []interface{}{"Category_id = ?", "1"}
	mockDB.On("First", mock.Anything, []interface{}{"category_id = ?", "1"}).Return(&gorm.DB{}).Run(func(args mock.Arguments) {
		// Simulate GORM populating the Category
		arg := args.Get(0).(*BusinessObjects.Category)
		*arg = mockCategory
	})

	// Create CategoryRepository with mock DB
	repo := Repositories.NewCategoryRepository(log, mockDB)

	// Call GetCategoryByID
	Category, err := repo.GetCategoryByID("1")

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, "1", Category.CategoryID)
	assert.Equal(t, "test1", Category.CategoryName)

	// Verify that all expectations were met
	mockDB.AssertExpectations(t)
}

func TestGetCategoryByIDUsingDDT(t *testing.T) {
	// Set up mock logger and DB
	log := logrus.New()
	mockDB := new(MockDB)

	// Define test cases in a table-driven format
	testCases := []struct {
		name         string                   // Name of the test case
		CategoryID   string                   // Input: Category ID
		mockCategory BusinessObjects.Category // The Category we expect to be returned
		shouldFail   bool                     // Whether the test should fail (if Category not found)
	}{
		{
			name:       "Category Found - ID 1",
			CategoryID: "1",
			mockCategory: BusinessObjects.Category{
				CategoryID:   "1",
				CategoryName: "test1",
			},
			shouldFail: false,
		},
		{
			name:       "Category Found - ID 2",
			CategoryID: "2",
			mockCategory: BusinessObjects.Category{
				CategoryID:   "2",
				CategoryName: "test2",
			},
			shouldFail: false,
		},
		{
			name:         "Category Not Found",
			CategoryID:   "999",
			mockCategory: BusinessObjects.Category{},
			shouldFail:   true, // Expect an error when the Category isn't found
		},
	}

	// Set up mock responses based on the test cases
	for _, tc := range testCases {
		if tc.shouldFail {
			// For the test case where we expect a failure, mock an error response
			mockDB.On("First", mock.Anything, []interface{}{"category_id = ?", tc.CategoryID}).Return(&gorm.DB{Error: gorm.ErrRecordNotFound})
		} else {
			// For test cases where a Category is found, mock the Category response
			mockDB.On("First", mock.Anything, []interface{}{"category_id = ?", tc.CategoryID}).Return(&gorm.DB{}).Run(func(args mock.Arguments) {
				arg := args.Get(0).(*BusinessObjects.Category)
				*arg = tc.mockCategory
			})
		}
	}

	// Iterate over test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create CategoryRepository with mock DB
			repo := Repositories.NewCategoryRepository(log, mockDB)

			// Call GetCategoryByID
			Category, err := repo.GetCategoryByID(tc.CategoryID)

			// Assertions based on whether an error was expected
			if tc.shouldFail {
				assert.Error(t, err)
				assert.Equal(t, BusinessObjects.Category{}, Category) // Ensure no Category is returned
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.mockCategory.CategoryID, Category.CategoryID)
				assert.Equal(t, tc.mockCategory.CategoryName, Category.CategoryName)
			}
		})
	}

	// Verify that all expectations were met
	mockDB.AssertExpectations(t)
}

func TestCreateCategory(t *testing.T) {
	// Set up mock logger and DB
	log := logrus.New()
	mockDB := new(MockDB)

	// Define test cases
	testCases := []struct {
		name       string                   // Name of the test case
		Category   BusinessObjects.Category // Input: Category to create
		shouldFail bool                     // Whether the test should fail
	}{
		{
			name: "Create Category Success",
			Category: BusinessObjects.Category{
				CategoryID:   "1",
				CategoryName: "test1",
			},
			shouldFail: false,
		},
	}

	// Set up mock responses based on the test cases
	for _, tc := range testCases {
		if tc.shouldFail {
			mockDB.On("Create", &tc.Category).Return(&gorm.DB{Error: gorm.ErrInvalidData})
		} else {
			mockDB.On("Create", &tc.Category).Return(&gorm.DB{})
			mockDB.On("First", mock.Anything, []interface{}{"category_id = ?", tc.Category.CategoryID}).Return(&gorm.DB{}).Run(func(args mock.Arguments) {
				arg := args.Get(0).(*BusinessObjects.Category)
				*arg = tc.Category
			})
		}
	}

	// Iterate over test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create CategoryRepository with mock DB
			repo := Repositories.NewCategoryRepository(log, mockDB)

			// Call CreateCategory
			err := repo.CreateCategory(tc.Category)

			if err != nil {
				t.Log(err)
			}

			Category, err := repo.GetCategoryByID(tc.Category.CategoryID)

			// Assertions based on whether an error was expected
			if tc.shouldFail {
				assert.Error(t, err)
				assert.Equal(t, BusinessObjects.Category{}, Category)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.Category.CategoryName, Category.CategoryName)
			}
		})
	}

	// Verify that all expectations were met
	mockDB.AssertExpectations(t)
}

func TestUpdateCategory(t *testing.T) {
	// Set up mock logger and DB
	log := logrus.New()
	mockDB := new(MockDB)

	// Define test cases
	testCases := []struct {
		name       string                   // Name of the test case
		Category   BusinessObjects.Category // Input: Category to update
		shouldFail bool                     // Whether the test should fail
	}{
		{
			name: "Update Category Success",
			Category: BusinessObjects.Category{
				CategoryID:   "1",
				CategoryName: "test1",
			},
			shouldFail: false,
		},
		{
			name: "Update Category Failure",
			Category: BusinessObjects.Category{
				CategoryID:   "2",
				CategoryName: "test2",
			},
			shouldFail: true, // Expect failure
		},
	}

	// Set up mock responses based on the test cases
	for _, tc := range testCases {
		if tc.shouldFail {
			mockDB.On("Save", &tc.Category).Return(&gorm.DB{Error: gorm.ErrInvalidData})
		} else {
			mockDB.On("Save", &tc.Category).Return(&gorm.DB{})
		}
	}

	// Iterate over test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create CategoryRepository with mock DB
			repo := Repositories.NewCategoryRepository(log, mockDB)

			// Call UpdateCategory
			err := repo.UpdateCategory(tc.Category)

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

func TestDeleteCategory(t *testing.T) {
	// Set up mock logger and DB
	log := logrus.New()
	mockDB := new(MockDB)

	// Define test cases
	testCases := []struct {
		CategoryID string
	}{
		{CategoryID: "1"},
		{CategoryID: "2"},
	}

	for _, tc := range testCases {
		// Set up mock expectations
		mockDB.On("Delete", &BusinessObjects.Category{}, "category_id = ?", tc.CategoryID).Return(&gorm.DB{})

		// Create CategoryRepository with mock DB
		repo := Repositories.NewCategoryRepository(log, mockDB)

		// Call DeleteCategory
		err := repo.DeleteCategory(tc.CategoryID)

		// Assertions
		assert.NoError(t, err)
		mockDB.AssertExpectations(t)
	}
}
