package Repositories

import (
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Util"
)

func GetPaginatedOrderDetailList(searchValue, sortBy, orderId, productId string, pageIndex, pageSize int) (Util.PaginatedList[BusinessObjects.OrderDetail], error) {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return Util.PaginatedList[BusinessObjects.OrderDetail]{}, err
	}

	var rates []BusinessObjects.OrderDetail
	query := db.Model(&BusinessObjects.OrderDetail{})

	if orderId != "" {
		query = query.Where("order_id = ?", orderId)
	}

	if productId != "" {
		query = query.Where("product_id = ?", productId)
	}

	if searchValue != "" {
		query = query.Where("order_id LIKE ?", "%"+searchValue+"%")
	}

	switch sortBy {
	case "order_id_asc":
		query = query.Order("order_id ASC")
	case "order_id_desc":
		query = query.Order("order_id DESC")
	case "product_id_asc":
		query = query.Order("product_id ASC")
	case "product_id_desc":
		query = query.Order("product_id DESC")
	case "quantity_asc":
		query = query.Order("quantity ASC")
	case "quantity_desc":
		query = query.Order("quantity DESC")
	case "unit_price_asc":
		query = query.Order("unit_price ASC")
	case "unit_price_desc":
		query = query.Order("unit_price DESC")
	default:
		query = query.Order("quantity DESC")
	}

	if err := query.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&rates).Error; err != nil {
		return Util.PaginatedList[BusinessObjects.OrderDetail]{}, err
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return Util.PaginatedList[BusinessObjects.OrderDetail]{}, err
	}

	return Util.PaginatedList[BusinessObjects.OrderDetail]{
		Items:      rates,
		TotalCount: total,
		PageIndex:  pageIndex,
		PageSize:   pageSize,
	}, nil
}

// GetAllOrderDetails retrieves all freight rates from the database
func GetAllOrderDetails() ([]BusinessObjects.OrderDetail, error) {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return nil, err
	}

	var rates []BusinessObjects.OrderDetail
	if err := db.Find(&rates).Error; err != nil {
		return nil, err
	}

	return rates, nil
}

// GetOrderDetailByID retrieves a freight rate by its ID
func GetOrderDetailByID(rateID string) (BusinessObjects.OrderDetail, error) {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return BusinessObjects.OrderDetail{}, err
	}

	var rate BusinessObjects.OrderDetail
	if err := db.First(&rate, "order_id = ?", rateID).Error; err != nil {
		return BusinessObjects.OrderDetail{}, err
	}

	return rate, nil
}

// CreateOrderDetail adds a new freight rate to the database
func CreateOrderDetail(rate BusinessObjects.OrderDetail) error {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return err
	}

	if err := db.Create(&rate).Error; err != nil {
		return err
	}

	return nil
}

// UpdateOrderDetail updates an existing freight rate
func UpdateOrderDetail(rate BusinessObjects.OrderDetail) error {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return err
	}

	if err := db.Save(&rate).Error; err != nil {
		return err
	}

	return nil
}

// DeleteOrderDetail removes a freight rate from the database by its ID
func DeleteOrderDetail(rateID string) error {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return err
	}

	if err := db.Delete(&BusinessObjects.OrderDetail{}, "order_id = ?", rateID).Error; err != nil {
		return err
	}

	return nil
}
