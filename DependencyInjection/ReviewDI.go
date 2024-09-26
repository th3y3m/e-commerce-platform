package DependencyInjection

import (
	"th3y3m/e-commerce-platform/Interface"
	"th3y3m/e-commerce-platform/Repositories"
	"th3y3m/e-commerce-platform/Services"

	"github.com/sirupsen/logrus"
)

func NewReviewRepositoryProvider() Interface.IReviewRepository {
	log := logrus.New()
	return Repositories.NewReviewRepository(log)
}

func NewReviewServiceProvider() Interface.IReviewService {
	reviewRepository := NewReviewRepositoryProvider()
	return Services.NewReviewService(reviewRepository)
}
