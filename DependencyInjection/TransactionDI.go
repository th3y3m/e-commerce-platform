package DependencyInjection

import (
	"th3y3m/e-commerce-platform/Interface"
	"th3y3m/e-commerce-platform/Repositories"
	"th3y3m/e-commerce-platform/Services"

	"github.com/sirupsen/logrus"
)

func NewTransactionRepositoryProvider() Interface.ITransactionRepository {
	log := logrus.New()
	return Repositories.NewTransactionRepository(log)
}

func NewTransactionServiceProvider() Interface.ITransactionService {
	transactionRepository := NewTransactionRepositoryProvider()
	return Services.NewTransactionService(transactionRepository)
}
