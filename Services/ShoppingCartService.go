package Services

import (
	"fmt"
	"net/http"
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Repositories"
	"th3y3m/e-commerce-platform/Util"
	"time"
)

func GetPaginatedShoppingCartList(sortBy, cartID, userID string, pageIndex, pageSize int, status *bool) (Util.PaginatedList[BusinessObjects.ShoppingCart], error) {
	return Repositories.GetPaginatedShoppingCartList(sortBy, cartID, userID, pageIndex, pageSize, status)
}

func GetAllShoppingCarts() ([]BusinessObjects.ShoppingCart, error) {
	return Repositories.GetAllShoppingCarts()
}

func GetShoppingCartByID(id string) (BusinessObjects.ShoppingCart, error) {
	return Repositories.GetShoppingCartByID(id)
}

func UpdateShoppingCartStatus(cartID string, status bool) error {
	return Repositories.UpdateShoppingCartStatus(cartID, status)
}

func CreateShoppingCart(userID string) (BusinessObjects.ShoppingCart, error) {
	cart := BusinessObjects.ShoppingCart{
		CartID:    "CART" + Util.GenerateID(10),
		UserID:    userID,
		Status:    true,
		CreatedAt: time.Now(),
	}

	newCart, err := Repositories.CreateShoppingCart(cart)
	if err != nil {
		return BusinessObjects.ShoppingCart{}, err
	}

	return newCart, nil
}

func UpdateShoppingCart(cart BusinessObjects.ShoppingCart) error {
	return Repositories.UpdateShoppingCart(cart)
}

func DeleteShoppingCart(id string) error {
	return Repositories.DeleteShoppingCart(id)
}

func GetUserShoppingCart(userID string) (BusinessObjects.ShoppingCart, error) {
	return Repositories.GetUserShoppingCart(userID)
}

func AddProductToShoppingCart(userID, productID string, quantity int) error {
	// Retrieve or create the shopping cart
	cart, err := Repositories.GetUserShoppingCart(userID)
	if err != nil {
		return err
	}

	// Retrieve the cart items
	cartItems, err := Repositories.GetCartItemByID(cart.CartID)
	if err != nil {
		return err
	}

	// Create a map to track product quantities
	productList := make(map[string]int)
	for _, item := range cartItems {
		productList[item.ProductID] = item.Quantity
	}

	// Update the quantity if the product exists, otherwise add it
	if val, ok := productList[productID]; ok {
		productList[productID] = val + quantity
	} else {
		productList[productID] = quantity
	}

	// Update or create the cart item
	cartItem := BusinessObjects.CartItem{
		CartID:    cart.CartID,
		ProductID: productID,
		Quantity:  productList[productID],
	}
	err = Repositories.UpdateOrCreateCartItem(cartItem)
	if err != nil {
		return err
	}

	return nil
}

