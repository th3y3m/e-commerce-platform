package API

import (
	"net/http"
	"strconv"
	"th3y3m/e-commerce-platform/Services"
	"time"

	"github.com/gin-gonic/gin"
)

func GetPaginatedDiscountList(c *gin.Context) {
	searchValue := c.DefaultQuery("searchValue", "")
	sortBy := c.DefaultQuery("sortBy", "")
	pageIndex, _ := strconv.Atoi(c.DefaultQuery("pageIndex", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	status, _ := strconv.ParseBool(c.DefaultQuery("status", ""))

	discounts, err := Services.GetPaginatedDiscountList(searchValue, sortBy, pageIndex, pageSize, &status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"discounts": discounts})
}

func GetAllDiscounts(c *gin.Context) {
	discounts, err := Services.GetAllDiscounts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"discounts": discounts})
}

func GetDiscountByID(c *gin.Context) {
	id := c.Param("id")
	discount, err := Services.GetDiscountByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"discount": discount})
}

func CreateDiscount(c *gin.Context) {
	var discount struct {
		DiscountType  string    `json:"discount_type" binding:"required"`
		DiscountValue float64   `json:"discount_value" binding:"required"`
		StartDate     time.Time `json:"start_date" binding:"required"`
		EndDate       time.Time `json:"end_date" binding:"required"`
	}
	if err := c.ShouldBindJSON(&discount); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := Services.CreateDiscount(discount.DiscountType, discount.DiscountValue, discount.StartDate, discount.EndDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Discount created successfully"})
}

func UpdateDiscount(c *gin.Context) {
	id := c.Param("id")

	var discount struct {
		DiscountType  string    `json:"discount_type" binding:"required"`
		DiscountValue float64   `json:"discount_value" binding:"required"`
		StartDate     time.Time `json:"start_date" binding:"required"`
		EndDate       time.Time `json:"end_date" binding:"required"`
	}
	if err := c.ShouldBindJSON(&discount); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := Services.UpdateDiscount(id, discount.DiscountType, discount.DiscountValue, discount.StartDate, discount.EndDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Discount updated successfully"})
}

func DeleteDiscount(c *gin.Context) {
	id := c.Param("id")
	err := Services.DeleteDiscount(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Discount deleted successfully"})
}
