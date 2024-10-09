package Services

import (
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Interface"
	"th3y3m/e-commerce-platform/Util"
	"time"
)

type TransactionService struct {
	transactionRepository Interface.ITransactionRepository
}

func NewTransactionService(transactionRepository Interface.ITransactionRepository) Interface.ITransactionService {
	return &TransactionService{transactionRepository}
}

func (t *TransactionService) GetPaginatedTransactionList(sortBy string, transactionID string, orderID string, pageIndex, pageSize int, minPrice, maxPrice *float64, status *bool, startDate time.Time, endDate time.Time) (Util.PaginatedList[BusinessObjects.Transaction], error) {
	return t.transactionRepository.GetPaginatedTransactionList(sortBy, transactionID, orderID, pageIndex, pageSize, minPrice, maxPrice, status, startDate, endDate)
}

func (t *TransactionService) GetAllTransactions() ([]BusinessObjects.Transaction, error) {
	return t.transactionRepository.GetAllTransactions()
}

func (t *TransactionService) GetTransactionByID(id string) (BusinessObjects.Transaction, error) {
	return t.transactionRepository.GetTransactionByID(id)
}

func (t *TransactionService) CreateTransaction(transaction BusinessObjects.NewTransaction) error {
	newTransaction := BusinessObjects.Transaction{
		TransactionID:   "TRX" + Util.GenerateID(10),
		OrderID:         transaction.OrderID,
		PaymentAmount:   transaction.PaymentAmount,
		TransactionDate: time.Now(),
		PaymentStatus:   transaction.PaymentStatus,
	}

	return t.transactionRepository.CreateTransaction(newTransaction)
}

func (t *TransactionService) UpdateTransaction(transaction BusinessObjects.Transaction) error {
	return t.transactionRepository.UpdateTransaction(transaction)
}

func (t *TransactionService) DeleteTransaction(id string) error {
	return t.transactionRepository.DeleteTransaction(id)
}
