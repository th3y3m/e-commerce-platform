package API

import (
	"net/http"
	"strconv"
	"th3y3m/e-commerce-platform/DependencyInjection"

	"github.com/gin-gonic/gin"
)

func CreateMoMoUrl(c *gin.Context) {
	MoMoConfig := DependencyInjection.NewMoMoServiceProvider()

	amount, _ := strconv.ParseFloat(c.Param("amount"), 64)
	orderID := c.Param("orderID")
	paymentUrl, err := MoMoConfig.CreateMoMoUrl(amount, orderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"paymentUrl": paymentUrl})
}

func ValidateMoMoResponse(c *gin.Context) {
	MoMoConfig := DependencyInjection.NewMoMoServiceProvider()

	queryParams := c.Request.URL.Query()
	res, err := MoMoConfig.ValidateMoMoResponse(queryParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": &res})
}
