package Repositories

import (
	"log"
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Util"
)

func GetAllProductIdOfOrder(id string) ([]BusinessObjects.OrderDetail, error) {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM OrderDetails WHERE OrderID = ?", id)

	if err != nil {
		log.Fatalf("Error querying the database: %v", err)
		return nil, err
	}

	defer rows.Close()

	products := []BusinessObjects.OrderDetail{}
	for rows.Next() {
		var product BusinessObjects.OrderDetail
		err := rows.Scan(&product.ProductID, &product.OrderID, &product.Quantity, &product.UnitPrice)
		if err != nil {
			log.Fatalf("Error scanning row: %v", err)
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func CreateOrderDetail(orderDetail BusinessObjects.OrderDetail) error {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO OrderDetails (ProductID, OrderID, Quantity, UnitPrice) VALUES (?, ?, ?, ?)", orderDetail.ProductID, orderDetail.OrderID, orderDetail.Quantity, orderDetail.UnitPrice)
	if err != nil {
		log.Fatalf("Error inserting into the database: %v", err)
		return err
	}

	return nil
}

func GetOrderDetailById(orderId string, productId string) (BusinessObjects.OrderDetail, error) {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return BusinessObjects.OrderDetail{}, err
	}
	defer db.Close()

	var orderDetail BusinessObjects.OrderDetail

	err = db.QueryRow("SELECT * FROM OrderDetails WHERE OrderID = ? AND ProductID = ?", orderId, productId).Scan(&orderDetail.ProductID, &orderDetail.OrderID, &orderDetail.Quantity, &orderDetail.UnitPrice)
	if err != nil {
		log.Fatalf("Error querying the database: %v", err)
		return BusinessObjects.OrderDetail{}, err
	}

	return orderDetail, nil
}
