package Services

import (
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Repositories"
	"th3y3m/e-commerce-platform/Util"
	"time"
)

func GetPaginatedOrderList(sortBy, orderID, customerId, courierId, voucherId string, pageIndex, pageSize int, startDate, endDate *time.Time, minPrice, maxPrice *float64, status string) (Util.PaginatedList[BusinessObjects.Order], error) {
	return Repositories.GetPaginatedOrderList(sortBy, orderID, customerId, courierId, voucherId, pageIndex, pageSize, startDate, endDate, minPrice, maxPrice, status)
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

func PlaceOrder(userId, cartId, shipAddress, CourierID, VoucherID string) error {
	productsList, err := Repositories.GetCartItemByCartID(cartId)
	if err != nil {
		return err
	}

	totalAmount := 0.0
	for _, product := range productsList {
		p, err := Repositories.GetProductByID(product.ProductID)
		if err != nil {
			return err
		}

		totalAmount += p.Price * float64(product.Quantity)
	}

	newOrder := BusinessObjects.NewOrder{
		CustomerID:            userId,
		CourierID:             CourierID,
		VoucherID:             VoucherID,
		TotalAmount:           totalAmount,
		PaymentMethod:         "",
		ShippingAddress:       shipAddress,
		FreightPrice:          0,
		EstimatedDeliveryDate: time.Now(),
		ActualDeliveryDate:    time.Now(),
		PaymentStatus:         "Pending",
	}

	err = CreateOrder(newOrder)
	if err != nil {
		return err
	}

	if err := UpdateShoppingCartStatus(cartId, false); err != nil {
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
