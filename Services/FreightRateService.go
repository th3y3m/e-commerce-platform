package Services

import (
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Repositories"
	"th3y3m/e-commerce-platform/Util"
)

func GetPaginatedFreightRateList(searchValue, sortBy, courierId string, pageIndex, pageSize int, status *bool) (Util.PaginatedList[BusinessObjects.FreightRate], error) {
	return Repositories.GetPaginatedFreightRateList(searchValue, sortBy, courierId, pageIndex, pageSize, status)
}

func GetAllFreightRates() ([]BusinessObjects.FreightRate, error) {
	return Repositories.GetAllFreightRates()
}

func GetFreightRateByID(id string) (BusinessObjects.FreightRate, error) {
	return Repositories.GetFreightRateByID(id)
}

func CreateFreightRate(courierId string, distanceMinKm, distanceMaxKm int, costPerKm float64) error {
	freightRate := BusinessObjects.FreightRate{
		RateID:        "FR" + Util.GenerateID(10),
		CourierID:     courierId,
		DistanceMinKM: distanceMinKm,
		DistanceMaxKM: distanceMaxKm,
		CostPerKM:     costPerKm,
		Status:        true,
	}

	err := Repositories.CreateFreightRate(freightRate)
	if err != nil {
		return err
	}

	return nil
}

func UpdateFreightRate(rateID, courierId string, distanceMinKm, distanceMaxKm int, costPerKm float64, status bool) error {
	freightRate, err := Repositories.GetFreightRateByID(rateID)
	if err != nil {
		return err
	}

	freightRate.CourierID = courierId
	freightRate.DistanceMinKM = distanceMinKm
	freightRate.DistanceMaxKM = distanceMaxKm
	freightRate.CostPerKM = costPerKm
	freightRate.Status = status

	return Repositories.UpdateFreightRate(freightRate)
}

func DeleteFreightRate(id string) error {
	return Repositories.DeleteFreightRate(id)
}
