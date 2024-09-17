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
	cartItems, err := Repositories.GetCartItemByCartID(cart.CartID)
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

func RemoveProductFromShoppingCart(userID, productID string, quantity int) error {
	// Retrieve the shopping cart
	cart, err := Repositories.GetUserShoppingCart(userID)
	if err != nil {
		return err
	}

	// Retrieve the cart items
	cartItems, err := Repositories.GetCartItemByCartID(cart.CartID)
	if err != nil {
		return err
	}

	// Create a map to track product quantities
	productList := make(map[string]int)
	for _, item := range cartItems {
		productList[item.ProductID] = item.Quantity
	}

	// Remove the product if it exists
	if val, ok := productList[productID]; ok {
		if val > quantity {
			productList[productID] = val - quantity
		} else {
			delete(productList, productID)
		}
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
	cartItems, err := Repositories.GetCartItemByCartID(cart.CartID)
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

func NumberOfItemsInCart(userID string) (int, error) {
	// Retrieve the shopping cart
	cart, err := Repositories.GetUserShoppingCart(userID)
	if err != nil {
		return 0, err
	}

	// Retrieve the cart items
	cartItems, err := Repositories.GetCartItemByCartID(cart.CartID)
	if err != nil {
		return 0, err
	}

	// Calculate the total number of items
	count := 0
	for _, item := range cartItems {
		count += item.Quantity
	}

	return count, nil
}

// Store the shopping cart in a cookie

func DeleteUnitItem(w http.ResponseWriter, r *http.Request, productId string, userId string) error {
	savedCart, err := r.Cookie("Cart_" + userId)
	if err == nil && savedCart != nil {
		// Retrieve cart items from the cookie (map[string]CartItem)
		cartItems, err := Util.GetCartFromCookie(savedCart.Value)
		if err != nil {
			return fmt.Errorf("error removing item from cart: %w", err)
		}

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
		strItemsInCart, err := Util.ConvertCartToString(cartItemSlice)
		if err != nil {
			return fmt.Errorf("error removing item from cart: %w", err)
		}

		err = Util.SaveCartToCookie(w, strItemsInCart, userId)
		if err != nil {
			return fmt.Errorf("error removing item from cart: %w", err)
		}
	} else {
		return fmt.Errorf("error removing item from cart: cookie not found or empty")
	}
	return nil
}

// RemoveFromCart removes a product from the cart.
func RemoveFromCart(w http.ResponseWriter, r *http.Request, productId string, userId string) error {
	savedCart, err := r.Cookie("Cart_" + userId)
	if err == nil && savedCart != nil {
		cartItems, err := Util.GetCartFromCookie(savedCart.Value)
		if err != nil {
			return fmt.Errorf("error removing item from cart: %w", err)
		}
		delete(cartItems, productId)

		var cartItemSlice []BusinessObjects.Item
		for _, item := range cartItems {
			cartItemSlice = append(cartItemSlice, item)
		}

		strItemsInCart, err := Util.ConvertCartToString(cartItemSlice)
		if err != nil {
			return fmt.Errorf("error removing item from cart: %w", err)
		}
		Util.SaveCartToCookie(w, strItemsInCart, userId)
	} else {
		fmt.Println("Error removing item from cart: cookie not found or empty")
	}
	return nil
}

// GetCart retrieves the cart items for a user.
func GetCart(r *http.Request, userId string) ([]BusinessObjects.Item, error) {
	var savedCart string

	cartCookie, err := r.Cookie("Cart_" + userId)
	if err == nil {
		savedCart = cartCookie.Value
	}

	if savedCart != "" {
		cart, err := Util.GetCartFromCookie(savedCart)
		if err != nil {
			return []BusinessObjects.Item{}, fmt.Errorf("error getting cart: %w", err)
		}
		var cartItemSlice []BusinessObjects.Item
		for _, item := range cart {
			cartItemSlice = append(cartItemSlice, item)
		}
		return cartItemSlice, nil
	}

	return []BusinessObjects.Item{}, nil
}

// DeleteCartInCookie removes the cart cookie for the user.
func DeleteCartInCookie(w http.ResponseWriter, userId string) error {
	err := Util.DeleteCartToCookie(w, userId)
	if err != nil {
		return fmt.Errorf("error deleting cart in cookie: %w", err)
	}
	return nil
}

// NumberOfItemsInCart returns the number of items in the user's cart.
func NumberOfItemsInCartCookie(r *http.Request, userId string) (int, error) {
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
		cartItems, err := Util.GetCartFromCookie(savedCart)
		if err != nil {
			return 0, fmt.Errorf("error getting number of items in cart: %w", err)
		}
		count := 0
		for _, item := range cartItems {
			count += item.Quantity
		}
		return count, nil
	}

	return 0, nil
}

// SaveCartToCookie adds or updates a product in the cart, then saves it to a cookie.
func SaveCartToCookieHandler(w http.ResponseWriter, r *http.Request, productId string, userId string) error {

	cartItems := make(map[string]BusinessObjects.Item)
	savedCart, err := r.Cookie("Cart_" + userId)
	if err == nil && savedCart != nil {
		cartItems, err = Util.GetCartFromCookie(savedCart.Value)
		if err != nil {
			return fmt.Errorf("error saving cart to cookie: %w", err)
		}
	} else {
		fmt.Println("Error saving cart to cookie: cookie not found or empty")
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

	strItemsInCart, err := Util.ConvertCartToString(cartItemSlice)
	if err != nil {
		return fmt.Errorf("error saving cart to cookie: %w", err)
	}
	err = Util.SaveCartToCookie(w, strItemsInCart, userId)
	if err != nil {
		return fmt.Errorf("error saving cart to cookie: %w", err)
	}
	return nil
}
