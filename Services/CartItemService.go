package Services

import (
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Repositories"
	"th3y3m/e-commerce-platform/Util"
)

func GetPaginatedCartItemList(searchValue, sortBy, cartId, productId string, pageIndex, pageSize int) (Util.PaginatedList[BusinessObjects.CartItem], error) {
	return Repositories.GetPaginatedCartItemList(searchValue, sortBy, cartId, productId, pageIndex, pageSize)
}

func GetAllCartItems() ([]BusinessObjects.CartItem, error) {
	return Repositories.GetAllCartItems()
}

func GetCartItemByID(id string) ([]BusinessObjects.CartItem, error) {
	return Repositories.GetCartItemByID(id)
}

// func CreateCartItem(cartId, productId string, quantity int) error {
// 	cartItem := BusinessObjects.CartItem{
// 		CartItemID: Util.GenerateUUID(),
// 		CartID:     cartId,
// 		ProductID:  productId,
// 		Quantity:   quantity,
// 		CreatedAt:  time.Now(),
// 	}

// 	return Repositories.CreateCartItem(cartItem)
// }

func UpdateCartItem(cartItem BusinessObjects.CartItem) error {
	return Repositories.UpdateCartItem(cartItem)
}

func DeleteCartItem(id string) error {
	return Repositories.DeleteCartItem(id)
}

// func AddProductToCart(cartId, productId string, quantity int) error {
// 	cart, err := Repositories.GetShoppingCartByID(cartId)
// 	if err != nil {
// 		return err
// 	}

// }

// AddItemToCart adds an item to the shopping cart
func AddItemToCart(cartId, productId string) error {
	cartItem, err := Repositories.GetShoppingCartByID(cartId)
	if err != nil {
		return err
	}

	productList := make(map[string]int)

	for _, item := range cartItem.CartItems {
		productList[item.ProductID] = item.Quantity
	}

	if quantity, ok := productList[productId]; ok {
		// Increment the quantity if the product already exists
		productList[productId] = quantity + 1
		for i, item := range cartItem.CartItems {
			if item.ProductID == productId {
				cartItem.CartItems[i].Quantity = quantity + 1
				break
			}
		}
	} else {
		// Add a new product if it doesn't exist
		cartItem.CartItems = append(cartItem.CartItems, BusinessObjects.CartItem{CartID: cartId, ProductID: productId, Quantity: 1})
	}

	return Repositories.UpdateShoppingCart(cartItem)
}

// RemoveItemFromCart removes an item from the shopping cart
func RemoveItemFromCart(cartId, productId string) error {
	cartItem, err := Repositories.GetShoppingCartByID(cartId)
	if err != nil {
		return err
	}

	productList := make(map[string]int)

	for _, item := range cartItem.CartItems {
		productList[item.ProductID] = item.Quantity
	}

	if quantity, ok := productList[productId]; ok {
		if quantity == 1 {
			// Remove the product if the quantity is 1
			for i, item := range cartItem.CartItems {
				if item.ProductID == productId {
					cartItem.CartItems = append(cartItem.CartItems[:i], cartItem.CartItems[i+1:]...)
					break
				}
			}
		} else {
			// Decrement the quantity if the product exists
			productList[productId] = quantity - 1
			for i, item := range cartItem.CartItems {
				if item.ProductID == productId {
					cartItem.CartItems[i].Quantity = quantity - 1
					break
				}
			}
		}
	}

	return Repositories.UpdateShoppingCart(cartItem)
}
