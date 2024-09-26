package Interface

import (
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Util"
)

type ICourierRepository interface {
	GetPaginatedCourierList(searchValue, sortBy string, pageIndex, pageSize int, status *bool) (Util.PaginatedList[BusinessObjects.Courier], error)
	GetAllCouriers() ([]BusinessObjects.Courier, error)
	GetCourierByID(courierID string) (BusinessObjects.Courier, error)
	CreateCourier(courier BusinessObjects.Courier) error
	UpdateCourier(courier BusinessObjects.Courier) error
	DeleteCourier(courierID string) error
}

type ICourierService interface {
	GetPaginatedCourierList(searchValue, sortBy string, pageIndex, pageSize int, status *bool) (Util.PaginatedList[BusinessObjects.Courier], error)
	GetAllCouriers() ([]BusinessObjects.Courier, error)
	GetCourierByID(courierID string) (BusinessObjects.Courier, error)
	CreateCourier(CourierName string) error
	UpdateCourier(courierID, CourierName string, status bool) error
	DeleteCourier(id string) error
}
