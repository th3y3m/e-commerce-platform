package Interface

import (
	"net/http"
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Util"
)

type IShoppingCartRepository interface {
	GetPaginatedShoppingCartList(sortBy, cartID, userID string, pageIndex, pageSize int, status *bool) (Util.PaginatedList[BusinessObjects.ShoppingCart], error)
	GetUserShoppingCart(userID string) (BusinessObjects.ShoppingCart, error)
	GetAllShoppingCarts() ([]BusinessObjects.ShoppingCart, error)
	GetShoppingCartByID(cartID string) (BusinessObjects.ShoppingCart, error)
	CreateShoppingCart(cart BusinessObjects.ShoppingCart) (BusinessObjects.ShoppingCart, error)
	UpdateShoppingCart(cart BusinessObjects.ShoppingCart) error
	DeleteShoppingCart(cartID string) error
	UpdateShoppingCartStatus(cartID string, status bool) error
}

type IShoppingCartService interface {
	GetPaginatedShoppingCartList(sortBy, cartID, userID string, pageIndex, pageSize int, status *bool) (Util.PaginatedList[BusinessObjects.ShoppingCart], error)
	GetAllShoppingCarts() ([]BusinessObjects.ShoppingCart, error)
	GetShoppingCartByID(id string) (BusinessObjects.ShoppingCart, error)
	UpdateShoppingCartStatus(cartID string, status bool) error
	CreateShoppingCart(userID string) (BusinessObjects.ShoppingCart, error)
	UpdateShoppingCart(cart BusinessObjects.ShoppingCart) error
	DeleteShoppingCart(id string) error
	GetUserShoppingCart(userID string) (BusinessObjects.ShoppingCart, error)
	AddProductToShoppingCart(userID, productID string, quantity int) error
	RemoveProductFromShoppingCart(userID, productID string, quantity int) error
	ClearShoppingCart(userID string) error
	NumberOfItemsInCart(userID string) (int, error)
	DeleteUnitItem(w http.ResponseWriter, r *http.Request, productId string, userId string) error
	RemoveFromCart(w http.ResponseWriter, r *http.Request, productId string, userId string) error
	GetCart(r *http.Request, userId string) ([]BusinessObjects.Item, error)
	DeleteCartInCookie(w http.ResponseWriter, userId string) error
	NumberOfItemsInCartCookie(r *http.Request, userId string) (int, error)
	SaveCartToCookieHandler(w http.ResponseWriter, r *http.Request, productId string, userId string) error
}
