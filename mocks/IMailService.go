// Code generated by mockery v2.46.1. DO NOT EDIT.

package mocks

import (
	BusinessObjects "th3y3m/e-commerce-platform/BusinessObjects"

	mock "github.com/stretchr/testify/mock"
)

// IMailService is an autogenerated mock type for the IMailService type
type IMailService struct {
	mock.Mock
}

// SendMail provides a mock function with given fields: to, token
func (_m *IMailService) SendMail(to string, token string) error {
	ret := _m.Called(to, token)

	if len(ret) == 0 {
		panic("no return value specified for SendMail")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(to, token)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendOrderDetails provides a mock function with given fields: Customer, Order, OrderDetails
func (_m *IMailService) SendOrderDetails(Customer BusinessObjects.User, Order BusinessObjects.Order, OrderDetails []BusinessObjects.OrderDetail) error {
	ret := _m.Called(Customer, Order, OrderDetails)

	if len(ret) == 0 {
		panic("no return value specified for SendOrderDetails")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(BusinessObjects.User, BusinessObjects.Order, []BusinessObjects.OrderDetail) error); ok {
		r0 = rf(Customer, Order, OrderDetails)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// VerifyToken provides a mock function with given fields: token
func (_m *IMailService) VerifyToken(token string) bool {
	ret := _m.Called(token)

	if len(ret) == 0 {
		panic("no return value specified for VerifyToken")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(token)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// NewIMailService creates a new instance of IMailService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIMailService(t interface {
	mock.TestingT
	Cleanup(func())
}) *IMailService {
	mock := &IMailService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
