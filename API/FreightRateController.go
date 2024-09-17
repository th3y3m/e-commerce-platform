package API

import (
	"net/http"
	"th3y3m/e-commerce-platform/Services"

	"github.com/gin-gonic/gin"
)

func GetAllFreightRates(c *gin.Context) {
	freightRates, err := Services.GetAllFreightRates()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"freightRates": freightRates})
}

func GetFreightRateByID(c *gin.Context) {

	id := c.Param("id")
	freightRate, err := Services.GetFreightRateByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"freight rate": freightRate})
}

func CreateFreightRate(c *gin.Context) {
	var freightRate struct {
		CourierID     string  `json:"courier_id" binding:"required"`
		DistanceMinKM int     `json:"distance_min_km" binding:"required"`
		DistanceMaxKM int     `json:"distance_max_km" binding:"required"`
		CostPerKM     float64 `json:"cost_per_km" binding:"required"`
		Status        bool    `json:"status" binding:"required"`
	}
	if err := c.ShouldBindJSON(&freightRate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := Services.CreateFreightRate(freightRate.CourierID, freightRate.DistanceMinKM, freightRate.DistanceMaxKM, freightRate.CostPerKM)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

func UpdateFreightRate(c *gin.Context) {
	var freightRate struct {
		CourierID     string  `json:"courier_id" binding:"required"`
		DistanceMinKM int     `json:"distance_min_km" binding:"required"`
		DistanceMaxKM int     `json:"distance_max_km" binding:"required"`
		CostPerKM     float64 `json:"cost_per_km" binding:"required"`
		Status        bool    `json:"status" binding:"required"`
	}
	if err := c.ShouldBindJSON(&freightRate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	err := Services.UpdateFreightRate(id, freightRate.CourierID, freightRate.DistanceMinKM, freightRate.DistanceMaxKM, freightRate.CostPerKM, freightRate.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

func DeleteFreightRate(c *gin.Context) {
	id := c.Param("id")
	err := Services.DeleteFreightRate(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
