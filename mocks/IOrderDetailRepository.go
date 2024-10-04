// Code generated by mockery v2.46.1. DO NOT EDIT.

package mocks

import (
	BusinessObjects "th3y3m/e-commerce-platform/BusinessObjects"

	Util "th3y3m/e-commerce-platform/Util"

	mock "github.com/stretchr/testify/mock"
)

// IOrderDetailRepository is an autogenerated mock type for the IOrderDetailRepository type
type IOrderDetailRepository struct {
	mock.Mock
}

// CreateOrderDetail provides a mock function with given fields: rate
func (_m *IOrderDetailRepository) CreateOrderDetail(rate BusinessObjects.OrderDetail) error {
	ret := _m.Called(rate)

	if len(ret) == 0 {
		panic("no return value specified for CreateOrderDetail")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(BusinessObjects.OrderDetail) error); ok {
		r0 = rf(rate)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteOrderDetail provides a mock function with given fields: rateID
func (_m *IOrderDetailRepository) DeleteOrderDetail(rateID string) error {
	ret := _m.Called(rateID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteOrderDetail")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(rateID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllOrderDetails provides a mock function with given fields:
func (_m *IOrderDetailRepository) GetAllOrderDetails() ([]BusinessObjects.OrderDetail, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAllOrderDetails")
	}

	var r0 []BusinessObjects.OrderDetail
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]BusinessObjects.OrderDetail, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []BusinessObjects.OrderDetail); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]BusinessObjects.OrderDetail)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOrderDetailByID provides a mock function with given fields: orderID
func (_m *IOrderDetailRepository) GetOrderDetailByID(orderID string) ([]BusinessObjects.OrderDetail, error) {
	ret := _m.Called(orderID)

	if len(ret) == 0 {
		panic("no return value specified for GetOrderDetailByID")
	}

	var r0 []BusinessObjects.OrderDetail
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]BusinessObjects.OrderDetail, error)); ok {
		return rf(orderID)
	}
	if rf, ok := ret.Get(0).(func(string) []BusinessObjects.OrderDetail); ok {
		r0 = rf(orderID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]BusinessObjects.OrderDetail)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(orderID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPaginatedOrderDetailList provides a mock function with given fields: searchValue, sortBy, orderId, productId, pageIndex, pageSize
func (_m *IOrderDetailRepository) GetPaginatedOrderDetailList(searchValue string, sortBy string, orderId string, productId string, pageIndex int, pageSize int) (Util.PaginatedList[BusinessObjects.OrderDetail], error) {
	ret := _m.Called(searchValue, sortBy, orderId, productId, pageIndex, pageSize)

	if len(ret) == 0 {
		panic("no return value specified for GetPaginatedOrderDetailList")
	}

	var r0 Util.PaginatedList[BusinessObjects.OrderDetail]
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, string, string, int, int) (Util.PaginatedList[BusinessObjects.OrderDetail], error)); ok {
		return rf(searchValue, sortBy, orderId, productId, pageIndex, pageSize)
	}
	if rf, ok := ret.Get(0).(func(string, string, string, string, int, int) Util.PaginatedList[BusinessObjects.OrderDetail]); ok {
		r0 = rf(searchValue, sortBy, orderId, productId, pageIndex, pageSize)
	} else {
		r0 = ret.Get(0).(Util.PaginatedList[BusinessObjects.OrderDetail])
	}

	if rf, ok := ret.Get(1).(func(string, string, string, string, int, int) error); ok {
		r1 = rf(searchValue, sortBy, orderId, productId, pageIndex, pageSize)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateOrderDetail provides a mock function with given fields: rate
func (_m *IOrderDetailRepository) UpdateOrderDetail(rate BusinessObjects.OrderDetail) error {
	ret := _m.Called(rate)

	if len(ret) == 0 {
		panic("no return value specified for UpdateOrderDetail")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(BusinessObjects.OrderDetail) error); ok {
		r0 = rf(rate)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewIOrderDetailRepository creates a new instance of IOrderDetailRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIOrderDetailRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *IOrderDetailRepository {
	mock := &IOrderDetailRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
