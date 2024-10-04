package Repositories_test

// import (
// 	"testing"
// 	"th3y3m/e-commerce-platform/BusinessObjects"
// 	"th3y3m/e-commerce-platform/Provider"
// 	"th3y3m/e-commerce-platform/Repositories"

// 	"github.com/sirupsen/logrus"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// // Mock function for GetTransactionByID
// func TestGetTransactionByID(t *testing.T) {
// 	// Initialize the logger
// 	logger := logrus.New()

// 	// Initialize the DbProvider
// 	dbProvider := Provider.NewDbProvider(logger)

// 	// Get the mock database
// 	mockDb, err := dbProvider.GetMockDb()
// 	if err != nil {
// 		t.Fatalf("Failed to get mock database: %v", err)
// 	}

// 	// Set up expectations
// 	transactionID := "12345"
// 	expectedTransaction := BusinessObjects.Transaction{
// 		TransactionID: transactionID,
// 		// Add other fields as necessary
// 	}
// 	mockDb.On("First", mock.Anything, "transaction_id = ?", transactionID).Return(func(transaction *BusinessObjects.Transaction, query string, args ...interface{}) error {
// 		*transaction = expectedTransaction
// 		return nil
// 	})

// 	// Initialize the TransactionRepository
// 	repo := Repositories.NewTransactionRepository(logger, dbProvider)

// 	// Call the function under test
// 	result, err := repo.GetTransactionByID(transactionID)

// 	// Assert the results
// 	assert.NoError(t, err)
// 	assert.Equal(t, expectedTransaction, result)

// 	// Assert that the expectations were met
// 	mockDb.AssertExpectations(t)
// }
