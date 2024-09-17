package Repositories

import (
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Util"
)

func GetPaginatedFreightRateList(searchValue, sortBy, courierID string, pageIndex, pageSize int, status *bool) (Util.PaginatedList[BusinessObjects.FreightRate], error) {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return Util.PaginatedList[BusinessObjects.FreightRate]{}, err
	}

	var rates []BusinessObjects.FreightRate
	query := db.Model(&BusinessObjects.FreightRate{})

	if courierID != "" {
		query = query.Where("courier_id = ?", courierID)
	}

	if searchValue != "" {
		query = query.Where("rate_id LIKE ?", "%"+searchValue+"%")
	}

	if status != nil {
		query = query.Where("status = ?", *status)
	}

	switch sortBy {
	case "rate_id_asc":
		query = query.Order("rate_id ASC")
	case "rate_id_desc":
		query = query.Order("rate_id DESC")
	case "courier_id_asc":
		query = query.Order("courier_id ASC")
	case "courier_id_desc":
		query = query.Order("courier_id DESC")
	case "distance_min_km_asc":
		query = query.Order("distance_min_km ASC")
	case "distance_min_km_desc":
		query = query.Order("distance_min_km DESC")
	case "distance_max_km_asc":
		query = query.Order("distance_max_km ASC")
	case "distance_max_km_desc":
		query = query.Order("distance_max_km DESC")
	case "cost_per_km_asc":
		query = query.Order("cost_per_km ASC")
	case "cost_per_km_desc":
		query = query.Order("cost_per_km DESC")
	case "status_asc":
		query = query.Order("status ASC")
	case "status_desc":
		query = query.Order("status DESC")
	default:
		query = query.Order("distance_min_km ASC")
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return Util.PaginatedList[BusinessObjects.FreightRate]{}, err
	}

	if err := query.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&rates).Error; err != nil {
		return Util.PaginatedList[BusinessObjects.FreightRate]{}, err
	}

	return Util.NewPaginatedList(rates, total, pageIndex, pageSize), nil
}

// GetAllFreightRates retrieves all freight rates from the database
func GetAllFreightRates() ([]BusinessObjects.FreightRate, error) {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return nil, err
	}

	var rates []BusinessObjects.FreightRate
	if err := db.Find(&rates).Error; err != nil {
		return nil, err
	}

	return rates, nil
}

// GetFreightRateByID retrieves a freight rate by its ID
func GetFreightRateByID(rateID string) (BusinessObjects.FreightRate, error) {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return BusinessObjects.FreightRate{}, err
	}

	var rate BusinessObjects.FreightRate
	if err := db.First(&rate, "rate_id = ?", rateID).Error; err != nil {
		return BusinessObjects.FreightRate{}, err
	}

	return rate, nil
}
func GetFreightRateByCourierID(courierID string) ([]BusinessObjects.FreightRate, error) {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return nil, err
	}

	var rates []BusinessObjects.FreightRate
	if err := db.Find(&rates, "courier_id = ?", courierID).Error; err != nil {
		return nil, err
	}

	return rates, nil
}

// CreateFreightRate adds a new freight rate to the database
func CreateFreightRate(rate BusinessObjects.FreightRate) error {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return err
	}

	if err := db.Create(&rate).Error; err != nil {
		return err
	}

	return nil
}

// UpdateFreightRate updates an existing freight rate
func UpdateFreightRate(rate BusinessObjects.FreightRate) error {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return err
	}

	if err := db.Save(&rate).Error; err != nil {
		return err
	}

	return nil
}

// DeleteFreightRate removes a freight rate from the database by its ID
func DeleteFreightRate(rateID string) error {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return err
	}

	// if err := db.Delete(&BusinessObjects.FreightRate{}, "rate_id = ?", rateID).Error; err != nil {
	// 	return err
	// }

	//Set status to false instead of deleting

	if err := db.Model(&BusinessObjects.FreightRate{}).Where("rate_id = ?", rateID).Update("status", false).Error; err != nil {
		return err
	}

	return nil
}
