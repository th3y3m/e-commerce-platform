package API

import (
	"net/http"
	"th3y3m/e-commerce-platform/DependencyInjection"

	"github.com/gin-gonic/gin"
)

func GetOrderDetailOfAOrder(c *gin.Context) {
	orderId := c.Param("orderId")
	module := DependencyInjection.NewOrderDetailServiceProvider()
	orderDetails, err := module.GetOrderDetailByID(orderId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, orderDetails)
}
