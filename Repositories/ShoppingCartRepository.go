package Repositories

import (
	"log"
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Util"
)

func GetAllShoppingCart() ([]BusinessObjects.ShoppingCart, error) {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM ShoppingCarts")
	if err != nil {
		log.Fatalf("Error querying the database: %v", err)
		return nil, err
	}
	defer rows.Close()

	shoppingCarts := []BusinessObjects.ShoppingCart{}
	for rows.Next() {
		var shoppingCart BusinessObjects.ShoppingCart
		err := rows.Scan(&shoppingCart.CartID, &shoppingCart.UserID, &shoppingCart.CreatedAt)
		if err != nil {
			log.Fatalf("Error scanning row: %v", err)
			return nil, err
		}
		shoppingCarts = append(shoppingCarts, shoppingCart)
	}

	return shoppingCarts, nil
}

func GetShoppingCartById(id string) (BusinessObjects.ShoppingCart, error) {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return BusinessObjects.ShoppingCart{}, err
	}
	defer db.Close()

	var shoppingCart BusinessObjects.ShoppingCart

	err = db.QueryRow("SELECT * FROM ShoppingCarts WHERE CartID = ?", id).Scan(&shoppingCart.CartID, &shoppingCart.UserID, &shoppingCart.CreatedAt)
	if err != nil {
		log.Fatalf("Error querying the database: %v", err)
		return BusinessObjects.ShoppingCart{}, err
	}

	return shoppingCart, nil
}

func CreateShoppingCart(shoppingCart BusinessObjects.ShoppingCart) error {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return err
	}

	_, err = db.Exec("INSERT INTO ShoppingCarts (CartID, UserID, CreatedAt) VALUES (?, ?, ?)", shoppingCart.CartID, shoppingCart.UserID, shoppingCart.CreatedAt)
	if err != nil {
		log.Fatalf("Error inserting into the database: %v", err)
		return err
	}

	return nil
}

func GetShoppingCartByUserId(userId string) ([]BusinessObjects.ShoppingCart, error) {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM ShoppingCarts WHERE UserID = ?", userId)
	if err != nil {
		log.Fatalf("Error querying the database: %v", err)
		return nil, err
	}
	defer rows.Close()

	shoppingCarts := []BusinessObjects.ShoppingCart{}
	for rows.Next() {
		var shoppingCart BusinessObjects.ShoppingCart
		err := rows.Scan(&shoppingCart.CartID, &shoppingCart.UserID, &shoppingCart.CreatedAt)
		if err != nil {
			log.Fatalf("Error scanning row: %v", err)
			return nil, err
		}
		shoppingCarts = append(shoppingCarts, shoppingCart)
	}

	return shoppingCarts, nil
}

func DeleteShoppingCartById(id string) error {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM ShoppingCarts WHERE CartID = ?", id)
	if err != nil {
		log.Fatalf("Error deleting from the database: %v", err)
		return err
	}

	return nil
}

func UpdateShoppingCart(shoppingCart BusinessObjects.ShoppingCart) error {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return err
	}
	defer db.Close()

	_, err = db.Exec("UPDATE ShoppingCarts SET UserID = ?, CreatedAt = ? WHERE CartID = ?", shoppingCart.UserID, shoppingCart.CreatedAt, shoppingCart.CartID)
	if err != nil {
		log.Fatalf("Error updating the database: %v", err)
		return err
	}

	return nil
}

func GetShoppingCartByCartId(cartId string) (BusinessObjects.ShoppingCart, error) {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return BusinessObjects.ShoppingCart{}, err
	}
	defer db.Close()

	var shoppingCart BusinessObjects.ShoppingCart

	err = db.QueryRow("SELECT * FROM ShoppingCarts WHERE CartID = ?", cartId).Scan(&shoppingCart.CartID, &shoppingCart.UserID, &shoppingCart.CreatedAt)
	if err != nil {
		log.Fatalf("Error querying the database: %v", err)
		return BusinessObjects.ShoppingCart{}, err
	}

	return shoppingCart, nil
}

func GetShoppingCartByUserIdAndCartId(userId string, cartId string) (BusinessObjects.ShoppingCart, error) {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return BusinessObjects.ShoppingCart{}, err
	}
	defer db.Close()

	var shoppingCart BusinessObjects.ShoppingCart

	err = db.QueryRow("SELECT * FROM ShoppingCarts WHERE UserID = ? AND CartID = ?", userId, cartId).Scan(&shoppingCart.CartID, &shoppingCart.UserID, &shoppingCart.CreatedAt)
	if err != nil {
		log.Fatalf("Error querying the database: %v", err)
		return BusinessObjects.ShoppingCart{}, err
	}

	return shoppingCart, nil
}

func DeleteShoppingCartByUserIdAndCartId(userId string, cartId string) error {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM ShoppingCarts WHERE UserID = ? AND CartID = ?", userId, cartId)
	if err != nil {
		log.Fatalf("Error deleting from the database: %v", err)
		return err
	}

	return nil
}
