package Services

import (
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Interface"
	"th3y3m/e-commerce-platform/Util"
)

type CategoryService struct {
	categoryRepository Interface.ICategoryRepository
}

func NewCategoryService(categoryRepository Interface.ICategoryRepository) *CategoryService {
	return &CategoryService{categoryRepository}
}

func (c *CategoryService) GetPaginatedCategoryList(searchValue, sortBy string, pageIndex, pageSize int, status *bool) (Util.PaginatedList[BusinessObjects.Category], error) {
	return c.categoryRepository.GetPaginatedCategoryList(searchValue, sortBy, pageIndex, pageSize, status)
}

func (c *CategoryService) GetAllCategories() ([]BusinessObjects.Category, error) {
	return c.categoryRepository.GetAllCategories()
}

func (c *CategoryService) GetCategoryByID(id string) (BusinessObjects.Category, error) {
	return c.categoryRepository.GetCategoryByID(id)
}

func (c *CategoryService) CreateCategory(CategoryName string) error {
	category := BusinessObjects.Category{
		CategoryID:   "CAT" + Util.GenerateID(10),
		CategoryName: CategoryName,
	}

	err := c.categoryRepository.CreateCategory(category)
	if err != nil {
		return err
	}

	return nil
}

func (c *CategoryService) UpdateCategory(categoryID, CategoryName string) error {
	category, err := c.categoryRepository.GetCategoryByID(categoryID)
	if err != nil {
		return err
	}

	category.CategoryName = CategoryName

	return c.categoryRepository.UpdateCategory(category)
}

func (c *CategoryService) DeleteCategory(id string) error {
	return c.categoryRepository.DeleteCategory(id)
}
