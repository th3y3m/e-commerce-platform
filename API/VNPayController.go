package API

import (
	"fmt"
	"net/http"
	"strconv"
	"th3y3m/e-commerce-platform/DependencyInjection"

	"github.com/gin-gonic/gin"
)

func CreateVNPayUrl(c *gin.Context) {
	amount, _ := strconv.ParseFloat(c.Param("amount"), 64)
	orderID := c.Param("orderID")
	vnpayConfig := DependencyInjection.NewVnpayServiceProvider()
	paymentUrl, err := vnpayConfig.CreateVNPayUrl(amount, orderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"paymentUrl": paymentUrl})
}

func ValidateVNPayResponse(c *gin.Context) {
	vnpayConfig := DependencyInjection.NewVnpayServiceProvider()
	queryParams := c.Request.URL.Query()
	fmt.Print(queryParams)
	res, err := vnpayConfig.ValidateVNPayResponse(queryParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": &res})
}
