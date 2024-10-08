package Services

import (
	"testing"
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Util"
	"th3y3m/e-commerce-platform/mocks"

	"github.com/stretchr/testify/assert"
)

func TestGetPaginatedOrderDetailList_Success(t *testing.T) {
	orderDetailRepository := &mocks.IOrderDetailRepository{}
	OrderDetailService := NewOrderDetailService(orderDetailRepository)

	searchValue := ""
	sortBy := ""
	orderId := ""
	productId := ""
	pageIndex := 1
	pageSize := 10

	orderDetailRepository.On("GetPaginatedOrderDetailList", searchValue, sortBy, orderId, productId, pageIndex, pageSize).Return(Util.PaginatedList[BusinessObjects.OrderDetail]{}, nil)

	_, err := OrderDetailService.GetPaginatedOrderDetailList(searchValue, sortBy, orderId, productId, pageIndex, pageSize)

	assert.NoError(t, err)
	orderDetailRepository.AssertExpectations(t)
}

func TestGetAllOrderDetails_Success(t *testing.T) {
	orderDetailRepository := &mocks.IOrderDetailRepository{}
	OrderDetailService := NewOrderDetailService(orderDetailRepository)

	orderDetailRepository.On("GetAllOrderDetails").Return([]BusinessObjects.OrderDetail{}, nil)

	_, err := OrderDetailService.GetAllOrderDetails()

	assert.NoError(t, err)
	orderDetailRepository.AssertExpectations(t)
}

func TestGetOrderDetailById_Success(t *testing.T) {
	orderDetailRepository := &mocks.IOrderDetailRepository{}
	OrderDetailService := NewOrderDetailService(orderDetailRepository)

	id := "1"

	orderDetailRepository.On("GetOrderDetailByID", id).Return([]BusinessObjects.OrderDetail{}, nil)

	_, err := OrderDetailService.GetOrderDetailByID(id)

	assert.NoError(t, err)
	orderDetailRepository.AssertExpectations(t)
}

func TestCreateOrderDetail_Success(t *testing.T) {
	orderDetailRepository := &mocks.IOrderDetailRepository{}
	OrderDetailService := NewOrderDetailService(orderDetailRepository)

	orderDetail := BusinessObjects.OrderDetail{
		OrderID:   "1",
		ProductID: "1",
		Quantity:  1,
		UnitPrice: 100,
	}

	orderDetailRepository.On("CreateOrderDetail", orderDetail).Return(nil)

	err := OrderDetailService.CreateOrderDetail(orderDetail.OrderID, orderDetail.ProductID, orderDetail.Quantity, orderDetail.UnitPrice)

	assert.NoError(t, err)
	orderDetailRepository.AssertExpectations(t)
}

func TestUpdateOrderDetail_Success(t *testing.T) {
	orderDetailRepository := &mocks.IOrderDetailRepository{}
	OrderDetailService := NewOrderDetailService(orderDetailRepository)

	orderDetail := BusinessObjects.OrderDetail{
		OrderID:   "1",
		ProductID: "1",
		Quantity:  1,
		UnitPrice: 100,
	}

	orderDetailRepository.On("UpdateOrderDetail", orderDetail).Return(nil)

	err := OrderDetailService.UpdateOrderDetail(orderDetail)

	assert.NoError(t, err)
	orderDetailRepository.AssertExpectations(t)
}

func TestDeleteOrderDetail_Success(t *testing.T) {
	orderDetailRepository := &mocks.IOrderDetailRepository{}
	OrderDetailService := NewOrderDetailService(orderDetailRepository)

	id := "1"

	orderDetailRepository.On("DeleteOrderDetail", id).Return(nil)

	err := OrderDetailService.DeleteOrderDetail(id)

	assert.NoError(t, err)
	orderDetailRepository.AssertExpectations(t)
}
