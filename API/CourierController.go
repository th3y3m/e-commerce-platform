package API

import (
	"net/http"
	"th3y3m/e-commerce-platform/Services"

	"github.com/gin-gonic/gin"
)

// CreateCourierRequest represents the request body for creating a courier.
type CreateCourierRequest struct {
	CourierName string `json:"courier_name" binding:"required"`
}

// UpdateCourierRequest represents the request body for updating a courier.
type UpdateCourierRequest struct {
	CourierName string `json:"courier_name" binding:"required"`
	Status      bool   `json:"status" binding:"required"`
}

// CourierResponse represents the response body for courier details.
type CourierResponse struct {
	ID          string `json:"id"`
	CourierName string `json:"courier_name"`
	Status      bool   `json:"status"`
}

// GetAllCouriers retrieves all couriers.
// @Summary Get all couriers
// @Description Retrieves a list of all available couriers.
// @Tags Couriers
// @Produce json
// @Success 200 {object} []CourierResponse
// @Failure 500 {object} API.ErrorResponse
// @Router /couriers [get]
func GetAllCouriers(c *gin.Context) {
	couriers, err := Services.GetAllCouriers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, couriers)
}

// GetCourierByID retrieves a courier by its ID.
// @Summary Get courier by ID
// @Description Retrieves a specific courier by providing the courier ID.
// @Tags Couriers
// @Produce json
// @Param id path string true "Courier ID"
// @Success 200 {object} CourierResponse
// @Failure 500 {object} API.ErrorResponse
// @Router /couriers/{id} [get]
func GetCourierByID(c *gin.Context) {
	id := c.Param("id")
	courier, err := Services.GetCourierByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, courier)
}

// CreateCourier creates a new courier.
// @Summary Create a new courier
// @Description Adds a new courier by providing the courier's name.
// @Tags Couriers
// @Accept json
// @Produce json
// @Param courier body API.CreateCourierRequest true "Courier name"
// @Success 200 {object} API.SuccessResponse
// @Failure 400 {object} API.ErrorResponse
// @Failure 500 {object} API.ErrorResponse
// @Router /couriers [post]
func CreateCourier(c *gin.Context) {
	var info struct {
		CourierName string `json:"courier_name" binding:"required"`
	}
	if err := c.ShouldBindJSON(&info); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := Services.CreateCourier(info.CourierName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Courier created successfully"})
}

// UpdateCourier updates an existing courier.
// @Summary Update courier by ID
// @Description Updates an existing courier by providing courier ID, name, and status.
// @Tags Couriers
// @Accept json
// @Produce json
// @Param id path string true "Courier ID"
// @Param courier body API.UpdateCourierRequest true "Updated courier data"
// @Success 200 {object} API.SuccessResponse
// @Failure 400 {object} API.ErrorResponse
// @Failure 500 {object} API.ErrorResponse
// @Router /couriers/{id} [put]
func UpdateCourier(c *gin.Context) {
	id := c.Param("id")

	var info struct {
		CourierName string `json:"courier_name" binding:"required"`
		Status      bool   `json:"status" binding:"required"`
	}
	if err := c.ShouldBindJSON(&info); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := Services.UpdateCourier(id, info.CourierName, info.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Courier updated successfully"})
}

// DeleteCourier deletes a courier by its ID.
// @Summary Delete courier by ID
// @Description Deletes a specific courier by providing the courier ID.
// @Tags Couriers
// @Produce json
// @Param id path string true "Courier ID"
// @Success 200 {object} API.SuccessResponse
// @Failure 500 {object} API.ErrorResponse
// @Router /couriers/{id} [delete]
func DeleteCourier(c *gin.Context) {
	id := c.Param("id")
	err := Services.DeleteCourier(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Courier deleted successfully"})
}
