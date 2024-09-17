package Services

import (
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Repositories"
	"th3y3m/e-commerce-platform/Util"
	"time"
)

func GetPaginatedTransactionList(sortBy string, transactionID string, orderID string, pageIndex, pageSize int, minPrice, maxPrice *float64, status *bool, startDate time.Time, endDate time.Time) (Util.PaginatedList[BusinessObjects.Transaction], error) {
	return Repositories.GetPaginatedTransactionList(sortBy, transactionID, orderID, pageIndex, pageSize, minPrice, maxPrice, status, startDate, endDate)
}

func GetAllTransactions() ([]BusinessObjects.Transaction, error) {
	return Repositories.GetAllTransactions()
}

func GetTransactionByID(id string) (BusinessObjects.Transaction, error) {
	return Repositories.GetTransactionByID(id)
}

func CreateTransaction(transaction BusinessObjects.NewTransaction) error {
	newTransaction := BusinessObjects.Transaction{
		TransactionID:   "TRX" + Util.GenerateID(10),
		OrderID:         transaction.OrderID,
		PaymentAmount:   transaction.PaymentAmount,
		TransactionDate: time.Now(),
		PaymentStatus:   transaction.PaymentStatus,
	}

	err := Repositories.CreateTransaction(newTransaction)
	if err != nil {
		return err
	}

	return nil
}

func UpdateTransaction(transaction BusinessObjects.Transaction) error {
	return Repositories.UpdateTransaction(transaction)
}

func DeleteTransaction(id string) error {
	return Repositories.DeleteTransaction(id)
}
