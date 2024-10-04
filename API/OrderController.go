package API

import (
	"net/http"
	"strconv"
	"th3y3m/e-commerce-platform/DependencyInjection"
	"time"

	"github.com/gin-gonic/gin"
)

func GetPaginatedOrderList(c *gin.Context) {
	module := DependencyInjection.NewOrderServiceProvider()
	sortBy := c.DefaultQuery("sortBy", "")
	orderID := c.DefaultQuery("orderID", "")
	customerID := c.DefaultQuery("customerID", "")
	courierId := c.DefaultQuery("courierId", "")
	voucherId := c.DefaultQuery("voucherId", "")
	pageIndex, _ := strconv.Atoi(c.DefaultQuery("pageIndex", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	minPrice, _ := strconv.ParseFloat(c.DefaultQuery("minPrice", ""), 64)
	maxPrice, _ := strconv.ParseFloat(c.DefaultQuery("maxPrice", ""), 64)
	status := c.DefaultQuery("status", "")

	// Get date values, check if they are empty or invalid
	var startDate, endDate *time.Time

	startDateStr := c.DefaultQuery("startDate", "")
	if startDateStr != "" {
		parsedStartDate, err := time.Parse(time.RFC3339, startDateStr)
		if err == nil {
			startDate = &parsedStartDate // set startDate pointer if valid
		}
	}

	endDateStr := c.DefaultQuery("endDate", "")
	if endDateStr != "" {
		parsedEndDate, err := time.Parse(time.RFC3339, endDateStr)
		if err == nil {
			endDate = &parsedEndDate // set endDate pointer if valid
		}
	}

	// Call the service with nil pointers if no valid dates are provided
	orders, err := module.GetPaginatedOrderList(sortBy, orderID, customerID, courierId, voucherId, pageIndex, pageSize, startDate, endDate, &minPrice, &maxPrice, status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, orders)
}

func GetAllOrders(c *gin.Context) {
	module := DependencyInjection.NewOrderServiceProvider()
	orders, err := module.GetAllOrders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, orders)
}

func GetOrderById(c *gin.Context) {
	module := DependencyInjection.NewOrderServiceProvider()
	id := c.Param("id")
	order, err := module.GetOrderById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, order)
}

func PlaceOrder(c *gin.Context) {
	module := DependencyInjection.NewOrderServiceProvider()
	userId := c.DefaultQuery("userId", "")
	cartId := c.DefaultQuery("cartId", "")
	shipAddress := c.DefaultQuery("shipAddress", "")
	courierId := c.DefaultQuery("courierId", "")
	voucherId := c.DefaultQuery("voucherId", "")
	paymentMethod := c.DefaultQuery("paymentMethod", "VnPay")

	url, err := module.PlaceOrder(userId, cartId, shipAddress, courierId, voucherId, paymentMethod)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, url)
}
