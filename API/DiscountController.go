package API

import (
	"net/http"
	"strconv"
	"th3y3m/e-commerce-platform/DependencyInjection"
	"time"

	"github.com/gin-gonic/gin"
)

// DiscountResponse represents the structure of a discount.
type DiscountResponse struct {
	ID            string    `json:"id"`
	DiscountType  string    `json:"discount_type"`
	DiscountValue float64   `json:"discount_value"`
	StartDate     time.Time `json:"start_date"`
	EndDate       time.Time `json:"end_date"`
	Status        bool      `json:"status"`
}

// CreateDiscountRequest represents the structure of the request to create a discount.
type CreateDiscountRequest struct {
	DiscountType  string    `json:"discount_type" binding:"required"`
	DiscountValue float64   `json:"discount_value" binding:"required"`
	StartDate     time.Time `json:"start_date" binding:"required"`
	EndDate       time.Time `json:"end_date" binding:"required"`
}

// UpdateDiscountRequest represents the structure of the request to update a discount.
type UpdateDiscountRequest struct {
	DiscountType  string    `json:"discount_type" binding:"required"`
	DiscountValue float64   `json:"discount_value" binding:"required"`
	StartDate     time.Time `json:"start_date" binding:"required"`
	EndDate       time.Time `json:"end_date" binding:"required"`
}

// GetPaginatedDiscountList retrieves a paginated list of discounts.
// @Summary Get paginated list of discounts
// @Description Retrieves a list of discounts with pagination and filtering options.
// @Tags Discounts
// @Produce json
// @Param searchValue query string false "Search value"
// @Param sortBy query string false "Sort by field"
// @Param pageIndex query int false "Page index" default(1)
// @Param pageSize query int false "Page size" default(10)
// @Param status query bool false "Discount status"
// @Success 200 {array} DiscountResponse
// @Failure 500 {object} API.ErrorResponse
// @Router /discounts/paginated [get]
func GetPaginatedDiscountList(c *gin.Context) {
	module := DependencyInjection.NewDiscountServiceProvider()
	searchValue := c.DefaultQuery("searchValue", "")
	sortBy := c.DefaultQuery("sortBy", "")
	pageIndex, _ := strconv.Atoi(c.DefaultQuery("pageIndex", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	status, _ := strconv.ParseBool(c.DefaultQuery("status", ""))

	discounts, err := module.GetPaginatedDiscountList(searchValue, sortBy, pageIndex, pageSize, &status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, discounts)
}

// GetAllDiscounts retrieves all discounts.
// @Summary Get all discounts
// @Description Retrieves a list of all discounts.
// @Tags Discounts
// @Produce json
// @Success 200 {array} DiscountResponse
// @Failure 500 {object} API.ErrorResponse
// @Router /discounts [get]
func GetAllDiscounts(c *gin.Context) {
	module := DependencyInjection.NewDiscountServiceProvider()
	discounts, err := module.GetAllDiscounts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, discounts)
}

// GetDiscountByID retrieves a discount by its ID.
// @Summary Get discount by ID
// @Description Retrieves a discount by providing the discount ID.
// @Tags Discounts
// @Produce json
// @Param id path string true "Discount ID"
// @Success 200 {object} DiscountResponse
// @Failure 500 {object} API.ErrorResponse
// @Router /discounts/{id} [get]
func GetDiscountByID(c *gin.Context) {
	module := DependencyInjection.NewDiscountServiceProvider()
	id := c.Param("id")
	discount, err := module.GetDiscountByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, discount)
}

// CreateDiscount creates a new discount.
// @Summary Create a new discount
// @Description Creates a new discount by providing discount details.
// @Tags Discounts
// @Accept json
// @Produce json
// @Param discount body API.CreateDiscountRequest true "Discount details"
// @Success 200 {object} API.SuccessResponse
// @Failure 400 {object} API.ErrorResponse
// @Failure 500 {object} API.ErrorResponse
// @Router /discounts [post]
func CreateDiscount(c *gin.Context) {
	module := DependencyInjection.NewDiscountServiceProvider()
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

	err := module.CreateDiscount(discount.DiscountType, discount.DiscountValue, discount.StartDate, discount.EndDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Discount created successfully"})
}

// UpdateDiscount updates an existing discount by its ID.
// @Summary Update discount by ID
// @Description Updates a discount by providing the discount ID and new discount details.
// @Tags Discounts
// @Accept json
// @Produce json
// @Param id path string true "Discount ID"
// @Param discount body API.UpdateDiscountRequest true "Updated discount details"
// @Success 200 {object} API.SuccessResponse
// @Failure 400 {object} API.ErrorResponse
// @Failure 500 {object} API.ErrorResponse
// @Router /discounts/{id} [put]
func UpdateDiscount(c *gin.Context) {
	module := DependencyInjection.NewDiscountServiceProvider()
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

	err := module.UpdateDiscount(id, discount.DiscountType, discount.DiscountValue, discount.StartDate, discount.EndDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Discount updated successfully"})
}

// DeleteDiscount deletes a discount by its ID.
// @Summary Delete discount by ID
// @Description Deletes a discount by providing the discount ID.
// @Tags Discounts
// @Produce json
// @Param id path string true "Discount ID"
// @Success 200 {object} API.SuccessResponse
// @Failure 500 {object} API.ErrorResponse
// @Router /discounts/{id} [delete]
func DeleteDiscount(c *gin.Context) {
	module := DependencyInjection.NewDiscountServiceProvider()
	id := c.Param("id")
	err := module.DeleteDiscount(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Discount deleted successfully"})
}
