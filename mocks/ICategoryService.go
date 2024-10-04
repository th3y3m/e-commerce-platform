// Code generated by mockery v2.46.1. DO NOT EDIT.

package mocks

import (
	BusinessObjects "th3y3m/e-commerce-platform/BusinessObjects"

	Util "th3y3m/e-commerce-platform/Util"

	mock "github.com/stretchr/testify/mock"
)

// ICategoryService is an autogenerated mock type for the ICategoryService type
type ICategoryService struct {
	mock.Mock
}

// CreateCategory provides a mock function with given fields: CategoryName
func (_m *ICategoryService) CreateCategory(CategoryName string) error {
	ret := _m.Called(CategoryName)

	if len(ret) == 0 {
		panic("no return value specified for CreateCategory")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(CategoryName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteCategory provides a mock function with given fields: categoryID
func (_m *ICategoryService) DeleteCategory(categoryID string) error {
	ret := _m.Called(categoryID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteCategory")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(categoryID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllCategories provides a mock function with given fields:
func (_m *ICategoryService) GetAllCategories() ([]BusinessObjects.Category, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAllCategories")
	}

	var r0 []BusinessObjects.Category
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]BusinessObjects.Category, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []BusinessObjects.Category); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]BusinessObjects.Category)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCategoryByID provides a mock function with given fields: categoryID
func (_m *ICategoryService) GetCategoryByID(categoryID string) (BusinessObjects.Category, error) {
	ret := _m.Called(categoryID)

	if len(ret) == 0 {
		panic("no return value specified for GetCategoryByID")
	}

	var r0 BusinessObjects.Category
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (BusinessObjects.Category, error)); ok {
		return rf(categoryID)
	}
	if rf, ok := ret.Get(0).(func(string) BusinessObjects.Category); ok {
		r0 = rf(categoryID)
	} else {
		r0 = ret.Get(0).(BusinessObjects.Category)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(categoryID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPaginatedCategoryList provides a mock function with given fields: searchValue, sortBy, pageIndex, pageSize, status
func (_m *ICategoryService) GetPaginatedCategoryList(searchValue string, sortBy string, pageIndex int, pageSize int, status *bool) (Util.PaginatedList[BusinessObjects.Category], error) {
	ret := _m.Called(searchValue, sortBy, pageIndex, pageSize, status)

	if len(ret) == 0 {
		panic("no return value specified for GetPaginatedCategoryList")
	}

	var r0 Util.PaginatedList[BusinessObjects.Category]
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, int, int, *bool) (Util.PaginatedList[BusinessObjects.Category], error)); ok {
		return rf(searchValue, sortBy, pageIndex, pageSize, status)
	}
	if rf, ok := ret.Get(0).(func(string, string, int, int, *bool) Util.PaginatedList[BusinessObjects.Category]); ok {
		r0 = rf(searchValue, sortBy, pageIndex, pageSize, status)
	} else {
		r0 = ret.Get(0).(Util.PaginatedList[BusinessObjects.Category])
	}

	if rf, ok := ret.Get(1).(func(string, string, int, int, *bool) error); ok {
		r1 = rf(searchValue, sortBy, pageIndex, pageSize, status)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateCategory provides a mock function with given fields: categoryID, CategoryName
func (_m *ICategoryService) UpdateCategory(categoryID string, CategoryName string) error {
	ret := _m.Called(categoryID, CategoryName)

	if len(ret) == 0 {
		panic("no return value specified for UpdateCategory")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(categoryID, CategoryName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewICategoryService creates a new instance of ICategoryService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewICategoryService(t interface {
	mock.TestingT
	Cleanup(func())
}) *ICategoryService {
	mock := &ICategoryService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
