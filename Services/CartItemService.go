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

func GetCartItemByID(id string) (BusinessObjects.CartItem, error) {
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
