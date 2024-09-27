package Repositories

import (
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Interface"
	"th3y3m/e-commerce-platform/Util"

	"github.com/sirupsen/logrus"
)

type CourierRepository struct {
	log *logrus.Logger
}

func NewCourierRepository(log *logrus.Logger) Interface.ICourierRepository {
	return &CourierRepository{log}
}

func (c *CourierRepository) GetPaginatedCourierList(searchValue, sortBy string, pageIndex, pageSize int, status *bool) (Util.PaginatedList[BusinessObjects.Courier], error) {
	c.log.Infof("Fetching paginated courier list with searchValue: %s, sortBy: %s, pageIndex: %d, pageSize: %d, status: %v", searchValue, sortBy, pageIndex, pageSize, status)
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		c.log.Error("Failed to connect to PostgreSQL:", err)
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
		c.log.Error("Failed to fetch paginated couriers:", err)
		return Util.PaginatedList[BusinessObjects.Courier]{}, err
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		c.log.Error("Failed to count couriers:", err)
		return Util.PaginatedList[BusinessObjects.Courier]{}, err
	}

	c.log.Infof("Successfully fetched paginated courier list with total count: %d", total)
	return Util.NewPaginatedList(couriers, total, pageIndex, pageSize), nil
}

// GetAllCouriers retrieves all couriers from the database
func (c *CourierRepository) GetAllCouriers() ([]BusinessObjects.Courier, error) {
	c.log.Info("Fetching all couriers")
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		c.log.Error("Failed to connect to PostgreSQL:", err)
		return nil, err
	}

	var couriers []BusinessObjects.Courier
	if err := db.Find(&couriers).Error; err != nil {
		c.log.Error("Failed to fetch all couriers:", err)
		return nil, err
	}

	c.log.Info("Successfully fetched all couriers")
	return couriers, nil
}

// GetCourierByID retrieves a courier by its ID
func (c *CourierRepository) GetCourierByID(courierID string) (BusinessObjects.Courier, error) {
	c.log.Infof("Fetching courier by ID: %s", courierID)
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		c.log.Error("Failed to connect to PostgreSQL:", err)
		return BusinessObjects.Courier{}, err
	}

	var courier BusinessObjects.Courier
	if err := db.First(&courier, "courier_id = ?", courierID).Error; err != nil {
		c.log.Error("Failed to fetch courier by ID:", err)
		return BusinessObjects.Courier{}, err
	}

	c.log.Infof("Successfully fetched courier by ID: %s", courierID)
	return courier, nil
}

// CreateCourier adds a new courier to the database
func (c *CourierRepository) CreateCourier(courier BusinessObjects.Courier) error {
	c.log.Infof("Creating new courier with name: %s", courier.Courier)
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		c.log.Error("Failed to connect to PostgreSQL:", err)
		return err
	}

	if err := db.Create(&courier).Error; err != nil {
		c.log.Error("Failed to create new courier:", err)
		return err
	}

	c.log.Infof("Successfully created new courier with name: %s", courier.Courier)
	return nil
}

// UpdateCourier updates an existing courier
func (c *CourierRepository) UpdateCourier(courier BusinessObjects.Courier) error {
	c.log.Infof("Updating courier with ID: %s", courier.CourierID)
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		c.log.Error("Failed to connect to PostgreSQL:", err)
		return err
	}

	if err := db.Save(&courier).Error; err != nil {
		c.log.Error("Failed to update courier:", err)
		return err
	}

	c.log.Infof("Successfully updated courier with ID: %s", courier.CourierID)
	return nil
}

// DeleteCourier removes a courier from the database by its ID
func (c *CourierRepository) DeleteCourier(courierID string) error {
	c.log.Infof("Deleting courier with ID: %s", courierID)
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		c.log.Error("Failed to connect to PostgreSQL:", err)
		return err
	}

	// if err := db.Delete(&BusinessObjects.Courier{}, "courier_id = ?", courierID).Error; err != nil {
	// 	return err
	// }

	// Set status to false
	if err := db.Model(&BusinessObjects.Courier{}).Where("courier_id = ?", courierID).Update("status", false).Error; err != nil {
		c.log.Error("Failed to delete courier:", err)
		return err
	}

	c.log.Infof("Successfully deleted courier with ID: %s", courierID)
	return nil
}
