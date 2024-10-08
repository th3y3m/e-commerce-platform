package API

import (
	"net/http"
	"strconv"
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/DependencyInjection"
	"time"

	"github.com/gin-gonic/gin"
)

func GetPaginatedVoucherList(c *gin.Context) {
	sortBy := c.DefaultQuery("sortBy", "voucherID")
	voucherID := c.DefaultQuery("voucherID", "")
	pageIndex, _ := strconv.Atoi(c.DefaultQuery("pageIndex", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	status, _ := strconv.ParseBool(c.DefaultQuery("status", ""))
	startDate, _ := time.Parse(time.RFC3339, c.DefaultQuery("startDate", ""))
	endDate, _ := time.Parse(time.RFC3339, c.DefaultQuery("endDate", ""))

	module := DependencyInjection.NewVoucherServiceProvider()
	vouchers, err := module.GetPaginatedVoucherList(sortBy, voucherID, pageIndex, pageSize, &status, startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, vouchers)
}

func GetAllVouchers(c *gin.Context) {
	module := DependencyInjection.NewVoucherServiceProvider()

	vouchers, err := module.GetAllVouchers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, vouchers)
}

func GetVoucherByID(c *gin.Context) {
	id := c.Param("id")
	module := DependencyInjection.NewVoucherServiceProvider()

	voucher, err := module.GetVoucherByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, voucher)
}

func CreateVoucher(c *gin.Context) {
	var newVoucher BusinessObjects.NewVoucher
	if err := c.ShouldBindJSON(&newVoucher); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	module := DependencyInjection.NewVoucherServiceProvider()

	err := module.CreateVoucher(newVoucher)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Voucher created successfully"})
}

func UpdateVoucher(c *gin.Context) {
	var updateVoucher BusinessObjects.Voucher
	if err := c.ShouldBindJSON(&updateVoucher); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	module := DependencyInjection.NewVoucherServiceProvider()

	err := module.UpdateVoucher(updateVoucher)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Voucher updated successfully"})
}

func DeleteVoucher(c *gin.Context) {
	id := c.Param("id")
	module := DependencyInjection.NewVoucherServiceProvider()

	err := module.DeleteVoucher(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Voucher deleted successfully"})
}
