package DependencyInjection

import (
	"th3y3m/e-commerce-platform/Interface"
	"th3y3m/e-commerce-platform/Repositories"
	"th3y3m/e-commerce-platform/Services"

	"github.com/sirupsen/logrus"
)

func NewCourierRepositoryProvider() Interface.ICourierRepository {
	log := logrus.New()
	return Repositories.NewCourierRepository(log)
}

func NewCourierServiceProvider() Interface.ICourierService {
	courierRepository := NewCourierRepositoryProvider()
	return Services.NewCourierService(courierRepository)
}
