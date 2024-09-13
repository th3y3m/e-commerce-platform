package Services

import (
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Repositories"
	"th3y3m/e-commerce-platform/Util"
)

func GetPaginatedategoryList(searchValue, sortBy string, pageIndex, pageSize int, status *bool) (Util.PaginatedList[BusinessObjects.Category], error) {
	return Repositories.GetPaginatedategoryList(searchValue, sortBy, pageIndex, pageSize, status)
}

func GetAllCategories() ([]BusinessObjects.Category, error) {
	return Repositories.GetAllCategories()
}

func GetCategoryByID(id string) (BusinessObjects.Category, error) {
	return Repositories.GetCategoryByID(id)
}

func CreateCategory(CategoryName string) error {
	category := BusinessObjects.Category{
		CategoryID:   "CAT" + Util.GenerateID(10),
		CategoryName: CategoryName,
	}

	err := Repositories.CreateCategory(category)
	if err != nil {
		return err
	}

	return nil
}

func UpdateCategory(categoryID, CategoryName string) error {
	category, err := Repositories.GetCategoryByID(categoryID)
	if err != nil {
		return err
	}

	category.CategoryName = CategoryName

	return Repositories.UpdateCategory(category)
}

func DeleteCategory(id string) error {
	return Repositories.DeleteCategory(id)
}
