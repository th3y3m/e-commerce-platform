package Repositories

import (
	"log"
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Util"
)

func GetAllOrder() ([]BusinessObjects.Order, error) {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM Orders")
	if err != nil {
		log.Fatalf("Error querying the database: %v", err)
		return nil, err
	}
	defer rows.Close()

	orders := []BusinessObjects.Order{}
	for rows.Next() {
		var order BusinessObjects.Order
		err := rows.Scan(&order.OrderID, &order.CustomerID, &order.OrderDate, &order.TotalAmount, &order.OrderStatus, &order.ShippingAddress, &order.CourierID, &order.FreightPrice, &order.EstimatedDeliveryDate, &order.ActualDeliveryDate, &order.PaymentMethod, &order.PaymentStatus, &order.VoucherID)
		if err != nil {
			log.Fatalf("Error scanning row: %v", err)
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}

func GetOrderById(id string) (BusinessObjects.Order, error) {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return BusinessObjects.Order{}, err
	}
	defer db.Close()

	var order BusinessObjects.Order

	err = db.QueryRow("SELECT * FROM Orders WHERE OrderID = ?", id).Scan(&order.OrderID, &order.CustomerID, &order.OrderDate, &order.ShippingAddress, &order.TotalAmount)
	if err != nil {
		log.Fatalf("Error querying the database: %v", err)
		return BusinessObjects.Order{}, err
	}

	return order, nil
}

func CreateOrder(order BusinessObjects.Order) error {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return nil
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO Orders (OrderID, CustomerID, OrderDate, TotalAmount, OrderStatus, ShippingAddress, CourierID, FreightPrice, EstimatedDeliveryDate, ActualDeliveryDate, PaymentMethod, PaymentStatus, VoucherID) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", order.OrderID, order.CustomerID, order.OrderDate, order.TotalAmount, order.OrderStatus, order.ShippingAddress, order.CourierID, order.FreightPrice, order.EstimatedDeliveryDate, order.ActualDeliveryDate, order.PaymentMethod, order.PaymentStatus, order.VoucherID)

	if err != nil {
		log.Fatalf("Error inserting into the database: %v", err)
		return err
	}

	return nil
}

func UpdateOrder(order BusinessObjects.Order) error {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return nil
	}
	defer db.Close()

	_, err = db.Exec("UPDATE Orders SET OrderDate = ?, TotalAmount = ?, OrderStatus = ?, ShippingAddress = ?, CourierID = ?, FreightPrice = ?, EstimatedDeliveryDate = ?, ActualDeliveryDate = ?, PaymentMethod = ?, PaymentStatus = ?, VoucherID = ? WHERE OrderID = ?", order.OrderDate, order.TotalAmount, order.OrderStatus, order.ShippingAddress, order.CourierID, order.FreightPrice, order.EstimatedDeliveryDate, order.ActualDeliveryDate, order.PaymentMethod, order.PaymentStatus, order.VoucherID, order.OrderID)

	if err != nil {
		log.Fatalf("Error updating the database: %v", err)
		return err
	}

	return nil
}

func DeleteOrder(id string) error {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return nil
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM Orders WHERE OrderID = ?", id)
	if err != nil {
		log.Fatalf("Error deleting from the database: %v", err)
		return err
	}

	return nil
}
