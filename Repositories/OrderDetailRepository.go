package Repositories

import (
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Util"

	"github.com/sirupsen/logrus"
)

type OrderDetailRepository struct {
	log *logrus.Logger
}

func NewOrderDetailRepository(log *logrus.Logger) *OrderDetailRepository {
	return &OrderDetailRepository{log}
}

func (o *OrderDetailRepository) GetPaginatedOrderDetailList(searchValue, sortBy, orderId, productId string, pageIndex, pageSize int) (Util.PaginatedList[BusinessObjects.OrderDetail], error) {
	o.log.Infof("Fetching paginated order detail list with searchValue: %s, sortBy: %s, orderId: %s, productId: %s, pageIndex: %d, pageSize: %d", searchValue, sortBy, orderId, productId, pageIndex, pageSize)
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		o.log.Error("Failed to connect to PostgreSQL:", err)
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
		o.log.Error("Failed to fetch paginated order details:", err)
		return Util.PaginatedList[BusinessObjects.OrderDetail]{}, err
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		o.log.Error("Failed to count order details:", err)
		return Util.PaginatedList[BusinessObjects.OrderDetail]{}, err
	}

	o.log.Infof("Successfully fetched paginated order detail list with total count: %d", total)
	return Util.NewPaginatedList(rates, total, pageIndex, pageSize), nil
}

// GetAllOrderDetails retrieves all order details
func (o *OrderDetailRepository) GetAllOrderDetails() ([]BusinessObjects.OrderDetail, error) {
	o.log.Info("Fetching all order details")
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		o.log.Error("Failed to connect to PostgreSQL:", err)
		return nil, err
	}

	var rates []BusinessObjects.OrderDetail
	if err := db.Find(&rates).Error; err != nil {
		o.log.Error("Failed to fetch all order details:", err)
		return nil, err
	}

	o.log.Info("Successfully fetched all order details")
	return rates, nil
}

// GetOrderDetailByID retrieves all products of an order
func (o *OrderDetailRepository) GetOrderDetailByID(orderID string) ([]BusinessObjects.OrderDetail, error) {
	o.log.Infof("Fetching order details by order ID: %s", orderID)
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		o.log.Error("Failed to connect to PostgreSQL:", err)
		return []BusinessObjects.OrderDetail{}, err
	}

	var orderDetails []BusinessObjects.OrderDetail
	if err := db.Where("order_id = ?", orderID).Find(&orderDetails).Error; err != nil {
		o.log.Error("Failed to fetch order details by order ID:", err)
		return []BusinessObjects.OrderDetail{}, err
	}

	o.log.Infof("Successfully fetched order details by order ID: %s", orderID)
	return orderDetails, nil
}

// CreateOrderDetail adds a new order detail to the database
func (o *OrderDetailRepository) CreateOrderDetail(rate BusinessObjects.OrderDetail) error {
	o.log.Infof("Creating new order detail with order ID: %s", rate.OrderID)
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		o.log.Error("Failed to connect to PostgreSQL:", err)
		return err
	}

	if err := db.Create(&rate).Error; err != nil {
		o.log.Error("Failed to create new order detail:", err)
		return err
	}

	o.log.Infof("Successfully created new order detail with order ID: %s", rate.OrderID)
	return nil
}

// UpdateOrderDetail updates an existing order detail in the database
func (o *OrderDetailRepository) UpdateOrderDetail(rate BusinessObjects.OrderDetail) error {
	o.log.Infof("Updating order detail with order ID: %s", rate.OrderID)
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		o.log.Error("Failed to connect to PostgreSQL:", err)
		return err
	}

	if err := db.Save(&rate).Error; err != nil {
		o.log.Error("Failed to update order detail:", err)
		return err
	}

	o.log.Infof("Successfully updated order detail with order ID: %s", rate.OrderID)
	return nil
}

// DeleteOrderDetail removes an order detail from the database
func (o *OrderDetailRepository) DeleteOrderDetail(rateID string) error {
	o.log.Infof("Deleting order detail with order ID: %s", rateID)
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		o.log.Error("Failed to connect to PostgreSQL:", err)
		return err
	}

	if err := db.Delete(&BusinessObjects.OrderDetail{}, "order_id = ?", rateID).Error; err != nil {
		o.log.Error("Failed to delete order detail:", err)
		return err
	}

	o.log.Infof("Successfully deleted order detail with order ID: %s", rateID)
	return nil
}
