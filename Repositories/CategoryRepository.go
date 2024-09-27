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

func (c *CategoryRepository) GetPaginatedCategoryList(searchValue, sortBy string, pageIndex, pageSize int, status *bool) (Util.PaginatedList[BusinessObjects.Category], error) {
	c.log.Infof("Fetching paginated category list with searchValue: %s, sortBy: %s, pageIndex: %d, pageSize: %d, status: %v", searchValue, sortBy, pageIndex, pageSize, status)
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		c.log.Error("Failed to connect to PostgreSQL:", err)
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
		c.log.Error("Failed to fetch paginated categories:", err)
		return Util.PaginatedList[BusinessObjects.Category]{}, err
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		c.log.Error("Failed to count categories:", err)
		return Util.PaginatedList[BusinessObjects.Category]{}, err
	}

	c.log.Infof("Successfully fetched paginated category list with total count: %d", total)
	return Util.NewPaginatedList(categories, total, pageIndex, pageSize), nil
}

// GetAllCategories retrieves all categories from the database
func (c *CategoryRepository) GetAllCategories() ([]BusinessObjects.Category, error) {
	c.log.Info("Fetching all categories")
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		c.log.Error("Failed to connect to PostgreSQL:", err)
		return nil, err
	}

	var categories []BusinessObjects.Category
	if err := db.Find(&categories).Error; err != nil {
		c.log.Error("Failed to fetch all categories:", err)
		return nil, err
	}

	c.log.Info("Successfully fetched all categories")
	return categories, nil
}

// GetCategoryByID retrieves a category by its ID
func (c *CategoryRepository) GetCategoryByID(categoryID string) (BusinessObjects.Category, error) {
	c.log.Infof("Fetching category by ID: %s", categoryID)
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		c.log.Error("Failed to connect to PostgreSQL:", err)
		return BusinessObjects.Category{}, err
	}

	var category BusinessObjects.Category
	if err := db.First(&category, "category_id = ?", categoryID).Error; err != nil {
		c.log.Error("Failed to fetch category by ID:", err)
		return BusinessObjects.Category{}, err
	}

	c.log.Infof("Successfully fetched category by ID: %s", categoryID)
	return category, nil
}

// CreateCategory adds a new category to the database
func (c *CategoryRepository) CreateCategory(category BusinessObjects.Category) error {
	c.log.Infof("Creating new category with name: %s", category.CategoryName)
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		c.log.Error("Failed to connect to PostgreSQL:", err)
		return err
	}

	if err := db.Create(&category).Error; err != nil {
		c.log.Error("Failed to create new category:", err)
		return err
	}

	c.log.Infof("Successfully created new category with name: %s", category.CategoryName)
	return nil
}

// UpdateCategory updates an existing category
func (c *CategoryRepository) UpdateCategory(category BusinessObjects.Category) error {
	c.log.Infof("Updating category with ID: %s", category.CategoryID)
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		c.log.Error("Failed to connect to PostgreSQL:", err)
		return err
	}

	if err := db.Save(&category).Error; err != nil {
		c.log.Error("Failed to update category:", err)
		return err
	}

	c.log.Infof("Successfully updated category with ID: %s", category.CategoryID)
	return nil
}

// DeleteCategory removes a category from the database by its ID
func (c *CategoryRepository) DeleteCategory(categoryID string) error {
	c.log.Infof("Deleting category with ID: %s", categoryID)
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		c.log.Error("Failed to connect to PostgreSQL:", err)
		return err
	}

	if err := db.Delete(&BusinessObjects.Category{}, "category_id = ?", categoryID).Error; err != nil {
		c.log.Error("Failed to delete category:", err)
		return err
	}

	c.log.Infof("Successfully deleted category with ID: %s", categoryID)
	return nil
}