func RemoveProductFromShoppingCart(userID, productID string) error {
	// Retrieve the shopping cart
	cart, err := Repositories.GetUserShoppingCart(userID)
	if err != nil {
		return err
	}

	// Retrieve the cart items
	cartItems, err := Repositories.GetCartItemByID(cart.CartID)
	if err != nil {
		return err
	}

	// Create a map to track product quantities
	productList := make(map[string]int)
	for _, item := range cartItems {
		productList[item.ProductID] = item.Quantity
	}

	// Remove the product if it exists
	if _, ok := productList[productID]; ok {
		delete(productList, productID)
	}

	// Update the cart items
	for _, item := range cartItems {
		if _, ok := productList[item.ProductID]; ok {
			cartItem := BusinessObjects.CartItem{
				CartID:    cart.CartID,
				ProductID: item.ProductID,
				Quantity:  productList[item.ProductID],
			}
			err = Repositories.UpdateOrCreateCartItem(cartItem)
			if err != nil {
				return err
			}
		} else {
			err = Repositories.DeleteCartItem(item.CartID, item.ProductID)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func ClearShoppingCart(userID string) error {
	// Retrieve the shopping cart
	cart, err := Repositories.GetUserShoppingCart(userID)
	if err != nil {
		return err
	}

	// Retrieve the cart items
	cartItems, err := Repositories.GetCartItemByID(cart.CartID)
	if err != nil {
		return err
	}

	// Delete all cart items
	for _, item := range cartItems {
		err = Repositories.DeleteCartItem(item.CartID, item.ProductID)
		if err != nil {
			return err
		}
	}

	return nil
}

// Store the shopping cart in a cookie

func DeleteUnitItem(w http.ResponseWriter, r *http.Request, productId string, userId string) {
	savedCart, err := r.Cookie("Cart_" + userId)
	if err == nil && savedCart != nil {
		// Retrieve cart items from the cookie (map[string]CartItem)
		cartItems := Util.GetCartFromCookie(savedCart.Value)

		// Check if the item exists in the cart
		if item, exists := cartItems[productId]; exists {
			item.Quantity-- // Decrease the quantity

			// If quantity is zero or less, remove the item from the cart
			if item.Quantity <= 0 {
				delete(cartItems, productId)
			} else {
				// Otherwise, reassign the updated item back to the map
				cartItems[productId] = item
			}
		}

		// Convert the map of CartItems to a slice of CartItem
		var cartItemSlice []BusinessObjects.Item
		for _, item := range cartItems {
			cartItemSlice = append(cartItemSlice, item)
		}

		// Convert the updated cart to a string and save it to the cookie
		strItemsInCart := Util.ConvertCartToString(cartItemSlice)
		Util.SaveCartToCookie(w, strItemsInCart, userId)
	} else {
		// Log the error if the cookie is not found or empty
		fmt.Println("Error deleting unit item: cookie not found or empty")
	}
}

// RemoveFromCart removes a product from the cart.
func RemoveFromCart(w http.ResponseWriter, r *http.Request, productId string, userId string) {
	savedCart, err := r.Cookie("Cart_" + userId)
	if err == nil && savedCart != nil {
		cartItems := Util.GetCartFromCookie(savedCart.Value)
		delete(cartItems, productId)

		var cartItemSlice []BusinessObjects.Item
		for _, item := range cartItems {
			cartItemSlice = append(cartItemSlice, item)
		}

		strItemsInCart := Util.ConvertCartToString(cartItemSlice)
		Util.SaveCartToCookie(w, strItemsInCart, userId)
	} else {
		fmt.Println("Error removing item from cart: cookie not found or empty")
	}
}

// GetCart retrieves the cart items for a user.
func GetCart(r *http.Request, userId string) []BusinessObjects.Item {
	var savedCart string

	cartCookie, err := r.Cookie("Cart_" + userId)
	if err == nil {
		savedCart = cartCookie.Value
	}

	if savedCart != "" {
		cart := Util.GetCartFromCookie(savedCart)
		var cartItemSlice []BusinessObjects.Item
		for _, item := range cart {
			cartItemSlice = append(cartItemSlice, item)
		}
		return cartItemSlice
	}

	return []BusinessObjects.Item{}
}

// DeleteCartInCookie removes the cart cookie for the user.
func DeleteCartInCookie(w http.ResponseWriter, userId string) {
	Util.DeleteCartToCookie(w, userId)
}

// NumberOfItemsInCart returns the number of items in the user's cart.
func NumberOfItemsInCart(r *http.Request, userId string) int {
	var savedCart string
	if userId == "" {
		cartCookie, err := r.Cookie("Cart")
		if err == nil {
			savedCart = cartCookie.Value
		}
	} else {
		cartCookie, err := r.Cookie("Cart_" + userId)
		if err == nil {
			savedCart = cartCookie.Value
		}
	}

	if savedCart != "" {
		cartItems := Util.GetCartFromCookie(savedCart)
		count := 0
		for _, item := range cartItems {
			count += item.Quantity
		}
		return count
	}

	return 0
}

// SaveCartToCookie adds or updates a product in the cart, then saves it to a cookie.
func SaveCartToCookieHandler(w http.ResponseWriter, r *http.Request, productId string, userId string) {

	cartItems := make(map[string]BusinessObjects.Item)
	savedCart, err := r.Cookie("Cart_" + userId)
	if err == nil && savedCart != nil {
		cartItems = Util.GetCartFromCookie(savedCart.Value)
	}

	item, exists := cartItems[productId]
	if !exists {
		item = BusinessObjects.Item{
			ProductID: productId,
			Quantity:  1,
		}
	} else {
		item.Quantity++
	}
	cartItems[productId] = item

	var cartItemSlice []BusinessObjects.Item
	for _, item := range cartItems {
		cartItemSlice = append(cartItemSlice, item)
	}

	strItemsInCart := Util.ConvertCartToString(cartItemSlice)
	Util.SaveCartToCookie(w, strItemsInCart, userId)
}
