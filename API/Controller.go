package API

import (
	"github.com/gin-gonic/gin"
)

func Controller() *gin.Engine {
	router := gin.Default()

	// sessionSecret := os.Getenv("SESSION_SECRET")
	// store := cookie.NewStore([]byte(sessionSecret), []byte(sessionSecret))
	// router.Use(sessions.Sessions("mysession", store))

	// Define routes
	router.POST("/login", Login)
	router.POST("/register", RegisterCustomer)
	router.GET("/auth/google", GoogleLogin)
	router.GET("/auth/google/callback", GoogleCallback)
	router.GET("/auth/google/logout", GoogleLogout)

	return router
}
