package Repositories

import (
	"log"
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Util"
)

func GetAllTransaction() ([]BusinessObjects.Transaction, error) {

	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM Transactions")
	if err != nil {
		log.Fatalf("Error querying the database: %v", err)
		return nil, err
	}
	defer rows.Close()

	transactions := []BusinessObjects.Transaction{}
	for rows.Next() {
		var transaction BusinessObjects.Transaction
		err := rows.Scan(&transaction.TransactionID, &transaction.OrderID, &transaction.PaymentAmount, &transaction.TransactionDate, &transaction.PaymentMethod, &transaction.PaymentStatus)
		if err != nil {
			log.Fatalf("Error scanning row: %v", err)
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func GetTransactionById(id string) (BusinessObjects.Transaction, error) {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return BusinessObjects.Transaction{}, err
	}
	defer db.Close()

	var transaction BusinessObjects.Transaction

	err = db.QueryRow("SELECT * FROM Transactions WHERE TransactionID = ?", id).Scan(&transaction.TransactionID, &transaction.OrderID, &transaction.PaymentAmount, &transaction.TransactionDate, &transaction.PaymentMethod, &transaction.PaymentStatus)
	if err != nil {
		log.Fatalf("Error querying the database: %v", err)
		return BusinessObjects.Transaction{}, err
	}

	return transaction, nil
}

func CreateTransaction(transaction BusinessObjects.Transaction) error {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO Transactions (OrderID, PaymentAmount, TransactionDate, PaymentMethod, PaymentStatus) VALUES (?, ?, ?, ?, ?)", transaction.OrderID, transaction.PaymentAmount, transaction.TransactionDate, transaction.PaymentMethod, transaction.PaymentStatus)
	if err != nil {
		log.Fatalf("Error inserting into the database: %v", err)
		return err
	}

	return nil
}

func UpdateTransaction(transaction BusinessObjects.Transaction) error {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return err
	}

	_, err = db.Exec("UPDATE Transactions SET OrderID = ?, PaymentAmount = ?, TransactionDate = ?, PaymentMethod = ?, PaymentStatus = ? WHERE TransactionID = ?", transaction.OrderID, transaction.PaymentAmount, transaction.TransactionDate, transaction.PaymentMethod, transaction.PaymentStatus, transaction.TransactionID)
	if err != nil {
		log.Fatalf("Error updating the database: %v", err)
		return err
	}

	return nil
}

func DeleteTransaction(id string) error {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM Transactions WHERE TransactionID = ?", id)
	if err != nil {
		log.Fatalf("Error deleting from the database: %v", err)
		return err
	}

	return nil
}
