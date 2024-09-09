package Repositories

import (
	"log"
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Util"
)

func GetAllProduct() ([]BusinessObjects.Product, error) {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM Products")
	if err != nil {
		log.Fatalf("Error querying the database: %v", err)
		return nil, err
	}
	defer rows.Close()

	products := []BusinessObjects.Product{}
	for rows.Next() {
		var product BusinessObjects.Product
		err := rows.Scan(&product.ProductID, &product.SellerID, &product.ProductName, &product.Description, &product.Price, &product.Quantity, &product.CategoryID, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			log.Fatalf("Error scanning row: %v", err)
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func GetProductById(id string) (BusinessObjects.Product, error) {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return BusinessObjects.Product{}, err
	}
	defer db.Close()

	var product BusinessObjects.Product

	err = db.QueryRow("SELECT * FROM Products WHERE ProductID = ?", id).Scan(&product.ProductID, &product.SellerID, &product.ProductName, &product.Description, &product.Price, &product.Quantity, &product.CategoryID, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		log.Fatalf("Error querying the database: %v", err)
		return BusinessObjects.Product{}, err
	}

	return product, nil
}

func CreateProduct(product BusinessObjects.Product) error {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO Products (ProductID, SellerID, ProductName, Description, Price, Quantity, CategoryID, CreatedAt, UpdatedAt) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)", product.ProductID, product.SellerID, product.ProductName, product.Description, product.Price, product.Quantity, product.CategoryID, product.CreatedAt, product.UpdatedAt)
	if err != nil {
		log.Fatalf("Error inserting into the database: %v", err)
		return err
	}

	return nil
}

func UpdateProduct(product BusinessObjects.Product) error {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return err
	}
	defer db.Close()

	_, err = db.Exec("UPDATE Products SET SellerID = ?, ProductName = ?, Description = ?, Price = ?, Quantity = ?, CategoryID = ?, UpdatedAt = ? WHERE ProductID = ?", product.SellerID, product.ProductName, product.Description, product.Price, product.Quantity, product.CategoryID, product.UpdatedAt, product.ProductID)
	if err != nil {
		log.Fatalf("Error updating the database: %v", err)
		return err
	}

	return nil
}

func DeleteProduct(id string) error {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM Products WHERE ProductID = ?", id)
	if err != nil {
		log.Fatalf("Error deleting from the database: %v", err)
		return err
	}

	return nil
}
