package API

import (
	"net/http"

	"th3y3m/e-commerce-platform/Services"

	"github.com/gin-gonic/gin"
)

func VerifyUserEmailHandler(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token is missing"})
		return
	}

	if !Services.VerifyToken(token) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token expired or invalid"})
		return
	}

	// user, err := Repositories.GetUserByToken(token)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user"})
	// 	return
	// }

	// Cancel the scheduled delete user task
	// inspector := asynq.NewInspector(asynq.RedisClientOpt{Addr: "localhost:6379"})
	// defer inspector.Close()

	// err = inspector.DeleteTask("default", user.UserID)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to cancel task"})
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{"message": "Email verified successfully!"})
}
