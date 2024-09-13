package Services

import (
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Repositories"
	"th3y3m/e-commerce-platform/Util"
)

func GetPaginatedOrderDetailList(searchValue, sortBy, orderId, productId string, pageIndex, pageSize int) (Util.PaginatedList[BusinessObjects.OrderDetail], error) {
	return Repositories.GetPaginatedOrderDetailList(searchValue, sortBy, orderId, productId, pageIndex, pageSize)
}

func GetAllOrderDetails() ([]BusinessObjects.OrderDetail, error) {
	return Repositories.GetAllOrderDetails()
}

func GetOrderDetailByID(id string) (BusinessObjects.OrderDetail, error) {
	return Repositories.GetOrderDetailByID(id)
}

func CreateOrderDetail(orderId, productId string, quantity int, unitPrice float64) error {
	orderDetail := BusinessObjects.OrderDetail{
		OrderID:   orderId,
		ProductID: productId,
		Quantity:  quantity,
		UnitPrice: unitPrice,
	}

	return Repositories.CreateOrderDetail(orderDetail)
}

func UpdateOrderDetail(orderDetail BusinessObjects.OrderDetail) error {
	return Repositories.UpdateOrderDetail(orderDetail)
}

func DeleteOrderDetail(id string) error {
	return Repositories.DeleteOrderDetail(id)
}
