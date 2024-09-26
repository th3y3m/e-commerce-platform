package Interface

import (
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Util"
)

type IFreightRateRepository interface {
	GetPaginatedFreightRateList(searchValue, sortBy, courierID string, pageIndex, pageSize int, status *bool) (Util.PaginatedList[BusinessObjects.FreightRate], error)
	GetAllFreightRates() ([]BusinessObjects.FreightRate, error)
	GetFreightRateByID(rateID string) (BusinessObjects.FreightRate, error)
	GetFreightRateByCourierID(courierID string) ([]BusinessObjects.FreightRate, error)
	CreateFreightRate(rate BusinessObjects.FreightRate) error
	UpdateFreightRate(rate BusinessObjects.FreightRate) error
	DeleteFreightRate(rateID string) error
}

type IFreightRateService interface {
	GetPaginatedFreightRateList(searchValue, sortBy, courierID string, pageIndex, pageSize int, status *bool) (Util.PaginatedList[BusinessObjects.FreightRate], error)
	GetAllFreightRates() ([]BusinessObjects.FreightRate, error)
	GetFreightRateByID(rateID string) (BusinessObjects.FreightRate, error)
	CreateFreightRate(courierId string, distanceMinKm, distanceMaxKm int, costPerKm float64) error
	UpdateFreightRate(rateID, courierId string, distanceMinKm, distanceMaxKm int, costPerKm float64, status bool) error
	DeleteFreightRate(rateID string) error
	GetFreightRateByCourierID(id string) ([]BusinessObjects.FreightRate, error)
	CalculateFreightRate(courierId string, distance float64) (float64, error)
}
