package Services

import (
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Interface"
	"th3y3m/e-commerce-platform/Util"
)

type CartItemService struct {
	cartItemRepository Interface.ICartItemRepository
	shoppingRepository Interface.IShoppingCartRepository
}

func NewCartItemService(cartItemRepository Interface.ICartItemRepository, shoppingRepository Interface.IShoppingCartRepository) Interface.ICartItemService {
	return &CartItemService{
		cartItemRepository: cartItemRepository,
		shoppingRepository: shoppingRepository,
	}
}

func (c *CartItemService) GetPaginatedCartItemList(searchValue, sortBy, cartId, productId string, pageIndex, pageSize int) (Util.PaginatedList[BusinessObjects.CartItem], error) {
	return c.cartItemRepository.GetPaginatedCartItemList(searchValue, sortBy, cartId, productId, pageIndex, pageSize)
}

func (c *CartItemService) GetAllCartItems() ([]BusinessObjects.CartItem, error) {
	return c.cartItemRepository.GetAllCartItems()
}

func (c *CartItemService) GetCartItemByCartID(id string) ([]BusinessObjects.CartItem, error) {
	return c.cartItemRepository.GetCartItemByCartID(id)
}

// func  (c *CartItemService) CreateCartItem(cartId, productId string, quantity int) error {
// 	cartItem := BusinessObjects.CartItem{
// 		CartItemID: Util.GenerateUUID(),
// 		CartID:     cartId,
// 		ProductID:  productId,
// 		Quantity:   quantity,
// 		CreatedAt:  time.Now(),
// 	}

// 	return Repositories.CreateCartItem(cartItem)
// }

func (c *CartItemService) UpdateCartItem(cartItem BusinessObjects.CartItem) error {
	return c.cartItemRepository.UpdateCartItem(cartItem)
}

func (c *CartItemService) DeleteCartItem(cartID, productID string) error {
	return c.cartItemRepository.DeleteCartItem(cartID, productID)
}

// RemoveItemFromCart removes an item from the shopping cart
func (c *CartItemService) RemoveItemFromCart(cartId, productId string) error {
	cartItem, err := c.shoppingRepository.GetShoppingCartByID(cartId)
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

	return c.shoppingRepository.UpdateShoppingCart(cartItem)
}
