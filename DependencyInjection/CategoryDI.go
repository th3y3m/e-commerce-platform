package DependencyInjection

import (
	"th3y3m/e-commerce-platform/Interface"
	"th3y3m/e-commerce-platform/Repositories"
	"th3y3m/e-commerce-platform/Services"
	"th3y3m/e-commerce-platform/Util"

	"github.com/sirupsen/logrus"
)

func NewCategoryRepositoryProvider() Interface.ICategoryRepository {
	log := logrus.New()
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		log.Error("Failed to connect to PostgreSQL:", err)
		return nil
	}
	return Repositories.NewCategoryRepository(log, db)
}

func NewCategoryServiceProvider() Interface.ICategoryService {
	categoryRepository := NewCategoryRepositoryProvider()
	return Services.NewCategoryService(categoryRepository)
}
