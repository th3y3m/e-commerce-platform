package API

import (
	"net/http"
	"strconv"
	"th3y3m/e-commerce-platform/Services"

	"github.com/gin-gonic/gin"
)

func GetShoppingCartByID(c *gin.Context) {
	cartID := c.Param("cartID")
	cart, err := Services.GetShoppingCartByID(cartID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"cart": cart})
}

func GetUserShoppingCart(c *gin.Context) {
	userID := c.Param("userID")
	cart, err := Services.GetUserShoppingCart(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"cart": cart})
}

func AddProductToCart(c *gin.Context) {
	userID := c.DefaultQuery("userID", "")
	productID := c.DefaultQuery("productID", "")
	quantity, _ := strconv.Atoi(c.DefaultQuery("quantity", "1"))
	err := Services.AddProductToShoppingCart(userID, productID, quantity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product added to cart successfully"})
}

func RemoveProductFromCart(c *gin.Context) {
	cartID := c.Param("cartID")
	productID := c.Param("productID")
	quantity, _ := strconv.Atoi(c.DefaultQuery("quantity", "1"))
	err := Services.RemoveProductFromShoppingCart(cartID, productID, quantity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product removed from cart successfully"})
}

func ClearShoppingCart(c *gin.Context) {
	cartID := c.Param("cartID")
	err := Services.ClearShoppingCart(cartID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Shopping cart cleared successfully"})
}

func NumberOfItemsInCart(c *gin.Context) {
	cartID := c.Param("cartID")
	items, err := Services.NumberOfItemsInCart(cartID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"items": items})
}

//Cookie

func DeleteUnitItem(c *gin.Context) {
	productID := c.Param("productID")
	userID := c.Param("userID")
	err := Services.DeleteUnitItem(c.Writer, c.Request, productID, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Unit item deleted successfully"})
}

func RemoveFromCart(c *gin.Context) {
	productID := c.Param("productID")
	userID := c.Param("userID")
	err := Services.RemoveFromCart(c.Writer, c.Request, productID, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product removed from cart successfully"})
}

func GetCartItems(c *gin.Context) {
	userID := c.Param("userID")
	items, err := Services.GetCart(c.Request, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"items": items})
}

func DeleteCartInCookie(c *gin.Context) {
	userId := c.Param("userId")
	err := Services.DeleteCartInCookie(c.Writer, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Cart deleted successfully"})
}

func NumberOfItemsInCartCookie(c *gin.Context) {
	userId := c.Param("userId")
	items, err := Services.NumberOfItemsInCartCookie(c.Request, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"items": items})
}

func SaveCartToCookieHandler(c *gin.Context) {
	userId := c.Param("userId")
	productId := c.Param("productId")
	err := Services.SaveCartToCookieHandler(c.Writer, c.Request, productId, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Cart saved to cookie successfully"})
}
