package API

import (
	"net/http"
	"th3y3m/e-commerce-platform/Services"

	"github.com/gin-gonic/gin"
)

func GetAllCouriers(c *gin.Context) {
	couriers, err := Services.GetAllCouriers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"couriers": couriers})
}

func GetCourierByID(c *gin.Context) {
	id := c.Param("id")
	courier, err := Services.GetCourierByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"courier": courier})
}

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

func DeleteCourier(c *gin.Context) {
	id := c.Param("id")
	err := Services.DeleteCourier(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Courier deleted successfully"})
}
