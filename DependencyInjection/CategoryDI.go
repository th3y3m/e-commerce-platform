package DependencyInjection

import (
	"th3y3m/e-commerce-platform/Interface"
	"th3y3m/e-commerce-platform/Repositories"
	"th3y3m/e-commerce-platform/Services"

	"github.com/sirupsen/logrus"
)

func NewCategoryRepositoryProvider() Interface.ICategoryRepository {
	log := logrus.New()
	return Repositories.NewCategoryRepository(log)
}

func NewCategoryServiceProvider() Interface.ICategoryService {
	categoryRepository := NewCategoryRepositoryProvider()
	return Services.NewCategoryService(categoryRepository)
}
