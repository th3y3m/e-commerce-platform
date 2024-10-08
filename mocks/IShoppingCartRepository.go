// Code generated by mockery v2.46.1. DO NOT EDIT.

package mocks

import (
	BusinessObjects "th3y3m/e-commerce-platform/BusinessObjects"

	Util "th3y3m/e-commerce-platform/Util"

	mock "github.com/stretchr/testify/mock"
)

// IShoppingCartRepository is an autogenerated mock type for the IShoppingCartRepository type
type IShoppingCartRepository struct {
	mock.Mock
}

// CreateShoppingCart provides a mock function with given fields: cart
func (_m *IShoppingCartRepository) CreateShoppingCart(cart BusinessObjects.ShoppingCart) (BusinessObjects.ShoppingCart, error) {
	ret := _m.Called(cart)

	if len(ret) == 0 {
		panic("no return value specified for CreateShoppingCart")
	}

	var r0 BusinessObjects.ShoppingCart
	var r1 error
	if rf, ok := ret.Get(0).(func(BusinessObjects.ShoppingCart) (BusinessObjects.ShoppingCart, error)); ok {
		return rf(cart)
	}
	if rf, ok := ret.Get(0).(func(BusinessObjects.ShoppingCart) BusinessObjects.ShoppingCart); ok {
		r0 = rf(cart)
	} else {
		r0 = ret.Get(0).(BusinessObjects.ShoppingCart)
	}

	if rf, ok := ret.Get(1).(func(BusinessObjects.ShoppingCart) error); ok {
		r1 = rf(cart)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteShoppingCart provides a mock function with given fields: cartID
func (_m *IShoppingCartRepository) DeleteShoppingCart(cartID string) error {
	ret := _m.Called(cartID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteShoppingCart")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(cartID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllShoppingCarts provides a mock function with given fields:
func (_m *IShoppingCartRepository) GetAllShoppingCarts() ([]BusinessObjects.ShoppingCart, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAllShoppingCarts")
	}

	var r0 []BusinessObjects.ShoppingCart
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]BusinessObjects.ShoppingCart, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []BusinessObjects.ShoppingCart); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]BusinessObjects.ShoppingCart)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPaginatedShoppingCartList provides a mock function with given fields: sortBy, cartID, userID, pageIndex, pageSize, status
func (_m *IShoppingCartRepository) GetPaginatedShoppingCartList(sortBy string, cartID string, userID string, pageIndex int, pageSize int, status *bool) (Util.PaginatedList[BusinessObjects.ShoppingCart], error) {
	ret := _m.Called(sortBy, cartID, userID, pageIndex, pageSize, status)

	if len(ret) == 0 {
		panic("no return value specified for GetPaginatedShoppingCartList")
	}

	var r0 Util.PaginatedList[BusinessObjects.ShoppingCart]
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, string, int, int, *bool) (Util.PaginatedList[BusinessObjects.ShoppingCart], error)); ok {
		return rf(sortBy, cartID, userID, pageIndex, pageSize, status)
	}
	if rf, ok := ret.Get(0).(func(string, string, string, int, int, *bool) Util.PaginatedList[BusinessObjects.ShoppingCart]); ok {
		r0 = rf(sortBy, cartID, userID, pageIndex, pageSize, status)
	} else {
		r0 = ret.Get(0).(Util.PaginatedList[BusinessObjects.ShoppingCart])
	}

	if rf, ok := ret.Get(1).(func(string, string, string, int, int, *bool) error); ok {
		r1 = rf(sortBy, cartID, userID, pageIndex, pageSize, status)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetShoppingCartByID provides a mock function with given fields: cartID
func (_m *IShoppingCartRepository) GetShoppingCartByID(cartID string) (BusinessObjects.ShoppingCart, error) {
	ret := _m.Called(cartID)

	if len(ret) == 0 {
		panic("no return value specified for GetShoppingCartByID")
	}

	var r0 BusinessObjects.ShoppingCart
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (BusinessObjects.ShoppingCart, error)); ok {
		return rf(cartID)
	}
	if rf, ok := ret.Get(0).(func(string) BusinessObjects.ShoppingCart); ok {
		r0 = rf(cartID)
	} else {
		r0 = ret.Get(0).(BusinessObjects.ShoppingCart)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(cartID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserShoppingCart provides a mock function with given fields: userID
func (_m *IShoppingCartRepository) GetUserShoppingCart(userID string) (BusinessObjects.ShoppingCart, error) {
	ret := _m.Called(userID)

	if len(ret) == 0 {
		panic("no return value specified for GetUserShoppingCart")
	}

	var r0 BusinessObjects.ShoppingCart
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (BusinessObjects.ShoppingCart, error)); ok {
		return rf(userID)
	}
	if rf, ok := ret.Get(0).(func(string) BusinessObjects.ShoppingCart); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Get(0).(BusinessObjects.ShoppingCart)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateShoppingCart provides a mock function with given fields: cart
func (_m *IShoppingCartRepository) UpdateShoppingCart(cart BusinessObjects.ShoppingCart) error {
	ret := _m.Called(cart)

	if len(ret) == 0 {
		panic("no return value specified for UpdateShoppingCart")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(BusinessObjects.ShoppingCart) error); ok {
		r0 = rf(cart)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateShoppingCartStatus provides a mock function with given fields: cartID, status
func (_m *IShoppingCartRepository) UpdateShoppingCartStatus(cartID string, status bool) error {
	ret := _m.Called(cartID, status)

	if len(ret) == 0 {
		panic("no return value specified for UpdateShoppingCartStatus")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, bool) error); ok {
		r0 = rf(cartID, status)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewIShoppingCartRepository creates a new instance of IShoppingCartRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIShoppingCartRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *IShoppingCartRepository {
	mock := &IShoppingCartRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
