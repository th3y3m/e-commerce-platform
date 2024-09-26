package DependencyInjection

import (
	"th3y3m/e-commerce-platform/Interface"
	"th3y3m/e-commerce-platform/Repositories"
	"th3y3m/e-commerce-platform/Services"

	"github.com/sirupsen/logrus"
)

func NewNewOrderDetailRepositoryProvider() Interface.IOrderDetailRepository {
	log := logrus.New()
	return Repositories.NewOrderDetailRepository(log)
}

func NewOrderDetailServiceProvider() Interface.IOrderDetailService {
	orderDetailRepository := NewNewOrderDetailRepositoryProvider()
	return Services.NewOrderDetailService(orderDetailRepository)
}
