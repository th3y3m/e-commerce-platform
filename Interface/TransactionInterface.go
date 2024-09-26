package Interface

import (
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Util"
	"time"
)

type ITransactionRepository interface {
	GetPaginatedTransactionList(sortBy, transactionID, orderID string, pageIndex, pageSize int, minPrice, maxPrice *float64, status *bool, startDate, endDate time.Time) (Util.PaginatedList[BusinessObjects.Transaction], error)
	GetAllTransactions() ([]BusinessObjects.Transaction, error)
	GetTransactionByID(transactionID string) (BusinessObjects.Transaction, error)
	CreateTransaction(transaction BusinessObjects.Transaction) error
	UpdateTransaction(transaction BusinessObjects.Transaction) error
	DeleteTransaction(transactionID string) error
}
type ITransactionService interface {
	GetPaginatedTransactionList(sortBy, transactionID, orderID string, pageIndex, pageSize int, minPrice, maxPrice *float64, status *bool, startDate, endDate time.Time) (Util.PaginatedList[BusinessObjects.Transaction], error)
	GetAllTransactions() ([]BusinessObjects.Transaction, error)
	GetTransactionByID(transactionID string) (BusinessObjects.Transaction, error)
	CreateTransaction(transaction BusinessObjects.NewTransaction) error
	UpdateTransaction(transaction BusinessObjects.Transaction) error
	DeleteTransaction(transactionID string) error
}
