package API

import (
	"net/http"
	"th3y3m/e-commerce-platform/Services"

	"github.com/gin-gonic/gin"
)

func GetShoppingCartByID(c *gin.Context) {
	cartID := c.Param("id")

	cart, err := Services.GetShoppingCartByID(cartID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"cart": cart})
}

func GetUserShoppingCart(c *gin.Context) {
	userID := c.Param("id")

	cart, err := Services.GetUserShoppingCart(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"cart": cart})
}

func AddProductToCart(c *gin.Context) {
	var item struct {
		UserID    string `json:"userID" binding:"required"`
		ProductID string `json:"productID" binding:"required"`
		Quantity  int    `json:"quantity" binding:"required"`
	}

	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := Services.AddProductToShoppingCart(item.UserID, item.ProductID, item.Quantity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product added to cart successfully"})
}

func RemoveProductFromCart(c *gin.Context) {
	var item struct {
		UserID    string `json:"userID" binding:"required"`
		ProductID string `json:"productID" binding:"required"`
		Quantity  int    `json:"quantity" binding:"required"`
	}

	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := Services.RemoveProductFromShoppingCart(item.UserID, item.ProductID, item.Quantity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product removed from cart successfully"})
}

func ClearShoppingCart(c *gin.Context) {
	userID := c.Param("id")

	err := Services.ClearShoppingCart(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Shopping cart cleared successfully"})
}

func NumberOfItemsInCart(c *gin.Context) {
	userID := c.Param("id")

	items, err := Services.NumberOfItemsInCart(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"items": items})
}

// Cookie handlers...

func DeleteUnitItem(c *gin.Context) {
	productID := c.Query("productID")
	userID := c.Query("userID")

	err := Services.DeleteUnitItem(c.Writer, c.Request, productID, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Unit item deleted successfully"})
}

func RemoveFromCart(c *gin.Context) {
	productID := c.Query("productID")
	userID := c.Query("userID")

	err := Services.RemoveFromCart(c.Writer, c.Request, productID, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product removed from cart successfully"})
}

func GetCartItems(c *gin.Context) {
	userID := c.Param("id")

	items, err := Services.GetCart(c.Request, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"items": items})
}

func DeleteCartInCookie(c *gin.Context) {
	userID := c.Param("id")

	err := Services.DeleteCartInCookie(c.Writer, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Cart deleted successfully"})
}

func NumberOfItemsInCartCookie(c *gin.Context) {
	userID := c.Param("id")

	items, err := Services.NumberOfItemsInCartCookie(c.Request, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"items": items})
}

func SaveCartToCookieHandler(c *gin.Context) {
	userID := c.Query("userID")
	productID := c.Query("productID")

	err := Services.SaveCartToCookieHandler(c.Writer, c.Request, productID, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Cart saved to cookie successfully"})
}
