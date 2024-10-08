// Code generated by mockery v2.46.1. DO NOT EDIT.

package mocks

import (
	BusinessObjects "th3y3m/e-commerce-platform/BusinessObjects"

	Util "th3y3m/e-commerce-platform/Util"

	mock "github.com/stretchr/testify/mock"

	time "time"
)

// ITransactionService is an autogenerated mock type for the ITransactionService type
type ITransactionService struct {
	mock.Mock
}

// CreateTransaction provides a mock function with given fields: transaction
func (_m *ITransactionService) CreateTransaction(transaction BusinessObjects.NewTransaction) error {
	ret := _m.Called(transaction)

	if len(ret) == 0 {
		panic("no return value specified for CreateTransaction")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(BusinessObjects.NewTransaction) error); ok {
		r0 = rf(transaction)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteTransaction provides a mock function with given fields: transactionID
func (_m *ITransactionService) DeleteTransaction(transactionID string) error {
	ret := _m.Called(transactionID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteTransaction")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(transactionID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllTransactions provides a mock function with given fields:
func (_m *ITransactionService) GetAllTransactions() ([]BusinessObjects.Transaction, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAllTransactions")
	}

	var r0 []BusinessObjects.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]BusinessObjects.Transaction, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []BusinessObjects.Transaction); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]BusinessObjects.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPaginatedTransactionList provides a mock function with given fields: sortBy, transactionID, orderID, pageIndex, pageSize, minPrice, maxPrice, status, startDate, endDate
func (_m *ITransactionService) GetPaginatedTransactionList(sortBy string, transactionID string, orderID string, pageIndex int, pageSize int, minPrice *float64, maxPrice *float64, status *bool, startDate time.Time, endDate time.Time) (Util.PaginatedList[BusinessObjects.Transaction], error) {
	ret := _m.Called(sortBy, transactionID, orderID, pageIndex, pageSize, minPrice, maxPrice, status, startDate, endDate)

	if len(ret) == 0 {
		panic("no return value specified for GetPaginatedTransactionList")
	}

	var r0 Util.PaginatedList[BusinessObjects.Transaction]
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, string, int, int, *float64, *float64, *bool, time.Time, time.Time) (Util.PaginatedList[BusinessObjects.Transaction], error)); ok {
		return rf(sortBy, transactionID, orderID, pageIndex, pageSize, minPrice, maxPrice, status, startDate, endDate)
	}
	if rf, ok := ret.Get(0).(func(string, string, string, int, int, *float64, *float64, *bool, time.Time, time.Time) Util.PaginatedList[BusinessObjects.Transaction]); ok {
		r0 = rf(sortBy, transactionID, orderID, pageIndex, pageSize, minPrice, maxPrice, status, startDate, endDate)
	} else {
		r0 = ret.Get(0).(Util.PaginatedList[BusinessObjects.Transaction])
	}

	if rf, ok := ret.Get(1).(func(string, string, string, int, int, *float64, *float64, *bool, time.Time, time.Time) error); ok {
		r1 = rf(sortBy, transactionID, orderID, pageIndex, pageSize, minPrice, maxPrice, status, startDate, endDate)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransactionByID provides a mock function with given fields: transactionID
func (_m *ITransactionService) GetTransactionByID(transactionID string) (BusinessObjects.Transaction, error) {
	ret := _m.Called(transactionID)

	if len(ret) == 0 {
		panic("no return value specified for GetTransactionByID")
	}

	var r0 BusinessObjects.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (BusinessObjects.Transaction, error)); ok {
		return rf(transactionID)
	}
	if rf, ok := ret.Get(0).(func(string) BusinessObjects.Transaction); ok {
		r0 = rf(transactionID)
	} else {
		r0 = ret.Get(0).(BusinessObjects.Transaction)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(transactionID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTransaction provides a mock function with given fields: transaction
func (_m *ITransactionService) UpdateTransaction(transaction BusinessObjects.Transaction) error {
	ret := _m.Called(transaction)

	if len(ret) == 0 {
		panic("no return value specified for UpdateTransaction")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(BusinessObjects.Transaction) error); ok {
		r0 = rf(transaction)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewITransactionService creates a new instance of ITransactionService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewITransactionService(t interface {
	mock.TestingT
	Cleanup(func())
}) *ITransactionService {
	mock := &ITransactionService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
