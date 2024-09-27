package Repositories

import (
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Interface"
	"th3y3m/e-commerce-platform/Util"

	"github.com/sirupsen/logrus"
)

type FreightRateRepository struct {
	log *logrus.Logger
}

func NewFreightRateRepository(log *logrus.Logger) Interface.IFreightRateRepository {
	return &FreightRateRepository{log}
}

func (c *FreightRateRepository) GetPaginatedFreightRateList(searchValue, sortBy, courierID string, pageIndex, pageSize int, status *bool) (Util.PaginatedList[BusinessObjects.FreightRate], error) {
	c.log.Infof("Fetching paginated freight rate list with searchValue: %s, sortBy: %s, courierID: %s, pageIndex: %d, pageSize: %d, status: %v", searchValue, sortBy, courierID, pageIndex, pageSize, status)
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		c.log.Error("Failed to connect to PostgreSQL:", err)
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
		c.log.Error("Failed to count freight rates:", err)
		return Util.PaginatedList[BusinessObjects.FreightRate]{}, err
	}

	if err := query.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&rates).Error; err != nil {
		c.log.Error("Failed to fetch paginated freight rates:", err)
		return Util.PaginatedList[BusinessObjects.FreightRate]{}, err
	}

	c.log.Infof("Successfully fetched paginated freight rate list with total count: %d", total)
	return Util.NewPaginatedList(rates, total, pageIndex, pageSize), nil
}

// GetAllFreightRates retrieves all freight rates from the database
func (c *FreightRateRepository) GetAllFreightRates() ([]BusinessObjects.FreightRate, error) {
	c.log.Info("Fetching all freight rates")
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		c.log.Error("Failed to connect to PostgreSQL:", err)
		return nil, err
	}

	var rates []BusinessObjects.FreightRate
	if err := db.Find(&rates).Error; err != nil {
		c.log.Error("Failed to fetch all freight rates:", err)
		return nil, err
	}

	c.log.Info("Successfully fetched all freight rates")
	return rates, nil
}

// GetFreightRateByID retrieves a freight rate by its ID
func (c *FreightRateRepository) GetFreightRateByID(rateID string) (BusinessObjects.FreightRate, error) {
	c.log.Infof("Fetching freight rate by ID: %s", rateID)
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		c.log.Error("Failed to connect to PostgreSQL:", err)
		return BusinessObjects.FreightRate{}, err
	}

	var rate BusinessObjects.FreightRate
	if err := db.First(&rate, "rate_id = ?", rateID).Error; err != nil {
		c.log.Error("Failed to fetch freight rate by ID:", err)
		return BusinessObjects.FreightRate{}, err
	}

	c.log.Infof("Successfully fetched freight rate by ID: %s", rateID)
	return rate, nil
}

func (c *FreightRateRepository) GetFreightRateByCourierID(courierID string) ([]BusinessObjects.FreightRate, error) {
	c.log.Infof("Fetching freight rates by courier ID: %s", courierID)
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		c.log.Error("Failed to connect to PostgreSQL:", err)
		return nil, err
	}

	var rates []BusinessObjects.FreightRate
	if err := db.Find(&rates, "courier_id = ?", courierID).Error; err != nil {
		c.log.Error("Failed to fetch freight rates by courier ID:", err)
		return nil, err
	}

	c.log.Infof("Successfully fetched freight rates by courier ID: %s", courierID)
	return rates, nil
}

// CreateFreightRate adds a new freight rate to the database
func (c *FreightRateRepository) CreateFreightRate(rate BusinessObjects.FreightRate) error {
	c.log.Infof("Creating new freight rate with ID: %s", rate.RateID)
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		c.log.Error("Failed to connect to PostgreSQL:", err)
		return err
	}

	if err := db.Create(&rate).Error; err != nil {
		c.log.Error("Failed to create new freight rate:", err)
		return err
	}

	c.log.Infof("Successfully created new freight rate with ID: %s", rate.RateID)
	return nil
}

// UpdateFreightRate updates an existing freight rate
func (c *FreightRateRepository) UpdateFreightRate(rate BusinessObjects.FreightRate) error {
	c.log.Infof("Updating freight rate with ID: %s", rate.RateID)
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		c.log.Error("Failed to connect to PostgreSQL:", err)
		return err
	}

	if err := db.Save(&rate).Error; err != nil {
		c.log.Error("Failed to update freight rate:", err)
		return err
	}

	c.log.Infof("Successfully updated freight rate with ID: %s", rate.RateID)
	return nil
}

// DeleteFreightRate removes a freight rate from the database by its ID
func (c *FreightRateRepository) DeleteFreightRate(rateID string) error {
	c.log.Infof("Deleting freight rate with ID: %s", rateID)
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		c.log.Error("Failed to connect to PostgreSQL:", err)
		return err
	}

	// if err := db.Delete(&BusinessObjects.FreightRate{}, "rate_id = ?", rateID).Error; err != nil {
	// 	return err
	// }

	// Set status to false instead of deleting
	if err := db.Model(&BusinessObjects.FreightRate{}).Where("rate_id = ?", rateID).Update("status", false).Error; err != nil {
		c.log.Error("Failed to delete freight rate:", err)
		return err
	}

	c.log.Infof("Successfully deleted freight rate with ID: %s", rateID)
	return nil
}
