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

func GetCartItemByCartID(id string) ([]BusinessObjects.CartItem, error) {
	return Repositories.GetCartItemByCartID(id)
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

func DeleteCartItem(cartID, productID string) error {
	return Repositories.DeleteCartItem(cartID, productID)
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
