package DependencyInjection

import (
	"th3y3m/e-commerce-platform/Interface"
	"th3y3m/e-commerce-platform/Services"
)

func NewMoMoServiceProvider() Interface.IMoMoService {
	transactionRepository := NewTransactionRepositoryProvider()
	orderRepository := NewOrderRepositoryProvider()
	return Services.NewMoMoService(transactionRepository, orderRepository)
}
