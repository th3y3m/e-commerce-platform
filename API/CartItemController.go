package API

import (
	"net/http"
	"th3y3m/e-commerce-platform/Services"

	"github.com/gin-gonic/gin"
)

func GetCartItemByCartID(c *gin.Context) {
	cartID := c.Param("cartID")
	cartItems, err := Services.GetCartItemByCartID(cartID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cartItems)
}
