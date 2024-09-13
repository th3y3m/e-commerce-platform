package Services

import (
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Repositories"
	"th3y3m/e-commerce-platform/Util"
	"time"
)

func GetPaginatedOrderList(sortBy string, orderID string, customerId string, courierId string, voucherId string, pageIndex int, pageSize int, startDate time.Time, endDate time.Time) (Util.PaginatedList[BusinessObjects.Order], error) {
	return Repositories.GetPaginatedOrderList(sortBy, orderID, customerId, courierId, voucherId, pageIndex, pageSize, startDate, endDate)
}

func GetAllOrders() ([]BusinessObjects.Order, error) {
	return Repositories.GetAllOrders()
}

func GetOrderById(id string) (BusinessObjects.Order, error) {
	return Repositories.GetOrderByID(id)
}

func CreateOrder(order BusinessObjects.NewOrder) error {
	newOrder := BusinessObjects.Order{
		OrderID:               "ORD" + Util.GenerateID(10),
		CustomerID:            order.CustomerID,
		CourierID:             order.CourierID,
		VoucherID:             order.VoucherID,
		TotalAmount:           order.TotalAmount,
		OrderDate:             time.Now(),
		OrderStatus:           "Pending",
		PaymentMethod:         order.PaymentMethod,
		ShippingAddress:       order.ShippingAddress,
		FreightPrice:          order.FreightPrice,
		EstimatedDeliveryDate: order.EstimatedDeliveryDate,
		ActualDeliveryDate:    order.ActualDeliveryDate,
		PaymentStatus:         order.PaymentStatus,
	}

	err := Repositories.CreateOrder(newOrder)
	if err != nil {
		return err
	}

	return nil
}

func UpdateOrder(order BusinessObjects.Order) error {
	return Repositories.UpdateOrder(order)
}

func CancelOrder(id string) error {
	return Repositories.DeleteOrder(id)
}
