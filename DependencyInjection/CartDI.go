package DependencyInjection

import (
	"th3y3m/e-commerce-platform/Interface"
	"th3y3m/e-commerce-platform/Repositories"
	"th3y3m/e-commerce-platform/Services"

	"github.com/sirupsen/logrus"
)

func NewShoppingCartRepositoryProvider() Interface.IShoppingCartRepository {
	log := logrus.New()
	return Repositories.NewShoppingCartRepository(log)
}

func NewShoppingCartServiceProvider() Interface.IShoppingCartService {
	shoppingCartRepository := NewShoppingCartRepositoryProvider()
	cartItemRepository := NewCartItemRepositoryProvider()
	return Services.NewShoppingCartService(shoppingCartRepository, cartItemRepository)
}

func NewCartItemRepositoryProvider() Interface.ICartItemRepository {
	log := logrus.New()
	return Repositories.NewCartItemRepository(log)
}

func NewCartItemServiceProvider() Interface.ICartItemService {
	cartItemRepository := NewCartItemRepositoryProvider()
	shoppingRepository := NewShoppingCartRepositoryProvider()
	return Services.NewCartItemService(cartItemRepository, shoppingRepository)
}
