package Services

import (
	"testing"
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Util"
	"th3y3m/e-commerce-platform/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetPaginatedCategoryList_Success(t *testing.T) {
	// Arrange
	categoryRepository := &mocks.ICategoryRepository{}
	categoryService := NewCategoryService(categoryRepository)

	searchValue := ""
	sortBy := ""
	pageIndex := 1
	pageSize := 10
	status := true

	categoryRepository.On("GetPaginatedCategoryList", searchValue, sortBy, pageIndex, pageSize, &status).Return(Util.PaginatedList[BusinessObjects.Category]{}, nil)

	// Act
	_, err := categoryService.GetPaginatedCategoryList(searchValue, sortBy, pageIndex, pageSize, &status)

	// Assert
	assert.NoError(t, err)
	categoryRepository.AssertExpectations(t)
}
func TestGetPaginatedCategoryList_Error(t *testing.T) {
	// Arrange
	categoryRepository := &mocks.ICategoryRepository{}
	categoryService := NewCategoryService(categoryRepository)

	searchValue := ""
	sortBy := ""
	pageIndex := 1
	pageSize := 10
	status := true

	error := assert.AnError

	categoryRepository.On("GetPaginatedCategoryList", searchValue, sortBy, pageIndex, pageSize, &status).Return(Util.PaginatedList[BusinessObjects.Category]{}, error)

	// Act
	_, err := categoryService.GetPaginatedCategoryList(searchValue, sortBy, pageIndex, pageSize, &status)

	// Assert
	assert.Error(t, err)
	categoryRepository.AssertExpectations(t)
}

func TestGetAllCategories_Success(t *testing.T) {
	// Arrange
	categoryRepository := &mocks.ICategoryRepository{}
	categoryService := NewCategoryService(categoryRepository)

	categoryRepository.On("GetAllCategories").Return([]BusinessObjects.Category{}, nil)

	// Act
	_, err := categoryService.GetAllCategories()

	// Assert
	assert.NoError(t, err)
	categoryRepository.AssertExpectations(t)
}
func TestGetAllCategories_Error(t *testing.T) {
	// Arrange
	categoryRepository := &mocks.ICategoryRepository{}
	categoryService := NewCategoryService(categoryRepository)

	error := assert.AnError

	categoryRepository.On("GetAllCategories").Return([]BusinessObjects.Category{}, error)

	// Act
	_, err := categoryService.GetAllCategories()

	// Assert
	assert.Error(t, err)
	categoryRepository.AssertExpectations(t)
}

func TestGetCategoryByID_Success(t *testing.T) {
	// Arrange
	categoryRepository := &mocks.ICategoryRepository{}
	categoryService := NewCategoryService(categoryRepository)

	id := "123"

	categoryRepository.On("GetCategoryByID", id).Return(BusinessObjects.Category{}, nil)

	// Act
	_, err := categoryService.GetCategoryByID(id)

	// Assert
	assert.NoError(t, err)
	categoryRepository.AssertExpectations(t)
}
func TestGetCategoryByID_Error(t *testing.T) {
	// Arrange
	categoryRepository := &mocks.ICategoryRepository{}
	categoryService := NewCategoryService(categoryRepository)

	id := "123"
	error := assert.AnError

	categoryRepository.On("GetCategoryByID", id).Return(BusinessObjects.Category{}, error)

	// Act
	_, err := categoryService.GetCategoryByID(id)

	// Assert
	assert.Error(t, err)
	categoryRepository.AssertExpectations(t)
}

func TestCreateCategory_Success(t *testing.T) {
	// Arrange
	categoryRepository := &mocks.ICategoryRepository{}
	categoryService := NewCategoryService(categoryRepository)

	categoryName := "categoryName"

	categoryRepository.On("CreateCategory", mock.AnythingOfType("BusinessObjects.Category")).Return(nil)

	// Act
	err := categoryService.CreateCategory(categoryName)

	// Assert
	assert.NoError(t, err)
	categoryRepository.AssertExpectations(t)
}
func TestCreateCategory_Error(t *testing.T) {
	// Arrange
	categoryRepository := &mocks.ICategoryRepository{}
	categoryService := NewCategoryService(categoryRepository)

	categoryName := "categoryName"

	error := assert.AnError

	categoryRepository.On("CreateCategory", mock.AnythingOfType("BusinessObjects.Category")).Return(error)

	// Act
	err := categoryService.CreateCategory(categoryName)

	// Assert
	assert.Error(t, err)
	categoryRepository.AssertExpectations(t)
}

func TestUpdateCategory_Success(t *testing.T) {
	// Arrange
	categoryRepository := &mocks.ICategoryRepository{}
	categoryService := NewCategoryService(categoryRepository)

	category := BusinessObjects.Category{CategoryID: "1", CategoryName: "categoryName"}
	categoryID := "1"
	categoryName := "test"

	categoryRepository.On("GetCategoryByID", categoryID).Return(category, nil)
	categoryRepository.On("UpdateCategory", BusinessObjects.Category{CategoryID: categoryID, CategoryName: categoryName}).Return(nil)

	// Act
	err := categoryService.UpdateCategory(categoryID, categoryName)

	// Assert
	assert.NoError(t, err)
	categoryRepository.AssertExpectations(t)
}

func TestUpdateCategory_Error(t *testing.T) {
	// Arrange
	categoryRepository := &mocks.ICategoryRepository{}
	categoryService := NewCategoryService(categoryRepository)

	categoryID := "123"
	categoryName := "categoryName"
	error := assert.AnError

	categoryRepository.On("GetCategoryByID", categoryID).Return(BusinessObjects.Category{}, error)

	// Act
	err := categoryService.UpdateCategory(categoryID, categoryName)

	// Assert
	assert.Error(t, err)
	categoryRepository.AssertExpectations(t)
}

func TestDeleteCategory_Success(t *testing.T) {
	// Arrange
	categoryRepository := &mocks.ICategoryRepository{}
	categoryService := NewCategoryService(categoryRepository)

	id := "123"

	categoryRepository.On("DeleteCategory", id).Return(nil)

	// Act
	err := categoryService.DeleteCategory(id)

	// Assert
	assert.NoError(t, err)
	categoryRepository.AssertExpectations(t)
}

func TestDeleteCategory_Error(t *testing.T) {
	// Arrange
	categoryRepository := &mocks.ICategoryRepository{}
	categoryService := NewCategoryService(categoryRepository)

	id := "123"
	error := assert.AnError

	categoryRepository.On("DeleteCategory", id).Return(error)

	// Act
	err := categoryService.DeleteCategory(id)

	// Assert
	assert.Error(t, err)
	categoryRepository.AssertExpectations(t)
}
