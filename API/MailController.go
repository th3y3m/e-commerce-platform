package API

import (
	"net/http"
	"th3y3m/e-commerce-platform/DependencyInjection"

	"github.com/gin-gonic/gin"
)

// VerifyUserEmailHandler verifies the user's email and cancels the deletion task
func VerifyUserEmailHandler(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token is missing"})
		return
	}
	service := DependencyInjection.NewNewAuthenticationServiceProvider()

	err := service.VerifyUserEmail(token)
	if err != nil {
		if err.Error() == "token expired or invalid" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Email verified successfully!"})
}
