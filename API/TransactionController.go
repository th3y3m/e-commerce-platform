package API

import (
	"net/http"
	"strconv"
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/DependencyInjection"
	"time"

	"github.com/gin-gonic/gin"
)

func GetPaginatedTransactionList(c *gin.Context) {
	sortBy := c.DefaultQuery("sortBy", "")
	transactionID := c.DefaultQuery("transactionID", "")
	orderID := c.DefaultQuery("orderID", "")
	pageIndex, _ := strconv.Atoi(c.DefaultQuery("pageIndex", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	minPrice, _ := strconv.ParseFloat(c.DefaultQuery("minPrice", ""), 64)
	maxPrice, _ := strconv.ParseFloat(c.DefaultQuery("maxPrice", ""), 64)
	status, _ := strconv.ParseBool(c.DefaultQuery("status", ""))
	startDate, _ := time.Parse(time.RFC3339, c.DefaultQuery("startDate", ""))
	endDate, _ := time.Parse(time.RFC3339, c.DefaultQuery("endDate", ""))

	module := DependencyInjection.NewTransactionServiceProvider()
	transactions, err := module.GetPaginatedTransactionList(sortBy, transactionID, orderID, pageIndex, pageSize, &minPrice, &maxPrice, &status, startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, transactions)
}

func GetAllTransactions(c *gin.Context) {
	module := DependencyInjection.NewTransactionServiceProvider()

	transactions, err := module.GetAllTransactions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, transactions)
}

func GetTransactionByID(c *gin.Context) {
	id := c.Param("id")
	module := DependencyInjection.NewTransactionServiceProvider()

	transaction, err := module.GetTransactionByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, transaction)
}

func CreateTransaction(c *gin.Context) {
	var newTransaction BusinessObjects.NewTransaction
	if err := c.ShouldBindJSON(&newTransaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	module := DependencyInjection.NewTransactionServiceProvider()

	err := module.CreateTransaction(newTransaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Transaction created successfully"})
}

func UpdateTransaction(c *gin.Context) {
	var transaction BusinessObjects.Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	module := DependencyInjection.NewTransactionServiceProvider()

	err := module.UpdateTransaction(transaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Transaction updated successfully"})
}

func DeleteTransaction(c *gin.Context) {
	id := c.Param("id")
	module := DependencyInjection.NewTransactionServiceProvider()

	err := module.DeleteTransaction(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Transaction deleted successfully"})
}
