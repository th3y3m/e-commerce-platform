package DependencyInjection

import (
	"th3y3m/e-commerce-platform/Interface"
	"th3y3m/e-commerce-platform/Repositories"
	"th3y3m/e-commerce-platform/Services"
	"th3y3m/e-commerce-platform/Util"

	"github.com/sirupsen/logrus"
)

func NewProductRepositoryProvider() Interface.IProductRepository {
	log := logrus.New()
	db, err := Util.ConnectToPostgreSQL()

	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}
	return Repositories.NewProductRepository(log, db)
}

func NewProductServiceProvider() Interface.IProductService {
	productRepository := NewProductRepositoryProvider()
	discountService := NewDiscountServiceProvider()
	productDiscountService := NewProductDiscountServiceProvider()
	return Services.NewProductService(productRepository, discountService, productDiscountService)
}
