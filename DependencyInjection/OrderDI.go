package DependencyInjection

import (
	"th3y3m/e-commerce-platform/Interface"
	"th3y3m/e-commerce-platform/Repositories"
	"th3y3m/e-commerce-platform/Services"

	"github.com/sirupsen/logrus"
)

func NewOrderRepositoryProvider() Interface.IOrderRepository {
	log := logrus.New()
	return Repositories.NewOrderRepository(log)
}

func NewOrderServiceProvider() Interface.IOrderService {
	orderRepository := NewOrderRepositoryProvider()
	cartItemRepository := NewCartItemRepositoryProvider()
	productRepository := NewProductRepositoryProvider()
	shoppingCartService := NewShoppingCartServiceProvider()
	orderDetailService := NewOrderDetailServiceProvider()
	transactionService := NewTransactionServiceProvider()
	momoService := NewMoMoServiceProvider()
	vnPayService := NewVnpayServiceProvider()
	mailService := NewMailServiceProvider()
	userService := NewUserServiceProvider()

	return Services.NewOrderService(orderRepository, cartItemRepository, productRepository, shoppingCartService, orderDetailService, transactionService, momoService, vnPayService, mailService, userService)
}
