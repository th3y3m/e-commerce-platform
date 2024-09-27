package DependencyInjection

import (
	"th3y3m/e-commerce-platform/Interface"
	"th3y3m/e-commerce-platform/Services"
)

func NewVnpayServiceProvider() Interface.IVnPayService {
	transactionRepository := NewTransactionRepositoryProvider()
	orderRepository := NewOrderRepositoryProvider()
	return Services.NewVnpayService(transactionRepository, orderRepository)
}
