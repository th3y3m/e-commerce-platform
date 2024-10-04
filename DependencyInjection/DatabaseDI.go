package DependencyInjection

import (
	"th3y3m/e-commerce-platform/Provider"

	"github.com/sirupsen/logrus"
)

func NewDbProvider() Provider.IDb {
	log := logrus.New()
	return Provider.NewDbProvider(log)
}
