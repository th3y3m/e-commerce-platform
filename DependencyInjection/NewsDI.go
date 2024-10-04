package DependencyInjection

import (
	"th3y3m/e-commerce-platform/Interface"
	"th3y3m/e-commerce-platform/Repositories"
	"th3y3m/e-commerce-platform/Services"

	"github.com/sirupsen/logrus"
)

func NewNewsRepositoryProvider() Interface.INewsRepository {
	log := logrus.New()
	db := NewDbProvider()
	return Repositories.NewNewsRepository(log, db)
}

func NewNewsServiceProvider() Interface.INewsService {
	repository := NewNewsRepositoryProvider()
	return Services.NewNewsService(repository)
}
