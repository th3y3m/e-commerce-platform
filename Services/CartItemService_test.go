package Services

import (
	"testing"
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Util"
	"th3y3m/e-commerce-platform/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetPaginatedCartItemList_Success(t *testing.T) {
	// Arrange
	cartItemRepository := &mocks.ICartItemRepository{}
	shoppingRepository := &mocks.IShoppingCartRepository{}
	cartItemService := NewCartItemService(cartItemRepository, shoppingRepository)

	searchValue := ""
	sortBy := ""
	cartId := ""
	productId := ""
	pageIndex := 1
	pageSize := 10

	cartItemRepository.On("GetPaginatedCartItemList", searchValue, sortBy, cartId, productId, pageIndex, pageSize).Return(Util.PaginatedList[BusinessObjects.CartItem]{}, nil)

	// Act
	_, err := cartItemService.GetPaginatedCartItemList(searchValue, sortBy, cartId, productId, pageIndex, pageSize)

	// Assert
	assert.NoError(t, err)
	cartItemRepository.AssertExpectations(t)
}
func TestGetPaginatedCartItemList_Error(t *testing.T) {
	// Arrange
	cartItemRepository := &mocks.ICartItemRepository{}
	shoppingRepository := &mocks.IShoppingCartRepository{}
	cartItemService := NewCartItemService(cartItemRepository, shoppingRepository)

	searchValue := ""
	sortBy := ""
	cartId := ""
	productId := ""
	pageIndex := 1
	pageSize := 10

	error := assert.AnError

	cartItemRepository.On("GetPaginatedCartItemList", searchValue, sortBy, cartId, productId, pageIndex, pageSize).Return(Util.PaginatedList[BusinessObjects.CartItem]{}, error)

	// Act
	_, err := cartItemService.GetPaginatedCartItemList(searchValue, sortBy, cartId, productId, pageIndex, pageSize)

	// Assert
	assert.Error(t, err)
	cartItemRepository.AssertExpectations(t)
}

func TestGetAllCartItems_Success(t *testing.T) {
	// Arrange
	cartItemRepository := &mocks.ICartItemRepository{}
	shoppingRepository := &mocks.IShoppingCartRepository{}
	cartItemService := NewCartItemService(cartItemRepository, shoppingRepository)

	cartItemRepository.On("GetAllCartItems").Return([]BusinessObjects.CartItem{}, nil)

	// Act
	_, err := cartItemService.GetAllCartItems()

	// Assert
	assert.NoError(t, err)
	cartItemRepository.AssertExpectations(t)
}

func TestGetCartItemByCartID_Success(t *testing.T) {
	// Arrange
	cartItemRepository := &mocks.ICartItemRepository{}
	shoppingRepository := &mocks.IShoppingCartRepository{}
	cartItemService := NewCartItemService(cartItemRepository, shoppingRepository)

	cartItemID := "cartItemID"
	cartItemRepository.On("GetCartItemByCartID", cartItemID).Return([]BusinessObjects.CartItem{}, nil)

	// Act
	_, err := cartItemService.GetCartItemByCartID(cartItemID)

	// Assert
	assert.NoError(t, err)
	cartItemRepository.AssertExpectations(t)
}
func TestGetCartItemByCartID_Error(t *testing.T) {
	// Arrange
	cartItemRepository := &mocks.ICartItemRepository{}
	shoppingRepository := &mocks.IShoppingCartRepository{}
	cartItemService := NewCartItemService(cartItemRepository, shoppingRepository)

	error := assert.AnError
	cartItemID := "cartItemID"
	cartItemRepository.On("GetCartItemByCartID", cartItemID).Return([]BusinessObjects.CartItem{}, error)

	// Act
	_, err := cartItemService.GetCartItemByCartID(cartItemID)

	// Assert
	assert.Error(t, err)
	cartItemRepository.AssertExpectations(t)
}

func TestUpdateCartItem_Success(t *testing.T) {
	// Arrange
	cartItemRepository := &mocks.ICartItemRepository{}
	shoppingRepository := &mocks.IShoppingCartRepository{}
	cartItemService := NewCartItemService(cartItemRepository, shoppingRepository)

	cartItem := BusinessObjects.CartItem{}
	cartItemRepository.On("UpdateCartItem", cartItem).Return(nil)

	// Act
	err := cartItemService.UpdateCartItem(cartItem)

	// Assert
	assert.NoError(t, err)
	cartItemRepository.AssertExpectations(t)
}
func TestUpdateCartItem_Error(t *testing.T) {
	// Arrange
	cartItemRepository := &mocks.ICartItemRepository{}
	shoppingRepository := &mocks.IShoppingCartRepository{}
	cartItemService := NewCartItemService(cartItemRepository, shoppingRepository)

	error := assert.AnError
	cartItem := BusinessObjects.CartItem{}
	cartItemRepository.On("UpdateCartItem", cartItem).Return(error)

	// Act
	err := cartItemService.UpdateCartItem(cartItem)

	// Assert
	assert.Error(t, err)
	cartItemRepository.AssertExpectations(t)
}

