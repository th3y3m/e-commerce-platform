package DependencyInjection

import (
	"th3y3m/e-commerce-platform/Interface"
	"th3y3m/e-commerce-platform/Repositories"
	"th3y3m/e-commerce-platform/Services"

	"github.com/sirupsen/logrus"
)

func NewDiscountRepositoryProvider() Interface.IDiscountRepository {
	log := logrus.New()
	return Repositories.NewDiscountRepository(log)
}

func NewDiscountServiceProvider() Interface.IDiscountService {
	discountRepository := NewDiscountRepositoryProvider()
	return Services.NewDiscountService(discountRepository)
}
