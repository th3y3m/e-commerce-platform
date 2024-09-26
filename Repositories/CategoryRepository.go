package Repositories

import (
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Interface"
	"th3y3m/e-commerce-platform/Util"

	"github.com/sirupsen/logrus"
)

type CategoryRepository struct {
	log *logrus.Logger
}

func NewCategoryRepository(log *logrus.Logger) Interface.ICategoryRepository {
	return &CategoryRepository{log}
}

func (c *CategoryRepository) GetPaginatedategoryList(searchValue, sortBy string, pageIndex, pageSize int, status *bool) (Util.PaginatedList[BusinessObjects.Category], error) {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return Util.PaginatedList[BusinessObjects.Category]{}, err
	}

	var categories []BusinessObjects.Category
	query := db.Model(&BusinessObjects.Category{})
	if searchValue != "" {
		query = query.Where("category_name LIKE ?", "%"+searchValue+"%")
	}
	if status != nil {
		query = query.Where("status = ?", *status)
	}

	switch sortBy {
	case "category_name_asc":
		query = query.Order("category_name ASC")
	case "category_name_desc":
		query = query.Order("category_name DESC")
	default:
		query = query.Order("category_name ASC")
	}

	if err := query.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&categories).Error; err != nil {
		return Util.PaginatedList[BusinessObjects.Category]{}, err
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return Util.PaginatedList[BusinessObjects.Category]{}, err
	}

	return Util.NewPaginatedList(categories, total, pageIndex, pageSize), nil
}

// GetAllCategories retrieves all categories from the database
func (c *CategoryRepository) GetAllCategories() ([]BusinessObjects.Category, error) {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return nil, err
	}

	var categories []BusinessObjects.Category
	if err := db.Find(&categories).Error; err != nil {
		return nil, err
	}

	return categories, nil
}

// GetCategoryByID retrieves a category by its ID
func (c *CategoryRepository) GetCategoryByID(categoryID string) (BusinessObjects.Category, error) {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return BusinessObjects.Category{}, err
	}

	var category BusinessObjects.Category
	if err := db.First(&category, "category_id = ?", categoryID).Error; err != nil {
		return BusinessObjects.Category{}, err
	}

	return category, nil
}

// CreateCategory adds a new category to the database
func (c *CategoryRepository) CreateCategory(category BusinessObjects.Category) error {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return err
	}

	if err := db.Create(&category).Error; err != nil {
		return err
	}

	return nil
}

// UpdateCategory updates an existing category
func (c *CategoryRepository) UpdateCategory(category BusinessObjects.Category) error {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return err
	}

	if err := db.Save(&category).Error; err != nil {
		return err
	}

	return nil
}

// DeleteCategory removes a category from the database by its ID
func (c *CategoryRepository) DeleteCategory(categoryID string) error {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return err
	}

	if err := db.Delete(&BusinessObjects.Category{}, "category_id = ?", categoryID).Error; err != nil {
		return err
	}

	return nil
}
