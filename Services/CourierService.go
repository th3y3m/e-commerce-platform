package Services

import (
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Repositories"
	"th3y3m/e-commerce-platform/Util"
)

func GetPaginatedCourierList(searchValue, sortBy string, pageIndex, pageSize int, status *bool) (Util.PaginatedList[BusinessObjects.Courier], error) {
	return Repositories.GetPaginatedCourierList(searchValue, sortBy, pageIndex, pageSize, status)
}

func GetAllCouriers() ([]BusinessObjects.Courier, error) {
	return Repositories.GetAllCouriers()
}

func GetCourierByID(id string) (BusinessObjects.Courier, error) {
	return Repositories.GetCourierByID(id)
}

func CreateCourier(CourierName string) error {
	courier := BusinessObjects.Courier{
		CourierID: "COUR" + Util.GenerateID(10),
		Courier:   CourierName,
		Status:    true,
	}

	err := Repositories.CreateCourier(courier)
	if err != nil {
		return err
	}

	return nil
}

func UpdateCourier(courierID, CourierName string, status bool) error {
	courier, err := Repositories.GetCourierByID(courierID)
	if err != nil {
		return err
	}

	courier.Courier = CourierName
	courier.Status = status

	return Repositories.UpdateCourier(courier)
}

func DeleteCourier(id string) error {
	return Repositories.DeleteCourier(id)
}
