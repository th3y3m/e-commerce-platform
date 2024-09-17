package API

import (
	"net/http"
	"strconv"
	"th3y3m/e-commerce-platform/Services"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	searchValue := c.DefaultQuery("searchValue", "")
	sortBy := c.DefaultQuery("sortBy", "")
	pageIndex, _ := strconv.Atoi(c.DefaultQuery("pageIndex", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	statusParam := c.DefaultQuery("status", "")
	var status *bool

	if statusParam != "" {
		statusValue := statusParam == "true"
		status = &statusValue
	}

	users, err := Services.GetPaginatedUserList(searchValue, sortBy, pageIndex, pageSize, status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	user, err := Services.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func UpdateProfile(c *gin.Context) {
	ID := c.Param("id")

	var info struct {
		FullName    string `json:"full_name" binding:"required"`
		PhoneNumber string `json:"phone_number" binding:"required"`
		Address     string `json:"address" binding:"required"`
	}
	if err := c.ShouldBindJSON(&info); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := Services.UpdateProfile(ID, info.FullName, info.PhoneNumber, info.Address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func BanUser(c *gin.Context) {
	id := c.Param("id")
	err := Services.BanUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User banned successfully"})
}

func UnBanUser(c *gin.Context) {
	id := c.Param("id")
	err := Services.UnBanUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User unbanned successfully"})
}
