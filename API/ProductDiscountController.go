package API

import (
	"net/http"
	"th3y3m/e-commerce-platform/DependencyInjection"

	"github.com/gin-gonic/gin"
)

func GetProductDiscountByID(c *gin.Context) {
	productID := c.Param("productID")
	module := DependencyInjection.NewProductDiscountServiceProvider()
	rate, err := module.GetProductDiscountByID(productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, rate)
}
