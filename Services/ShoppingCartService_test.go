package Services

import (
	"testing"
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Util"
	"th3y3m/e-commerce-platform/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetPaginatedShoppingCartList_Success(t *testing.T) {
	shoppingCartRepository := &mocks.IShoppingCartRepository{}
	cartItemRepository := &mocks.ICartItemRepository{}
	shoppingCartService := NewShoppingCartService(shoppingCartRepository, cartItemRepository)
	var status *bool
	shoppingCartRepository.On("GetPaginatedShoppingCartList", "sortBy", "cartID", "userID", 1, 10, status).Return(Util.PaginatedList[BusinessObjects.ShoppingCart]{}, nil)

	_, err := shoppingCartService.GetPaginatedShoppingCartList("sortBy", "cartID", "userID", 1, 10, status)

	assert.NoError(t, err)
	shoppingCartRepository.AssertExpectations(t)
}

func TestGetAllShoppingCarts_Success(t *testing.T) {
	shoppingCartRepository := &mocks.IShoppingCartRepository{}
	cartItemRepository := &mocks.ICartItemRepository{}
	shoppingCartService := NewShoppingCartService(shoppingCartRepository, cartItemRepository)
	shoppingCartRepository.On("GetAllShoppingCarts").Return([]BusinessObjects.ShoppingCart{}, nil)

	_, err := shoppingCartService.GetAllShoppingCarts()

	assert.NoError(t, err)
	shoppingCartRepository.AssertExpectations(t)
}

func TestGetShoppingCartByID_Success(t *testing.T) {
	shoppingCartRepository := &mocks.IShoppingCartRepository{}
	cartItemRepository := &mocks.ICartItemRepository{}
	shoppingCartService := NewShoppingCartService(shoppingCartRepository, cartItemRepository)
	shoppingCartRepository.On("GetShoppingCartByID", "id").Return(BusinessObjects.ShoppingCart{}, nil)

	_, err := shoppingCartService.GetShoppingCartByID("id")

	assert.NoError(t, err)
	shoppingCartRepository.AssertExpectations(t)
}

func TestUpdateShoppingCartStatus_Success(t *testing.T) {
	shoppingCartRepository := &mocks.IShoppingCartRepository{}
	cartItemRepository := &mocks.ICartItemRepository{}
	shoppingCartService := NewShoppingCartService(shoppingCartRepository, cartItemRepository)
	shoppingCartRepository.On("UpdateShoppingCartStatus", "id", true).Return(nil)

	err := shoppingCartService.UpdateShoppingCartStatus("id", true)

	assert.NoError(t, err)
	shoppingCartRepository.AssertExpectations(t)
}

func TestCreateShoppingCart_Success(t *testing.T) {
	shoppingCartRepository := &mocks.IShoppingCartRepository{}
	cartItemRepository := &mocks.ICartItemRepository{}
	shoppingCartService := NewShoppingCartService(shoppingCartRepository, cartItemRepository)
	shoppingCart := BusinessObjects.ShoppingCart{
		UserID: "userID",
	}
	shoppingCartRepository.On("CreateShoppingCart", mock.AnythingOfType("BusinessObjects.ShoppingCart")).Return(BusinessObjects.ShoppingCart{}, nil)

	_, err := shoppingCartService.CreateShoppingCart(shoppingCart.UserID)

	assert.NoError(t, err)
	shoppingCartRepository.AssertExpectations(t)
}
