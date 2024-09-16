// package API

// import (
// 	"github.com/gin-gonic/gin"
// )

// func Controller() *gin.Engine {
// 	router := gin.Default()

// 	// sessionSecret := os.Getenv("SESSION_SECRET")
// 	// store := cookie.NewStore([]byte(sessionSecret), []byte(sessionSecret))
// 	// router.Use(sessions.Sessions("mysession", store))

// 	// Define routes
// 	router.POST("/login", Login)
// 	router.POST("/register", RegisterCustomer)
// 	router.GET("/verify", VerifyUserEmailHandler)

// 	router.GET("/auth/google", GoogleLogin)
// 	router.GET("/auth/google/callback", GoogleCallback)
// 	router.GET("/auth/google/logout", GoogleLogout)

// 	router.GET("/auth/facebook", FacebookLogin)
// 	router.GET("/auth/facebook/callback", FacebookCallback)
// 	router.GET("/auth/facebook/logout", FacebookLogout)

// 	return router
// }

package API

import (
	"log"
	"os"
	"path/filepath"
	"th3y3m/e-commerce-platform/Middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func Controller() *gin.Engine {
	router := gin.Default()

	// Debug: Print the current working directory
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get working directory: %v", err)
	}
	log.Printf("Current working directory: %s", wd)

	// Load Casbin enforcer
	enforcer, err := casbin.NewEnforcer(filepath.Join(wd, "rbac", "rbac_model.conf"), filepath.Join(wd, "rbac", "rbac_policy.csv"))
	if err != nil {
		log.Fatalf("Failed to load Casbin enforcer: %v", err)
	}
	log.Println("Casbin enforcer loaded successfully")

	// Public routes
	router.POST("/login", Login)
	router.POST("/register", RegisterCustomer)
	router.GET("/verify", VerifyUserEmailHandler)
	router.GET("/auth/google", GoogleLogin)
	router.GET("/auth/google/callback", GoogleCallback)
	router.GET("/auth/google/logout", GoogleLogout)
	router.GET("/auth/facebook", FacebookLogin)
	router.GET("/auth/facebook/callback", FacebookCallback)
	router.GET("/auth/facebook/logout", FacebookLogout)

	// Protected routes with JWT and Casbin middleware
	protected := router.Group("/")
	protected.Use(Middleware.AuthMiddleware(enforcer))
	{
		protected.GET("/auth/users", GetUsers)
	}

	log.Println("Routes registered successfully")

	return router
}
