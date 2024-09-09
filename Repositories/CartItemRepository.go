package Repositories

import (
	"log"
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Util"
)

func GetAllCartItems() ([]BusinessObjects.CartItem, error) {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM CartItems")
	if err != nil {
		log.Fatalf("Error querying the database: %v", err)
		return nil, err
	}
	defer rows.Close()

	cartItems := []BusinessObjects.CartItem{}

	for rows.Next() {
		var cartItem BusinessObjects.CartItem
		err := rows.Scan(&cartItem.CartItemID, &cartItem.CartID, &cartItem.ProductID, &cartItem.Quantity)
		if err != nil {
			log.Fatalf("Error scanning row: %v", err)
			return nil, err
		}
		cartItems = append(cartItems, cartItem)
	}

	return cartItems, nil
}

func GetCartItemByID(cartItemID string) (BusinessObjects.CartItem, error) {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return BusinessObjects.CartItem{}, err
	}
	defer db.Close()

	var cartItem BusinessObjects.CartItem

	err = db.QueryRow("SELECT * FROM CartItems WHERE CartItemID = ?", cartItemID).Scan(&cartItem.CartItemID, &cartItem.CartID, &cartItem.ProductID, &cartItem.Quantity)
	if err != nil {
		log.Fatalf("Error querying the database: %v", err)
		return BusinessObjects.CartItem{}, err
	}

	return cartItem, nil
}

func CreateCartItem(cartItem BusinessObjects.CartItem) error {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return err
	}

	_, err = db.Exec("INSERT INTO CartItems (CartItemID, CartID, ProductID, Quantity) VALUES (?, ?, ?, ?)", cartItem.CartItemID, cartItem.CartID, cartItem.ProductID, cartItem.Quantity)
	if err != nil {
		log.Fatalf("Error inserting into the database: %v", err)
		return err
	}

	return nil
}

func UpdateCartItem(cartItem BusinessObjects.CartItem) error {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return err
	}

	_, err = db.Exec("UPDATE CartItems SET CartID = ?, ProductID = ?, Quantity = ? WHERE CartItemID = ?", cartItem.CartID, cartItem.ProductID, cartItem.Quantity, cartItem.CartItemID)
	if err != nil {
		log.Fatalf("Error updating the database: %v", err)
	}

	return nil
}

func DeleteCartItem(cartItem BusinessObjects.CartItem) error {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
	}

	_, err = db.Exec("DELETE FROM CartItems WHERE CartItemID = ?", cartItem.CartItemID)
	if err != nil {
		log.Fatalf("Error deleting from the database: %v", err)
	}

	return nil
}
