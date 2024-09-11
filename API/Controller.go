package API

import (
	"github.com/gin-gonic/gin"
)

func Controller() *gin.Engine {
	router := gin.Default()

	router.GET("/users", GetUsers)

	return router
}
