package DependencyInjection

import (
	"th3y3m/e-commerce-platform/Interface"
	"th3y3m/e-commerce-platform/Repositories"
	"th3y3m/e-commerce-platform/Services"

	"github.com/sirupsen/logrus"
)

func NewProductDiscountRepositoryProvider() Interface.IProductDiscountRepository {
	log := logrus.New()
	return Repositories.NewProductDiscountRepository(log)
}

func NewProductDiscountServiceProvider() Interface.IProductDiscountService {
	productDiscountRepository := NewProductDiscountRepositoryProvider()
	return Services.NewProductDiscountService(productDiscountRepository)
}
