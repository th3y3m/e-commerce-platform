package Services

import (
	"testing"
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Util"
	"th3y3m/e-commerce-platform/mocks"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetPaginatedTransactionList_Success(t *testing.T) {
	transactionRepository := &mocks.ITransactionRepository{}

	var minPrice, maxPrice *float64
	var status *bool

	transactionService := NewTransactionService(transactionRepository)
	transactionRepository.On("GetPaginatedTransactionList", "sortBy", "transactionID", "orderID", 1, 10, minPrice, maxPrice, status, time.Time{}, time.Time{}).Return(Util.PaginatedList[BusinessObjects.Transaction]{}, nil)
	_, err := transactionService.GetPaginatedTransactionList("sortBy", "transactionID", "orderID", 1, 10, minPrice, maxPrice, status, time.Time{}, time.Time{})
	assert.NoError(t, err)
	transactionRepository.AssertExpectations(t)
}

func TestGetAllTransactions_Success(t *testing.T) {
	transactionRepository := &mocks.ITransactionRepository{}

	transactionService := NewTransactionService(transactionRepository)
	transactionRepository.On("GetAllTransactions").Return([]BusinessObjects.Transaction{}, nil)
	_, err := transactionService.GetAllTransactions()
	assert.NoError(t, err)
	transactionRepository.AssertExpectations(t)
}

func TestGetTransactionByID_Success(t *testing.T) {
	transactionRepository := &mocks.ITransactionRepository{}

	transactionService := NewTransactionService(transactionRepository)
	transactionRepository.On("GetTransactionByID", "id").Return(BusinessObjects.Transaction{}, nil)
	_, err := transactionService.GetTransactionByID("id")
	assert.NoError(t, err)
	transactionRepository.AssertExpectations(t)
}

func TestCreateTransaction_Success(t *testing.T) {
	transactionRepository := &mocks.ITransactionRepository{}

	transactionService := NewTransactionService(transactionRepository)
	transactionRepository.On("CreateTransaction", mock.AnythingOfType("BusinessObjects.Transaction")).Return(nil)
	err := transactionService.CreateTransaction(BusinessObjects.NewTransaction{})
	assert.NoError(t, err)
	transactionRepository.AssertExpectations(t)
}

func TestUpdateTransaction_Success(t *testing.T) {
	transactionRepository := &mocks.ITransactionRepository{}

	transactionService := NewTransactionService(transactionRepository)
	transactionRepository.On("UpdateTransaction", mock.AnythingOfType("BusinessObjects.Transaction")).Return(nil)
	err := transactionService.UpdateTransaction(BusinessObjects.Transaction{})
	assert.NoError(t, err)
	transactionRepository.AssertExpectations(t)
}

func TestDeleteTransaction_Success(t *testing.T) {
	transactionRepository := &mocks.ITransactionRepository{}

	transactionService := NewTransactionService(transactionRepository)
	transactionRepository.On("DeleteTransaction", "id").Return(nil)
	err := transactionService.DeleteTransaction("id")
	assert.NoError(t, err)
	transactionRepository.AssertExpectations(t)
}
