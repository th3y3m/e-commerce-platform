package Services

import (
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Interface"
	"th3y3m/e-commerce-platform/Util"
)

type CourierService struct {
	courierRepository Interface.ICourierRepository
}

func NewCourierService(courierRepository Interface.ICourierRepository) Interface.ICourierService {
	return &CourierService{courierRepository}
}

func (c *CourierService) GetPaginatedCourierList(searchValue, sortBy string, pageIndex, pageSize int, status *bool) (Util.PaginatedList[BusinessObjects.Courier], error) {
	return c.courierRepository.GetPaginatedCourierList(searchValue, sortBy, pageIndex, pageSize, status)
}

func (c *CourierService) GetAllCouriers() ([]BusinessObjects.Courier, error) {
	return c.courierRepository.GetAllCouriers()
}

func (c *CourierService) GetCourierByID(id string) (BusinessObjects.Courier, error) {
	return c.courierRepository.GetCourierByID(id)
}

func (c *CourierService) CreateCourier(CourierName string) error {
	courier := BusinessObjects.Courier{
		CourierID: "COUR" + Util.GenerateID(10),
		Courier:   CourierName,
		Status:    true,
	}

	err := c.courierRepository.CreateCourier(courier)
	if err != nil {
		return err
	}

	return nil
}

func (c *CourierService) UpdateCourier(courierID, CourierName string, status bool) error {
	courier, err := c.courierRepository.GetCourierByID(courierID)
	if err != nil {
		return err
	}

	courier.Courier = CourierName
	courier.Status = status

	return c.courierRepository.UpdateCourier(courier)
}

func (c *CourierService) DeleteCourier(id string) error {
	return c.courierRepository.DeleteCourier(id)
}
