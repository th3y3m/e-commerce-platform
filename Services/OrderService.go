package Services

import (
	"encoding/json"
	"log"
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Interface"
	"th3y3m/e-commerce-platform/Util"
	"time"

	"github.com/streadway/amqp"
)

type OrderService struct {
	orderRepository     Interface.IOrderRepository
	cartItemRepository  Interface.ICartItemRepository
	productRepository   Interface.IProductRepository
	shoppingCartService Interface.IShoppingCartService
	OrderDetailService  Interface.IOrderDetailService
	TransactionService  Interface.ITransactionService
	momoService         Interface.IMoMoService
	VnpayService        Interface.IVnPayService
	mailService         Interface.IMailService
	userService         Interface.IUserService
}

func NewOrderService(orderRepository Interface.IOrderRepository, cartItemRepository Interface.ICartItemRepository, productRepository Interface.IProductRepository, shoppingCartService Interface.IShoppingCartService, OrderDetailService Interface.IOrderDetailService, TransactionService Interface.ITransactionService, momoService Interface.IMoMoService, VnpayService Interface.IVnPayService, mailService Interface.IMailService, userService Interface.IUserService) Interface.IOrderService {
	return &OrderService{
		orderRepository:     orderRepository,
		cartItemRepository:  cartItemRepository,
		productRepository:   productRepository,
		shoppingCartService: shoppingCartService,
		OrderDetailService:  OrderDetailService,
		TransactionService:  TransactionService,
		momoService:         momoService,
		VnpayService:        VnpayService,
		mailService:         mailService,
		userService:         userService,
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

func (o *OrderService) CreateOrder(order BusinessObjects.NewOrder) (BusinessObjects.Order, error) {
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
		return BusinessObjects.Order{}, err
	}

	return newOrder, nil
}

func (o *OrderService) PlaceOrder(userId, cartId, shipAddress, CourierID, VoucherID, paymentMethod string) (string, error) {
	// Step 1: Create the order synchronously
	order, err := o.ProcessOrder(userId, cartId, shipAddress, CourierID, VoucherID, paymentMethod)
	if err != nil {
		return "", err
	}

	err = o.PublishInventoryUpdateEvent(userId, cartId)
	if err != nil {
		return "", err
	}

	// Publish Notification Event
	err = o.PublishOrderNotificationEvent(order.OrderID)
	if err != nil {
		return "", err
	}

	paymentURL, err := o.ProcessPayment(order)
	if err != nil {
		return "", err
	}

	return paymentURL, nil
}

func (o *OrderService) ProcessOrder(userId, cartId, shipAddress, CourierID, VoucherID, paymentMethod string) (BusinessObjects.Order, error) {
	productsList, err := o.cartItemRepository.GetCartItemByCartID(cartId)
	if err != nil {
		return BusinessObjects.Order{}, err
	}

	totalAmount := 0.0
	for _, product := range productsList {
		p, err := o.productRepository.GetProductByID(product.ProductID)
		if err != nil {
			return BusinessObjects.Order{}, err
		}

		totalAmount += p.Price * float64(product.Quantity)
	}

	newOrder := BusinessObjects.NewOrder{
		CustomerID:            userId,
		CourierID:             CourierID,
		VoucherID:             VoucherID,
		TotalAmount:           totalAmount,
		PaymentMethod:         paymentMethod,
		ShippingAddress:       shipAddress,
		FreightPrice:          10000,
		EstimatedDeliveryDate: time.Now(),
		ActualDeliveryDate:    time.Now(),
		PaymentStatus:         "Pending",
	}

	createdOrder, err := o.CreateOrder(newOrder)
	if err != nil {
		return BusinessObjects.Order{}, err
	}

	for _, item := range productsList {
		product, err := o.productRepository.GetProductByID(item.ProductID)
		if err != nil {
			return BusinessObjects.Order{}, err
		}
		err = o.OrderDetailService.CreateOrderDetail(
			createdOrder.OrderID,
			item.ProductID,
			item.Quantity,
			product.Price,
		)
		if err != nil {
			return BusinessObjects.Order{}, err
		}
	}

	return createdOrder, nil
}

func (o *OrderService) ProcessPayment(order BusinessObjects.Order) (string, error) {
	transaction := BusinessObjects.NewTransaction{
		OrderID:       order.OrderID,
		PaymentAmount: order.TotalAmount,
		PaymentStatus: "Pending",
		PaymentMethod: order.PaymentMethod,
	}

	err := o.TransactionService.CreateTransaction(transaction)
	if err != nil {
		return "", err
	}

	if order.PaymentMethod == "MoMo" {
		return o.momoService.CreateMoMoUrl(order.TotalAmount, order.OrderID)
	}

	if order.PaymentMethod == "VnPay" {
		return o.VnpayService.CreateVNPayUrl(order.TotalAmount, order.OrderID)
	}

	return "", nil
}

func (o *OrderService) UpdateInventory(userId, cartId string) error {
	productsList, err := o.cartItemRepository.GetCartItemByCartID(cartId)
	if err != nil {
		return err
	}

	for _, product := range productsList {
		p, err := o.productRepository.GetProductByID(product.ProductID)
		if err != nil {
			return err
		}

		p.Quantity -= product.Quantity
		err = o.productRepository.UpdateProduct(p)
		if err != nil {
			return err
		}
	}

	return nil
}

func (o *OrderService) SendNotification(orderID string) error {
	order, err := o.orderRepository.GetOrderByID(orderID)
	if err != nil {
		return err
	}

	customer, err := o.userService.GetUserByID(order.CustomerID)
	if err != nil {
		return err
	}

	orderDetails, err := o.OrderDetailService.GetOrderDetailByID(order.OrderID)
	if err != nil {
		return err
	}

	o.mailService.SendOrderDetails(customer, order, orderDetails)

	return nil
}

func (o *OrderService) UpdateOrder(order BusinessObjects.Order) error {
	return o.orderRepository.UpdateOrder(order)
}

func (o *OrderService) CancelOrder(id string) error {
	return o.orderRepository.DeleteOrder(id)
}

func (o *OrderService) PublishInventoryUpdateEvent(userId, cartId string) error {
	// Connect to RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"inventory_update_queue",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	// Create the message
	message := map[string]string{
		"userId": userId,
		"cartId": cartId,
	}
	body, _ := json.Marshal(message)

	// Publish the message to RabbitMQ
	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	if err != nil {
		return err
	}

	return nil
}

func (o *OrderService) PublishOrderNotificationEvent(orderId string) error {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"order_notification_queue",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	message := map[string]string{
		"orderId": orderId,
	}
	body, _ := json.Marshal(message)

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	if err != nil {
		return err
	}

	return nil
}

func (o *OrderService) ConsumeMailNotifycation() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"order_notification_queue",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)

			// Deserialize message and process payment
			var message map[string]string
			json.Unmarshal(d.Body, &message)

			orderId := message["orderId"]

			o.SendNotification(orderId)
		}
	}()

	log.Printf("Waiting for messages. To exit press CTRL+C")
	<-forever
}

func (o *OrderService) ConsumeInventoryUpdates() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"inventory_update_queue",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)

			// Deserialize message and update inventory
			var message map[string]string
			json.Unmarshal(d.Body, &message)

			userId := message["userId"]
			cartId := message["cartId"]

			// Call your inventory update logic here
			o.UpdateInventory(userId, cartId)
		}
	}()

	log.Printf("Waiting for messages. To exit press CTRL+C")
	<-forever
}
