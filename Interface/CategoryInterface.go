package Interface

import (
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Util"
)

type ICategoryRepository interface {
	GetPaginatedategoryList(searchValue, sortBy string, pageIndex, pageSize int, status *bool) (Util.PaginatedList[BusinessObjects.Category], error)
	GetAllCategories() ([]BusinessObjects.Category, error)
	GetCategoryByID(categoryID string) (BusinessObjects.Category, error)
	CreateCategory(category BusinessObjects.Category) error
	UpdateCategory(category BusinessObjects.Category) error
	DeleteCategory(categoryID string) error
}
type ICategoryService interface {
	GetPaginatedategoryList(searchValue, sortBy string, pageIndex, pageSize int, status *bool) (Util.PaginatedList[BusinessObjects.Category], error)
	GetAllCategories() ([]BusinessObjects.Category, error)
	GetCategoryByID(categoryID string) (BusinessObjects.Category, error)
	CreateCategory(CategoryName string) error
	UpdateCategory(categoryID, CategoryName string) error
	DeleteCategory(categoryID string) error
}
