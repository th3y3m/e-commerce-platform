package API

import (
	"net/http"
	"th3y3m/e-commerce-platform/DependencyInjection"

	"github.com/gin-gonic/gin"
)

func GetShoppingCartByID(c *gin.Context) {
	cartID := c.Param("id")

	module := DependencyInjection.NewShoppingCartServiceProvider()

	cart, err := module.GetShoppingCartByID(cartID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cart)
}

func GetUserShoppingCart(c *gin.Context) {
	userID := c.Param("id")

	module := DependencyInjection.NewShoppingCartServiceProvider()

	cart, err := module.GetUserShoppingCart(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cart)
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
	module := DependencyInjection.NewShoppingCartServiceProvider()

	err := module.AddProductToShoppingCart(item.UserID, item.ProductID, item.Quantity)
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
	module := DependencyInjection.NewShoppingCartServiceProvider()

	err := module.RemoveProductFromShoppingCart(item.UserID, item.ProductID, item.Quantity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product removed from cart successfully"})
}

func ClearShoppingCart(c *gin.Context) {
	userID := c.Param("id")
	module := DependencyInjection.NewShoppingCartServiceProvider()

	err := module.ClearShoppingCart(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Shopping cart cleared successfully"})
}

func NumberOfItemsInCart(c *gin.Context) {
	userID := c.Param("id")
	module := DependencyInjection.NewShoppingCartServiceProvider()

	items, err := module.NumberOfItemsInCart(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}

// Cookie handlers...

func DeleteUnitItem(c *gin.Context) {
	productID := c.Query("productID")
	userID := c.Query("userID")
	module := DependencyInjection.NewShoppingCartServiceProvider()

	err := module.DeleteUnitItem(c.Writer, c.Request, productID, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Unit item deleted successfully"})
}

func RemoveFromCart(c *gin.Context) {
	productID := c.Query("productID")
	userID := c.Query("userID")
	module := DependencyInjection.NewShoppingCartServiceProvider()

	err := module.RemoveFromCart(c.Writer, c.Request, productID, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product removed from cart successfully"})
}

func GetCartItems(c *gin.Context) {
	userID := c.Param("id")
	module := DependencyInjection.NewShoppingCartServiceProvider()

	items, err := module.GetCart(c.Request, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"items": items})
}

func DeleteCartInCookie(c *gin.Context) {
	userID := c.Param("id")
	module := DependencyInjection.NewShoppingCartServiceProvider()

	err := module.DeleteCartInCookie(c.Writer, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Cart deleted successfully"})
}

func NumberOfItemsInCartCookie(c *gin.Context) {
	userID := c.Param("id")
	module := DependencyInjection.NewShoppingCartServiceProvider()

	items, err := module.NumberOfItemsInCartCookie(c.Request, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}

func SaveCartToCookieHandler(c *gin.Context) {
	userID := c.Query("userID")
	productID := c.Query("productID")
	module := DependencyInjection.NewShoppingCartServiceProvider()

	err := module.SaveCartToCookieHandler(c.Writer, c.Request, productID, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Cart saved to cookie successfully"})
}
