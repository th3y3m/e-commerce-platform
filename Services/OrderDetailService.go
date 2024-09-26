package Services

import (
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Interface"
	"th3y3m/e-commerce-platform/Util"
)

type OrderDetailService struct {
	orderDetailRepository Interface.IOrderDetailRepository
}

func NewOrderDetailService(orderDetailRepository Interface.IOrderDetailRepository) Interface.IOrderDetailService {
	return &OrderDetailService{orderDetailRepository}
}

func (o *OrderDetailService) GetPaginatedOrderDetailList(searchValue, sortBy, orderId, productId string, pageIndex, pageSize int) (Util.PaginatedList[BusinessObjects.OrderDetail], error) {
	return o.orderDetailRepository.GetPaginatedOrderDetailList(searchValue, sortBy, orderId, productId, pageIndex, pageSize)
}

func (o *OrderDetailService) GetAllOrderDetails() ([]BusinessObjects.OrderDetail, error) {
	return o.orderDetailRepository.GetAllOrderDetails()
}

func (o *OrderDetailService) GetOrderDetailByID(id string) ([]BusinessObjects.OrderDetail, error) {
	return o.orderDetailRepository.GetOrderDetailByID(id)
}

func (o *OrderDetailService) CreateOrderDetail(orderId, productId string, quantity int, unitPrice float64) error {
	orderDetail := BusinessObjects.OrderDetail{
		OrderID:   orderId,
		ProductID: productId,
		Quantity:  quantity,
		UnitPrice: unitPrice,
	}

	return o.orderDetailRepository.CreateOrderDetail(orderDetail)
}

func (o *OrderDetailService) UpdateOrderDetail(orderDetail BusinessObjects.OrderDetail) error {
	return o.orderDetailRepository.UpdateOrderDetail(orderDetail)
}

func (o *OrderDetailService) DeleteOrderDetail(id string) error {
	return o.orderDetailRepository.DeleteOrderDetail(id)
}
