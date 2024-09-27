package DependencyInjection

import (
	"th3y3m/e-commerce-platform/Interface"
	"th3y3m/e-commerce-platform/Repositories"
	"th3y3m/e-commerce-platform/Services"
	"th3y3m/e-commerce-platform/Util"

	"github.com/sirupsen/logrus"
)

func NewCourierRepositoryProvider() Interface.ICourierRepository {
	log := logrus.New()
	db, err := Util.ConnectToPostgreSQL()

	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}
	return Repositories.NewCourierRepository(log, db)
}

func NewCourierServiceProvider() Interface.ICourierService {
	courierRepository := NewCourierRepositoryProvider()
	return Services.NewCourierService(courierRepository)
}
