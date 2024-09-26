package Interface

import (
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Util"
)

type ICartItemRepository interface {
	GetPaginatedCartItemList(searchValue, sortBy, cartId, productId string, pageIndex, pageSize int) (Util.PaginatedList[BusinessObjects.CartItem], error)
	GetAllCartItems() ([]BusinessObjects.CartItem, error)
	GetCartItemByCartID(cartItemID string) ([]BusinessObjects.CartItem, error)
	CreateCartItem(cartItem BusinessObjects.CartItem) error
	UpdateCartItem(cartItem BusinessObjects.CartItem) error
	DeleteCartItem(cartID, productID string) error
	UpdateOrCreateCartItem(cartItem BusinessObjects.CartItem) error
}

type ICartItemService interface {
	GetPaginatedCartItemList(searchValue, sortBy, cartId, productId string, pageIndex, pageSize int) (Util.PaginatedList[BusinessObjects.CartItem], error)
	GetAllCartItems() ([]BusinessObjects.CartItem, error)
	GetCartItemByCartID(cartItemID string) ([]BusinessObjects.CartItem, error)
	UpdateCartItem(cartItem BusinessObjects.CartItem) error
	DeleteCartItem(cartID, productID string) error
	RemoveItemFromCart(cartId, productId string) error
}
