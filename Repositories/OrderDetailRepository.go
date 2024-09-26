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

	return Util.NewPaginatedList(rates, total, pageIndex, pageSize), nil
}

// GetAllOrderDetails retrieves all order details
func (o *OrderDetailRepository) GetAllOrderDetails() ([]BusinessObjects.OrderDetail, error) {
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

// GetOrderDetailByID retrieves all products of a order
func (o *OrderDetailRepository) GetOrderDetailByID(orderID string) ([]BusinessObjects.OrderDetail, error) {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return []BusinessObjects.OrderDetail{}, err
	}

	var orderDetails []BusinessObjects.OrderDetail
	if err := db.Where("order_id = ?", orderID).Find(&orderDetails).Error; err != nil {
		return []BusinessObjects.OrderDetail{}, err
	}

	return orderDetails, nil
}

// CreateOrderDetail adds a new order detail to the database
func (o *OrderDetailRepository) CreateOrderDetail(rate BusinessObjects.OrderDetail) error {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return err
	}

	if err := db.Create(&rate).Error; err != nil {
		return err
	}

	return nil
}

// UpdateOrderDetail updates an existing order detail in the database
func (o *OrderDetailRepository) UpdateOrderDetail(rate BusinessObjects.OrderDetail) error {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return err
	}

	if err := db.Save(&rate).Error; err != nil {
		return err
	}

	return nil
}

// DeleteOrderDetail removes a order detail from the database
func (o *OrderDetailRepository) DeleteOrderDetail(rateID string) error {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return err
	}

	if err := db.Delete(&BusinessObjects.OrderDetail{}, "order_id = ?", rateID).Error; err != nil {
		return err
	}

	return nil
}
