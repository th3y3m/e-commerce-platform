package Interface

import (
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Util"
)

type IOrderDetailRepository interface {
	GetPaginatedOrderDetailList(searchValue, sortBy, orderId, productId string, pageIndex, pageSize int) (Util.PaginatedList[BusinessObjects.OrderDetail], error)
	GetAllOrderDetails() ([]BusinessObjects.OrderDetail, error)
	GetOrderDetailByID(orderID string) ([]BusinessObjects.OrderDetail, error)
	CreateOrderDetail(rate BusinessObjects.OrderDetail) error
	UpdateOrderDetail(rate BusinessObjects.OrderDetail) error
	DeleteOrderDetail(rateID string) error
}
type IOrderDetailService interface {
	GetPaginatedOrderDetailList(searchValue, sortBy, orderId, productId string, pageIndex, pageSize int) (Util.PaginatedList[BusinessObjects.OrderDetail], error)
	GetAllOrderDetails() ([]BusinessObjects.OrderDetail, error)
	GetOrderDetailByID(orderID string) ([]BusinessObjects.OrderDetail, error)
	CreateOrderDetail(orderId, productId string, quantity int, unitPrice float64) error
	UpdateOrderDetail(rate BusinessObjects.OrderDetail) error
	DeleteOrderDetail(rateID string) error
}
