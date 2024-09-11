package Repositories

import (
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Util"
	"time"
)

func GetPaginatedOrderList(sortBy, orderID, customerId, courierId, voucherId string, pageIndex, pageSize int, startDate, endDate time.Time) (Util.PaginatedList[BusinessObjects.Order], error) {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return Util.PaginatedList[BusinessObjects.Order]{}, err
	}

	var orders []BusinessObjects.Order
	query := db.Model(&BusinessObjects.Order{})

	if customerId != "" {
		query = query.Where("customer_id = ?", customerId)
	}

	if courierId != "" {
		query = query.Where("courier_id = ?", courierId)
	}

	if orderID != "" {
		query = query.Where("order_id = ?", orderID)
	}

	if voucherId != "" {
		query = query.Where("voucher_id = ?", voucherId)
	}

	if !startDate.IsZero() {
		query = query.Where("order_date >= ?", startDate)
	}

	if !endDate.IsZero() {
		query = query.Where("order_date <= ?", endDate)
	}

	switch sortBy {
	case "order_id_asc":
		query = query.Order("order_id ASC")
	case "order_id_desc":
		query = query.Order("order_id DESC")
	case "customer_id_asc":
		query = query.Order("customer_id ASC")
	case "customer_id_desc":
		query = query.Order("customer_id DESC")
	case "total_price_asc":
		query = query.Order("total_price ASC")
	case "total_price_desc":
		query = query.Order("total_price DESC")
	case "payment_status_asc":
		query = query.Order("payment_status ASC")
	case "payment_status_desc":
		query = query.Order("payment_status DESC")
	case "order_date_asc":
		query = query.Order("order_date ASC")
	case "order_date_desc":
		query = query.Order("order_date DESC")
	default:
		query = query.Order("order_date DESC")
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return Util.PaginatedList[BusinessObjects.Order]{}, err
	}

	if err := query.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&orders).Error; err != nil {
		return Util.PaginatedList[BusinessObjects.Order]{}, err
	}

	return Util.NewPaginatedList(orders, total, pageIndex, pageSize), nil
}

// GetAllOrders retrieves all orders from the database
func GetAllOrders() ([]BusinessObjects.Order, error) {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return nil, err
	}

	var orders []BusinessObjects.Order
	if err := db.Find(&orders).Error; err != nil {
		return nil, err
	}

	return orders, nil
}

// GetOrderByID retrieves an order by its ID
func GetOrderByID(orderID string) (BusinessObjects.Order, error) {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return BusinessObjects.Order{}, err
	}

	var order BusinessObjects.Order
	if err := db.First(&order, "order_id = ?", orderID).Error; err != nil {
		return BusinessObjects.Order{}, err
	}

	return order, nil
}

// CreateOrder adds a new order to the database
func CreateOrder(order BusinessObjects.Order) error {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return err
	}

	if err := db.Create(&order).Error; err != nil {
		return err
	}

	return nil
}

// UpdateOrder updates an existing order
func UpdateOrder(order BusinessObjects.Order) error {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return err
	}

	if err := db.Save(&order).Error; err != nil {
		return err
	}

	return nil
}

// DeleteOrder removes an order from the database by its ID
func DeleteOrder(orderID string) error {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return err
	}

	if err := db.Delete(&BusinessObjects.Order{}, "order_id = ?", orderID).Error; err != nil {
		return err
	}

	return nil
}
