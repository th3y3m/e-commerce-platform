package API

// func AddToCart(c *gin.Context) {
// 	var productID string
// 	if err := c.ShouldBindJSON(&productID); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	cart, err := Services.AddItemToCart(cartItem.UserID)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	cartItem.CartID = cart.CartID
// 	newCartItem, err := Services.CreateCartItem(cartItem)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, newCartItem)
// }
