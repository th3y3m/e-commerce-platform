package API

import (
	"net/http"
	"th3y3m/e-commerce-platform/Services"

	"github.com/gin-gonic/gin"
)

type CartItem struct {
	CartID    string `json:"cart_id"`
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

// GetCartItemByCartID retrieves cart items by CartID.
// @Summary Get cart items by cart ID
// @Description Retrieves all items in the cart by providing the cart ID.
// @Tags Cart
// @Produce json
// @Param cartID path string true "Cart ID"
// @Success 200 {object} []CartItem
// @Failure 500 {object} API.ErrorResponse
// @Router /cartItems/{cartID} [get]
func GetCartItemByCartID(c *gin.Context) {
	cartID := c.Param("cartID")
	cartItems, err := Services.GetCartItemByCartID(cartID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cartItems)
}
