package API

import (
	"net/http"
	"th3y3m/e-commerce-platform/Services"

	"github.com/gin-gonic/gin"
)

func GetAllCategories(c *gin.Context) {
	categories, err := Services.GetAllCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"categories": categories})
}

func GetCategoryByID(c *gin.Context) {
	id := c.Param("id")
	category, err := Services.GetCategoryByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"category": category})
}

func CreateCategory(c *gin.Context) {
	var info struct {
		CategoryName string `json:"category_name" binding:"required"`
	}
	if err := c.ShouldBindJSON(&info); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := Services.CreateCategory(info.CategoryName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Category created successfully"})
}

func UpdateCategory(c *gin.Context) {
	id := c.Param("id")

	var info struct {
		CategoryName string `json:"category_name" binding:"required"`
	}
	if err := c.ShouldBindJSON(&info); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := Services.UpdateCategory(id, info.CategoryName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Category updated successfully"})
}

func DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	err := Services.DeleteCategory(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}
