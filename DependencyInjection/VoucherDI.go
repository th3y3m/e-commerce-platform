package DependencyInjection

import (
	"th3y3m/e-commerce-platform/Interface"
	"th3y3m/e-commerce-platform/Repositories"
	"th3y3m/e-commerce-platform/Services"

	"github.com/sirupsen/logrus"
)

func NewVoucherRepositoryProvider() Interface.IVoucherRepository {
	log := logrus.New()
	return Repositories.NewVoucherRepository(log)
}

func NewVoucherServiceProvider() Interface.IVoucherService {
	voucherRepository := NewVoucherRepositoryProvider()
	return Services.NewVoucherService(voucherRepository)
}
