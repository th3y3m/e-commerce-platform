package Services

import (
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Interface"
	"th3y3m/e-commerce-platform/Util"
)

type FreightRateService struct {
	freightRateRepository Interface.IFreightRateRepository
}

func NewFreightRateService(freightRateRepository Interface.IFreightRateRepository) Interface.IFreightRateService {
	return &FreightRateService{freightRateRepository}
}

func (c *FreightRateService) GetPaginatedFreightRateList(searchValue, sortBy, courierId string, pageIndex, pageSize int, status *bool) (Util.PaginatedList[BusinessObjects.FreightRate], error) {
	return c.freightRateRepository.GetPaginatedFreightRateList(searchValue, sortBy, courierId, pageIndex, pageSize, status)
}

func (c *FreightRateService) GetAllFreightRates() ([]BusinessObjects.FreightRate, error) {
	return c.freightRateRepository.GetAllFreightRates()
}

func (c *FreightRateService) GetFreightRateByID(id string) (BusinessObjects.FreightRate, error) {
	return c.freightRateRepository.GetFreightRateByID(id)
}

func (c *FreightRateService) CreateFreightRate(courierId string, distanceMinKm, distanceMaxKm int, costPerKm float64) error {
	freightRate := BusinessObjects.FreightRate{
		RateID:        "FR" + Util.GenerateID(10),
		CourierID:     courierId,
		DistanceMinKM: distanceMinKm,
		DistanceMaxKM: distanceMaxKm,
		CostPerKM:     costPerKm,
		Status:        true,
	}

	err := c.freightRateRepository.CreateFreightRate(freightRate)
	if err != nil {
		return err
	}

	return nil
}

func (c *FreightRateService) UpdateFreightRate(rateID, courierId string, distanceMinKm, distanceMaxKm int, costPerKm float64, status bool) error {
	freightRate, err := c.freightRateRepository.GetFreightRateByID(rateID)
	if err != nil {
		return err
	}

	freightRate.CourierID = courierId
	freightRate.DistanceMinKM = distanceMinKm
	freightRate.DistanceMaxKM = distanceMaxKm
	freightRate.CostPerKM = costPerKm
	freightRate.Status = status

	return c.freightRateRepository.UpdateFreightRate(freightRate)
}

func (c *FreightRateService) DeleteFreightRate(id string) error {
	return c.freightRateRepository.DeleteFreightRate(id)
}
func (c *FreightRateService) GetFreightRateByCourierID(id string) ([]BusinessObjects.FreightRate, error) {
	return c.freightRateRepository.GetFreightRateByCourierID(id)
}

func (c *FreightRateService) CalculateFreightRate(courierId string, distance float64) (float64, error) {
	rates, err := c.GetFreightRateByCourierID(courierId)
	if err != nil {
		return 0, err
	}

	for _, rate := range rates {
		if distance >= float64(rate.DistanceMinKM) && distance <= float64(rate.DistanceMaxKM) {
			return rate.CostPerKM * distance, nil
		}
	}

	return 0, nil
}
