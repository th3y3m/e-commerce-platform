package Services

import (
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Interface"
	"th3y3m/e-commerce-platform/Provider/RabbitMQ"
	"th3y3m/e-commerce-platform/Util"
	"time"
)

type OrderService struct {
	orderRepository     Interface.IOrderRepository
	cartItemRepository  Interface.ICartItemRepository
	productRepository   Interface.IProductRepository
	shoppingCartService Interface.IShoppingCartService
}

func NewOrderService(orderRepository Interface.IOrderRepository, cartItemRepository Interface.ICartItemRepository, productRepository Interface.IProductRepository, shoppingCartService Interface.IShoppingCartService,
) Interface.IOrderService {
	return &OrderService{
		orderRepository:     orderRepository,
		cartItemRepository:  cartItemRepository,
		productRepository:   productRepository,
		shoppingCartService: shoppingCartService,
	}
}

func (o *OrderService) GetPaginatedOrderList(sortBy, orderID, customerId, courierId, voucherId string, pageIndex, pageSize int, startDate, endDate *time.Time, minPrice, maxPrice *float64, status string) (Util.PaginatedList[BusinessObjects.Order], error) {
	return o.orderRepository.GetPaginatedOrderList(sortBy, orderID, customerId, courierId, voucherId, pageIndex, pageSize, startDate, endDate, minPrice, maxPrice, status)
}

func (o *OrderService) GetAllOrders() ([]BusinessObjects.Order, error) {
	return o.orderRepository.GetAllOrders()
}

func (o *OrderService) GetOrderById(id string) (BusinessObjects.Order, error) {
	return o.orderRepository.GetOrderByID(id)
}

func (o *OrderService) CreateOrder(order BusinessObjects.NewOrder) error {
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

	err := o.orderRepository.CreateOrder(newOrder)
	if err != nil {
		return err
	}

	return nil
}

func (o *OrderService) PlaceOrder(userId, cartId, shipAddress, CourierID, VoucherID string) error {
	err := RabbitMQ.PublishMessage("Order placed")
	if err != nil {
		return err
	}

	productsList, err := o.cartItemRepository.GetCartItemByCartID(cartId)
	if err != nil {
		return err
	}

	totalAmount := 0.0
	for _, product := range productsList {
		p, err := o.productRepository.GetProductByID(product.ProductID)
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

	err = o.CreateOrder(newOrder)
	if err != nil {
		return err
	}

	if err := o.shoppingCartService.UpdateShoppingCartStatus(cartId, false); err != nil {
		return err
	}

	return nil
}

func (o *OrderService) UpdateOrder(order BusinessObjects.Order) error {
	return o.orderRepository.UpdateOrder(order)
}

func (o *OrderService) CancelOrder(id string) error {
	return o.orderRepository.DeleteOrder(id)
}
