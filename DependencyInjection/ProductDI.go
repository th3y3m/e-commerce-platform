package DependencyInjection

import (
	"th3y3m/e-commerce-platform/Interface"
	"th3y3m/e-commerce-platform/Repositories"
	"th3y3m/e-commerce-platform/Services"

	"github.com/sirupsen/logrus"
)

func NewProductRepositoryProvider() Interface.IProductRepository {
	log := logrus.New()
	return Repositories.NewProductRepository(log)
}

func NewProductServiceProvider() Interface.IProductService {
	productRepository := NewProductRepositoryProvider()
	discountService := NewDiscountServiceProvider()
	productDiscountService := NewProductDiscountServiceProvider()
	return Services.NewProductService(productRepository, discountService, productDiscountService)
}
