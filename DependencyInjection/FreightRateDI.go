package DependencyInjection

import (
	"th3y3m/e-commerce-platform/Interface"
	"th3y3m/e-commerce-platform/Repositories"
	"th3y3m/e-commerce-platform/Services"

	"github.com/sirupsen/logrus"
)

func NewFreightRateRepositoryProvider() Interface.IFreightRateRepository {
	log := logrus.New()
	return Repositories.NewFreightRateRepository(log)
}

func NewFreightRateServiceProvider() Interface.IFreightRateService {
	freightRateRepository := NewFreightRateRepositoryProvider()
	return Services.NewFreightRateService(freightRateRepository)
}
