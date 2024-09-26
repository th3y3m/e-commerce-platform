package API

import (
	"net/http"
	"th3y3m/e-commerce-platform/DependencyInjection"

	"github.com/gin-gonic/gin"
)

// Category is the structure for product categories.
type Category struct {
	CategoryID   string `json:"category_id"`
	CategoryName string `json:"category_name"`
}

// SuccessResponse is the structure for successful API responses.
type SuccessResponse struct {
	Message string `json:"message"`
}

// CategoryRequest is used when creating or updating a category.
type CategoryRequest struct {
	CategoryName string `json:"category_name" binding:"required"`
}

// GetAllCategories retrieves all categories.
// @Summary Get all categories
// @Description Retrieves all product categories.
// @Tags Categories
// @Produce json
// @Success 200 {object} []BusinessObjects.Category
// @Failure 500 {object} API.ErrorResponse
// @Router /categories [get]
func GetAllCategories(c *gin.Context) {
	module := DependencyInjection.NewCategoryServiceProvider()
	categories, err := module.GetAllCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, categories)
}

// GetCategoryByID retrieves a category by ID.
// @Summary Get category by ID
// @Description Retrieves category details by providing the category ID.
// @Tags Categories
// @Produce json
// @Param id path string true "Category ID"
// @Success 200 {object} Category
// @Failure 500 {object} API.ErrorResponse
// @Router /categories/{id} [get]
func GetCategoryByID(c *gin.Context) {
	id := c.Param("id")
	module := DependencyInjection.NewCategoryServiceProvider()

	category, err := module.GetCategoryByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, category)
}

// CreateCategory creates a new category.
// @Summary Create a new category
// @Description Creates a new product category.
// @Tags Categories
// @Accept json
// @Produce json
// @Param category body API.CategoryRequest true "Category Name"
// @Success 200 {object} API.SuccessResponse
// @Failure 500 {object} API.ErrorResponse
// @Router /categories [post]
func CreateCategory(c *gin.Context) {
	var info CategoryRequest
	module := DependencyInjection.NewCategoryServiceProvider()

	if err := c.ShouldBindJSON(&info); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := module.CreateCategory(info.CategoryName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Category created successfully"})
}

// UpdateCategory updates a category.
// @Summary Update category by ID
// @Description Updates an existing category by providing category ID and new data.
// @Tags Categories
// @Accept json
// @Produce json
// @Param id path string true "Category ID"
// @Param category body API.CategoryRequest true "Updated Category Data"
// @Success 200 {object} API.SuccessResponse
// @Failure 500 {object} API.ErrorResponse
// @Router /categories/{id} [put]
func UpdateCategory(c *gin.Context) {
	id := c.Param("id")
	module := DependencyInjection.NewCategoryServiceProvider()

	var info CategoryRequest
	if err := c.ShouldBindJSON(&info); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := module.UpdateCategory(id, info.CategoryName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Category updated successfully"})
}

// DeleteCategory deletes a category.
// @Summary Delete category by ID
// @Description Deletes an existing category by providing the category ID.
// @Tags Categories
// @Produce json
// @Param id path string true "Category ID"
// @Success 200 {object} API.SuccessResponse
// @Failure 500 {object} API.ErrorResponse
// @Router /categories/{id} [delete]
func DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	module := DependencyInjection.NewCategoryServiceProvider()

	err := module.DeleteCategory(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}
