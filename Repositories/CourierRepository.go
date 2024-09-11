package Repositories

import (
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Util"
)

func GetPaginatedCourierList(searchValue, sortBy string, pageIndex, pageSize int, status *bool) (Util.PaginatedList[BusinessObjects.Courier], error) {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return Util.PaginatedList[BusinessObjects.Courier]{}, err
	}

	var couriers []BusinessObjects.Courier
	query := db.Model(&BusinessObjects.Courier{})
	if searchValue != "" {
		query = query.Where("courier LIKE ?", "%"+searchValue+"%")
	}
	if status != nil {
		query = query.Where("status = ?", *status)
	}

	switch sortBy {
	case "courier_asc":
		query = query.Order("courier ASC")
	case "courier_desc":
		query = query.Order("courier DESC")
	default:
		query = query.Order("courier ASC")
	}

	if err := query.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&couriers).Error; err != nil {
		return Util.PaginatedList[BusinessObjects.Courier]{}, err
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return Util.PaginatedList[BusinessObjects.Courier]{}, err
	}

	return Util.NewPaginatedList(couriers, total, pageIndex, pageSize), nil
}

// GetAllCouriers retrieves all couriers from the database
func GetAllCouriers() ([]BusinessObjects.Courier, error) {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return nil, err
	}

	var couriers []BusinessObjects.Courier
	if err := db.Find(&couriers).Error; err != nil {
		return nil, err
	}

	return couriers, nil
}

// GetCourierByID retrieves a courier by its ID
func GetCourierByID(courierID string) (BusinessObjects.Courier, error) {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return BusinessObjects.Courier{}, err
	}

	var courier BusinessObjects.Courier
	if err := db.First(&courier, "courier_id = ?", courierID).Error; err != nil {
		return BusinessObjects.Courier{}, err
	}

	return courier, nil
}

// CreateCourier adds a new courier to the database
func CreateCourier(courier BusinessObjects.Courier) error {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return err
	}

	if err := db.Create(&courier).Error; err != nil {
		return err
	}

	return nil
}

// UpdateCourier updates an existing courier
func UpdateCourier(courier BusinessObjects.Courier) error {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return err
	}

	if err := db.Save(&courier).Error; err != nil {
		return err
	}

	return nil
}

// DeleteCourier removes a courier from the database by its ID
func DeleteCourier(courierID string) error {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return err
	}

	if err := db.Delete(&BusinessObjects.Courier{}, "courier_id = ?", courierID).Error; err != nil {
		return err
	}

	return nil
}
