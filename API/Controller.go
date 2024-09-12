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
	router.GET("/verify", VerifyUserEmailHandler)

	router.GET("/auth/google", GoogleLogin)
	router.GET("/auth/google/callback", GoogleCallback)
	router.GET("/auth/google/logout", GoogleLogout)

	router.GET("/auth/facebook", FacebookLogin)
	router.GET("/auth/facebook/callback", FacebookCallback)
	router.GET("/auth/facebook/logout", FacebookLogout)

	return router
}

// package API

// import (
// 	"log"
// 	"th3y3m/e-commerce-platform/Middleware"

// 	"github.com/casbin/casbin/v2"
// 	"github.com/gin-gonic/gin"
// )

// func Controller() *gin.Engine {
// 	router := gin.Default()

// 	// Load Casbin enforcer
// 	enforcer, err := casbin.NewEnforcer("path/to/rbac_model.conf", "path/to/rbac_policy.csv")
// 	if err != nil {
// 		log.Fatalf("Failed to load Casbin enforcer: %v", err)
// 	}

// 	// Public routes
// 	router.POST("/login", Login)
// 	router.POST("/register", RegisterCustomer)
// 	router.GET("/verify", VerifyUserEmailHandler)

// 	// Protected routes with JWT and Casbin middleware
// 	protected := router.Group("/")
// 	protected.Use(Middleware.AuthMiddleware(enforcer))
// 	{
// 		protected.GET("/auth/google", GoogleLogin)
// 		protected.GET("/auth/google/callback", GoogleCallback)
// 		protected.GET("/auth/google/logout", GoogleLogout)
// 		protected.GET("/auth/facebook", FacebookLogin)
// 		protected.GET("/auth/facebook/callback", FacebookCallback)
// 		protected.GET("/auth/facebook/logout", FacebookLogout)
// 	}

// 	return router
// }
