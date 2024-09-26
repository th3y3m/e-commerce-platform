package Interface

import (
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Util"
	"time"
)

type IOrderRepository interface {
	GetPaginatedOrderList(sortBy, orderID, customerId, courierId, voucherId string, pageIndex, pageSize int, startDate, endDate *time.Time, minPrice, maxPrice *float64, status string) (Util.PaginatedList[BusinessObjects.Order], error)
	GetAllOrders() ([]BusinessObjects.Order, error)
	GetOrderByID(orderID string) (BusinessObjects.Order, error)
	CreateOrder(order BusinessObjects.Order) error
	UpdateOrder(order BusinessObjects.Order) error
	DeleteOrder(orderID string) error
}
type IOrderService interface {
	GetPaginatedOrderList(sortBy, orderID, customerId, courierId, voucherId string, pageIndex, pageSize int, startDate, endDate *time.Time, minPrice, maxPrice *float64, status string) (Util.PaginatedList[BusinessObjects.Order], error)
	GetAllOrders() ([]BusinessObjects.Order, error)
	GetOrderById(orderID string) (BusinessObjects.Order, error)
	CreateOrder(order BusinessObjects.NewOrder) error
	PlaceOrder(userId, cartId, shipAddress, CourierID, VoucherID string) error
	UpdateOrder(order BusinessObjects.Order) error
	CancelOrder(orderID string) error
}
