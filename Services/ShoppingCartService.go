package Services

import (
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

func CreateCart(userID string) (BusinessObjects.ShoppingCart, error) {
	cart, err := Repositories.GetUserShoppingCart(userID)
	if err != nil {
		return BusinessObjects.ShoppingCart{}, err
	}

	if cart.CartID == "" {
		return CreateShoppingCart(userID)
	}

	return cart, nil
}