func TestDeleteCartItem_Success(t *testing.T) {
	// Arrange
	cartItemRepository := &mocks.ICartItemRepository{}
	shoppingRepository := &mocks.IShoppingCartRepository{}
	cartItemService := NewCartItemService(cartItemRepository, shoppingRepository)

	cartID := "cartID"
	productID := "productID"
	cartItemRepository.On("DeleteCartItem", cartID, productID).Return(nil)

	// Act
	err := cartItemService.DeleteCartItem(cartID, productID)

	// Assert
	assert.NoError(t, err)
	cartItemRepository.AssertExpectations(t)
}

func TestDeleteCartItem_Error(t *testing.T) {
	// Arrange
	cartItemRepository := &mocks.ICartItemRepository{}
	shoppingRepository := &mocks.IShoppingCartRepository{}
	cartItemService := NewCartItemService(cartItemRepository, shoppingRepository)

	cartID := "cartID"
	productID := "productID"
	error := assert.AnError
	cartItemRepository.On("DeleteCartItem", cartID, productID).Return(error)

	// Act
	err := cartItemService.DeleteCartItem(cartID, productID)

	// Assert
	assert.Error(t, err)
	cartItemRepository.AssertExpectations(t)
}

func TestRemoveItemFromCart_Success_Quantity1(t *testing.T) {
	// Arrange
	cartItemRepository := &mocks.ICartItemRepository{}
	shoppingRepository := &mocks.IShoppingCartRepository{}
	cartItemService := NewCartItemService(cartItemRepository, shoppingRepository)

	cartID := "cartID"
	productID := "productID"

	cartItem := BusinessObjects.ShoppingCart{
		CartID: "cartID",
		UserID: "userID",
		CartItems: []BusinessObjects.CartItem{
			{
				ProductID: "productID",
				Quantity:  1,
			},
		},
	}

	shoppingRepository.On("GetShoppingCartByID", cartID).Return(cartItem, nil)
	shoppingRepository.On("UpdateShoppingCart", mock.AnythingOfType("BusinessObjects.ShoppingCart")).Return(nil)

	// Act
	err := cartItemService.RemoveItemFromCart(cartID, productID)

	// Assert
	assert.NoError(t, err)
	shoppingRepository.AssertExpectations(t)
}
func TestRemoveItemFromCart_Success_Quantity2(t *testing.T) {
	// Arrange
	cartItemRepository := &mocks.ICartItemRepository{}
	shoppingRepository := &mocks.IShoppingCartRepository{}
	cartItemService := NewCartItemService(cartItemRepository, shoppingRepository)

	cartID := "cartID"
	productID := "productID"

	cartItem := BusinessObjects.ShoppingCart{
		CartID: "cartID",
		UserID: "userID",
		CartItems: []BusinessObjects.CartItem{
			{
				ProductID: "productID",
				Quantity:  2,
			},
		},
	}

	shoppingRepository.On("GetShoppingCartByID", cartID).Return(cartItem, nil)
	shoppingRepository.On("UpdateShoppingCart", mock.AnythingOfType("BusinessObjects.ShoppingCart")).Return(nil)

	// Act
	err := cartItemService.RemoveItemFromCart(cartID, productID)

	// Assert
	assert.NoError(t, err)
	shoppingRepository.AssertExpectations(t)
}

func TestRemoveItemFromCart_Error(t *testing.T) {
	// Arrange
	cartItemRepository := &mocks.ICartItemRepository{}
	shoppingRepository := &mocks.IShoppingCartRepository{}
	cartItemService := NewCartItemService(cartItemRepository, shoppingRepository)

	cartID := "cartID"
	productID := "productID"

	shoppingRepository.On("GetShoppingCartByID", cartID).Return(BusinessObjects.ShoppingCart{}, assert.AnError)

	// Act
	err := cartItemService.RemoveItemFromCart(cartID, productID)

	// Assert
	assert.Error(t, err)
	shoppingRepository.AssertExpectations(t)
}
