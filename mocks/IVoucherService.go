// Code generated by mockery v2.46.1. DO NOT EDIT.

package mocks

import (
	BusinessObjects "th3y3m/e-commerce-platform/BusinessObjects"

	Util "th3y3m/e-commerce-platform/Util"

	mock "github.com/stretchr/testify/mock"

	time "time"
)

// IVoucherService is an autogenerated mock type for the IVoucherService type
type IVoucherService struct {
	mock.Mock
}

// CreateVoucher provides a mock function with given fields: vou
func (_m *IVoucherService) CreateVoucher(vou BusinessObjects.NewVoucher) error {
	ret := _m.Called(vou)

	if len(ret) == 0 {
		panic("no return value specified for CreateVoucher")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(BusinessObjects.NewVoucher) error); ok {
		r0 = rf(vou)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteVoucher provides a mock function with given fields: voucherID
func (_m *IVoucherService) DeleteVoucher(voucherID string) error {
	ret := _m.Called(voucherID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteVoucher")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(voucherID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllVouchers provides a mock function with given fields:
func (_m *IVoucherService) GetAllVouchers() ([]BusinessObjects.Voucher, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAllVouchers")
	}

	var r0 []BusinessObjects.Voucher
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]BusinessObjects.Voucher, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []BusinessObjects.Voucher); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]BusinessObjects.Voucher)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPaginatedVoucherList provides a mock function with given fields: sortBy, voucherID, pageIndex, pageSize, status, startDate, endDate
func (_m *IVoucherService) GetPaginatedVoucherList(sortBy string, voucherID string, pageIndex int, pageSize int, status *bool, startDate time.Time, endDate time.Time) (Util.PaginatedList[BusinessObjects.Voucher], error) {
	ret := _m.Called(sortBy, voucherID, pageIndex, pageSize, status, startDate, endDate)

	if len(ret) == 0 {
		panic("no return value specified for GetPaginatedVoucherList")
	}

	var r0 Util.PaginatedList[BusinessObjects.Voucher]
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, int, int, *bool, time.Time, time.Time) (Util.PaginatedList[BusinessObjects.Voucher], error)); ok {
		return rf(sortBy, voucherID, pageIndex, pageSize, status, startDate, endDate)
	}
	if rf, ok := ret.Get(0).(func(string, string, int, int, *bool, time.Time, time.Time) Util.PaginatedList[BusinessObjects.Voucher]); ok {
		r0 = rf(sortBy, voucherID, pageIndex, pageSize, status, startDate, endDate)
	} else {
		r0 = ret.Get(0).(Util.PaginatedList[BusinessObjects.Voucher])
	}

	if rf, ok := ret.Get(1).(func(string, string, int, int, *bool, time.Time, time.Time) error); ok {
		r1 = rf(sortBy, voucherID, pageIndex, pageSize, status, startDate, endDate)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetVoucherByID provides a mock function with given fields: voucherID
func (_m *IVoucherService) GetVoucherByID(voucherID string) (BusinessObjects.Voucher, error) {
	ret := _m.Called(voucherID)

	if len(ret) == 0 {
		panic("no return value specified for GetVoucherByID")
	}

	var r0 BusinessObjects.Voucher
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (BusinessObjects.Voucher, error)); ok {
		return rf(voucherID)
	}
	if rf, ok := ret.Get(0).(func(string) BusinessObjects.Voucher); ok {
		r0 = rf(voucherID)
	} else {
		r0 = ret.Get(0).(BusinessObjects.Voucher)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(voucherID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateVoucher provides a mock function with given fields: voucher
func (_m *IVoucherService) UpdateVoucher(voucher BusinessObjects.Voucher) error {
	ret := _m.Called(voucher)

	if len(ret) == 0 {
		panic("no return value specified for UpdateVoucher")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(BusinessObjects.Voucher) error); ok {
		r0 = rf(voucher)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewIVoucherService creates a new instance of IVoucherService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIVoucherService(t interface {
	mock.TestingT
	Cleanup(func())
}) *IVoucherService {
	mock := &IVoucherService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
