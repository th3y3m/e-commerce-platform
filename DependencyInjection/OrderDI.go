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

	return Services.NewOrderService(orderRepository, cartItemRepository, productRepository, shoppingCartService)
}
