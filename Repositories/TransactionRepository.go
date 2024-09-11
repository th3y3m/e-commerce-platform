package Repositories

import (
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Util"
	"time"
)

func GetPaginatedTransactionList(sortBy, transactionID, orderID string, pageIndex, pageSize int, minPrice, maxPrice *int, status *bool, startDate, endDate time.Time) (Util.PaginatedList[BusinessObjects.Transaction], error) {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return Util.PaginatedList[BusinessObjects.Transaction]{}, err
	}

	var transactions []BusinessObjects.Transaction
	query := db.Model(&BusinessObjects.Transaction{})

	if transactionID != "" {
		query = query.Where("transaction_id = ?", transactionID)
	}

	if orderID != "" {
		query = query.Where("order_id = ?", orderID)
	}

	if minPrice != nil {
		query = query.Where("price >= ?", *minPrice)
	}

	if maxPrice != nil {
		query = query.Where("price <= ?", *maxPrice)
	}

	if !startDate.IsZero() {
		query = query.Where("transaction_date >= ?", startDate)
	}

	if !endDate.IsZero() {
		query = query.Where("transaction_date <= ?", endDate)
	}

	if status != nil {
		query = query.Where("payment_status = ?", *status)
	}

	switch sortBy {
	case "transaction_id_asc":
		query = query.Order("transaction_id ASC")
	case "transaction_id_desc":
		query = query.Order("transaction_id DESC")
	case "order_id_asc":
		query = query.Order("order_id ASC")
	case "order_id_desc":
		query = query.Order("order_id DESC")
	case "payment_amount_asc":
		query = query.Order("payment_amount ASC")
	case "payment_amount_desc":
		query = query.Order("payment_amount DESC")
	case "transaction_date_asc":
		query = query.Order("transaction_date ASC")
	case "transaction_date_desc":
		query = query.Order("transaction_date DESC")
	case "payment_status_asc":
		query = query.Order("payment_status ASC")
	case "payment_status_desc":
		query = query.Order("payment_status DESC")
	case "payment_method_asc":
		query = query.Order("payment_method ASC")
	case "payment_method_desc":
		query = query.Order("payment_method DESC")
	default:
		query = query.Order("transaction_date DESC")
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return Util.PaginatedList[BusinessObjects.Transaction]{}, err
	}

	if err := query.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&transactions).Error; err != nil {
		return Util.PaginatedList[BusinessObjects.Transaction]{}, err
	}

	return Util.PaginatedList[BusinessObjects.Transaction]{
		Items: transactions, TotalCount: total, PageIndex: pageIndex, PageSize: pageSize}, nil
}

// GetAllTransactions retrieves all freight transactions from the database
func GetAllTransactions() ([]BusinessObjects.Transaction, error) {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return nil, err
	}

	var transactions []BusinessObjects.Transaction
	if err := db.Find(&transactions).Error; err != nil {
		return nil, err
	}

	return transactions, nil
}

// GetTransactionByID retrieves a freight transaction by its ID
func GetTransactionByID(transactionID string) (BusinessObjects.Transaction, error) {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return BusinessObjects.Transaction{}, err
	}

	var transaction BusinessObjects.Transaction
	if err := db.First(&transaction, "transaction_id = ?", transactionID).Error; err != nil {
		return BusinessObjects.Transaction{}, err
	}

	return transaction, nil
}

// CreateTransaction adds a new freight transaction to the database
func CreateTransaction(transaction BusinessObjects.Transaction) error {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return err
	}

	if err := db.Create(&transaction).Error; err != nil {
		return err
	}

	return nil
}

// UpdateTransaction updates an existing freight transaction
func UpdateTransaction(transaction BusinessObjects.Transaction) error {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return err
	}

	if err := db.Save(&transaction).Error; err != nil {
		return err
	}

	return nil
}

// DeleteTransaction removes a freight transaction from the database by its ID
func DeleteTransaction(transactionID string) error {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return err
	}

	if err := db.Delete(&BusinessObjects.Transaction{}, "transaction_id = ?", transactionID).Error; err != nil {
		return err
	}

	return nil
}
