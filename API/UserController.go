package API

import (
	"net/http"
	"th3y3m/e-commerce-platform/Services"

	"github.com/gin-gonic/gin"
)

// UserController is the controller for user management
func UserController() {
	router := gin.Default()

	router.Run("localhost:8081") // Separate port for user management
}

func GetUsers(c *gin.Context) {
	// Get all users
	users, err := Services.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}
