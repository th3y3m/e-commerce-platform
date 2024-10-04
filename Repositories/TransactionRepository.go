package Repositories

import (
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Interface"
	"th3y3m/e-commerce-platform/Provider"
	"th3y3m/e-commerce-platform/Util"
	"time"

	"github.com/sirupsen/logrus"
)

type TransactionRepository struct {
	log *logrus.Logger
	db  Provider.IDb
}

func NewTransactionRepository(log *logrus.Logger, db Provider.IDb) Interface.ITransactionRepository {
	return &TransactionRepository{log: log, db: db}
}

func (t *TransactionRepository) GetPaginatedTransactionList(sortBy, transactionID, orderID string, pageIndex, pageSize int, minPrice, maxPrice *float64, status *bool, startDate, endDate time.Time) (Util.PaginatedList[BusinessObjects.Transaction], error) {
	t.log.Infof("Fetching paginated transaction list with sortBy: %s, transactionID: %s, orderID: %s, pageIndex: %d, pageSize: %d, minPrice: %v, maxPrice: %v, status: %v, startDate: %v, endDate: %v", sortBy, transactionID, orderID, pageIndex, pageSize, minPrice, maxPrice, status, startDate, endDate)
	db, err := t.db.GetDB()
	if err != nil {
		t.log.Error("Failed to connect to PostgreSQL:", err)
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
		t.log.Error("Failed to count transactions:", err)
		return Util.PaginatedList[BusinessObjects.Transaction]{}, err
	}

	if err := query.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&transactions).Error; err != nil {
		t.log.Error("Failed to fetch paginated transactions:", err)
		return Util.PaginatedList[BusinessObjects.Transaction]{}, err
	}

	t.log.Infof("Successfully fetched paginated transaction list with total count: %d", total)
	return Util.NewPaginatedList(transactions, total, pageIndex, pageSize), nil
}

// GetAllTransactions retrieves all freight transactions from the database
func (t *TransactionRepository) GetAllTransactions() ([]BusinessObjects.Transaction, error) {
	t.log.Info("Fetching all transactions")
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		t.log.Error("Failed to connect to PostgreSQL:", err)
		return nil, err
	}

	var transactions []BusinessObjects.Transaction
	if err := db.Find(&transactions).Error; err != nil {
		t.log.Error("Failed to fetch all transactions:", err)
		return nil, err
	}

	t.log.Info("Successfully fetched all transactions")
	return transactions, nil
}

// GetTransactionByID retrieves a freight transaction by its ID
func (t *TransactionRepository) GetTransactionByID(transactionID string) (BusinessObjects.Transaction, error) {
	t.log.Infof("Fetching transaction by ID: %s", transactionID)
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		t.log.Error("Failed to connect to PostgreSQL:", err)
		return BusinessObjects.Transaction{}, err
	}

	var transaction BusinessObjects.Transaction
	if err := db.First(&transaction, "transaction_id = ?", transactionID).Error; err != nil {
		t.log.Error("Failed to fetch transaction by ID:", err)
		return BusinessObjects.Transaction{}, err
	}

	t.log.Infof("Successfully fetched transaction by ID: %s", transactionID)
	return transaction, nil
}

// CreateTransaction adds a new freight transaction to the database
func (t *TransactionRepository) CreateTransaction(transaction BusinessObjects.Transaction) error {
	t.log.Infof("Creating new transaction with ID: %s", transaction.TransactionID)
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		t.log.Error("Failed to connect to PostgreSQL:", err)
		return err
	}

	if err := db.Create(&transaction).Error; err != nil {
		t.log.Error("Failed to create new transaction:", err)
		return err
	}

	t.log.Infof("Successfully created new transaction with ID: %s", transaction.TransactionID)
	return nil
}

// UpdateTransaction updates an existing freight transaction
func (t *TransactionRepository) UpdateTransaction(transaction BusinessObjects.Transaction) error {
	t.log.Infof("Updating transaction with ID: %s", transaction.TransactionID)
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		t.log.Error("Failed to connect to PostgreSQL:", err)
		return err
	}

	if err := db.Save(&transaction).Error; err != nil {
		t.log.Error("Failed to update transaction:", err)
		return err
	}

	t.log.Infof("Successfully updated transaction with ID: %s", transaction.TransactionID)
	return nil
}

// DeleteTransaction removes a freight transaction from the database by its ID
func (t *TransactionRepository) DeleteTransaction(transactionID string) error {
	t.log.Infof("Deleting transaction with ID: %s", transactionID)
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		t.log.Error("Failed to connect to PostgreSQL:", err)
		return err
	}

	// if err := db.Delete(&BusinessObjects.Transaction{}, "transaction_id = ?", transactionID).Error; err != nil {
	// 	return err
	// }

	if err := db.Model(&BusinessObjects.Transaction{}).Where("transaction_id = ?", transactionID).Update("payment_status", "Cancel").Error; err != nil {
		t.log.Error("Failed to delete transaction:", err)
		return err
	}

	t.log.Infof("Successfully deleted transaction with ID: %s", transactionID)
	return nil
}
