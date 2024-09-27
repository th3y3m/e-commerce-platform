package API

import (
	"log"
	"os"
	"path/filepath"
	"th3y3m/e-commerce-platform/Middleware"

	_ "th3y3m/e-commerce-platform/docs" // Import generated docs

	"github.com/casbin/casbin/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title E-Commerce Platform API
// @version 1.0
// @description This is an e-commerce platform API.
// @host localhost:8080
// @BasePath /
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

	// CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Swagger route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Public routes
	router.POST("/login", Login)
	router.POST("/register", RegisterCustomer)
	router.GET("/verify", VerifyUserEmailHandler)
	router.GET("/google", GoogleLogin)
	router.GET("/google/callback", GoogleCallback)
	router.GET("/google/logout", GoogleLogout)
	router.GET("/facebook", FacebookLogin)
	router.GET("/facebook/callback", FacebookCallback)
	router.GET("/facebook/logout", FacebookLogout)

	router.GET("/couriers", GetAllCouriers)
	router.GET("/users", GetUsers)

	router.GET("/categories", GetAllCategories)
	router.GET("/categories/:id", GetCategoryByID)

	router.GET("/products", GetPaginatedProductList)
	router.GET("/products/:id", GetProductByID)
	router.GET("/products/:id/price", GetProductPriceAfterDiscount)

	// Protected routes with JWT and Casbin middleware
	protected := router.Group("/auth")
	protected.Use(Middleware.AuthMiddleware(enforcer))
	{
		// routes
	}
	router.GET("/users/:id", GetUserByID)
	router.PUT("/users/UpdateProfile/:id", UpdateProfile)
	router.PUT("/users/Ban/:id", BanUser)
	router.PUT("/users/UnBan/:id", UnBanUser)
	router.POST("/products", CreateProduct)
	router.PUT("/products/:id", UpdateProduct)
	router.DELETE("/products/:id", DeleteProduct)

	router.POST("/categories", CreateCategory)
	router.PUT("/categories/:id", UpdateCategory)
	router.DELETE("/categories/:id", DeleteCategory)

	router.GET("/couriers/:id", GetCourierByID)
	router.POST("/couriers", CreateCourier)
	router.PUT("/couriers/:id", UpdateCourier)
	router.DELETE("/couriers/:id", DeleteCourier)

	router.GET("/discounts", GetAllDiscounts)
	router.GET("/discounts/:id", GetDiscountByID)
	router.POST("/discounts", CreateDiscount)
	router.PUT("/discounts/:id", UpdateDiscount)
	router.DELETE("/discounts/:id", DeleteDiscount)

	router.GET("/productDiscounts/:id", GetProductDiscountByID)

	router.GET("/cartItems/:id", GetCartItemByCartID)

	router.GET("/orders", GetPaginatedOrderList)
	router.GET("/orders/:id", GetOrderById)
	router.POST("/orders", PlaceOrder)

	router.GET("/orderDetails/:id", GetOrderDetailOfAOrder)

	router.GET("/freightRates", GetAllFreightRates)
	router.GET("/freightRates/:id", GetFreightRateByID)
	router.POST("/freightRates", CreateFreightRate)
	router.PUT("/freightRates/:id", UpdateFreightRate)
	router.DELETE("/freightRates/:id", DeleteFreightRate)

	router.GET("/vouchers", GetAllVouchers)
	router.GET("/vouchers/:id", GetVoucherByID)
	router.POST("/vouchers", CreateVoucher)
	router.PUT("/vouchers/:id", UpdateVoucher)
	router.DELETE("/vouchers/:id", DeleteVoucher)

	router.GET("/reviews", GetAllReviews)
	router.GET("/reviews/:id", GetReviewByID)
	router.POST("/reviews", CreateReview)
	router.PUT("/reviews/:id", UpdateReview)
	router.DELETE("/reviews/:id", DeleteReview)

	router.GET("/transactions", GetPaginatedTransactionList)
	router.GET("/transactions/:id", GetTransactionByID)
	router.POST("/transactions", CreateTransaction)
	router.PUT("/transactions/:id", UpdateTransaction)

	router.GET("/shoppingCart/:id", GetShoppingCartByID)
	router.GET("/shoppingCart", GetUserShoppingCart)
	router.POST("/shoppingCart", AddProductToCart)
	router.PUT("/shoppingCart", RemoveProductFromCart)
	router.DELETE("/shoppingCart/:id", ClearShoppingCart)
	router.GET("/shoppingCart/numberofitems/:id", NumberOfItemsInCart)

	router.PUT("/cookie/deleteUnitItem", DeleteUnitItem)
	router.PUT("/cookie/removeFromCart", RemoveFromCart)
	router.GET("/cookie/getCartItems/:id", GetCartItems)
	router.PUT("/cookie/deleteCartInCookie/:id", DeleteCartInCookie)
	router.GET("/cookie/numberOfItemsInCartCookie/:id", NumberOfItemsInCartCookie)
	router.POST("/cookie/saveCartToCookieHandler", SaveCartToCookieHandler)

	router.GET("/news", GetAllNews)
	router.GET("/news/:id", GetNewsByID)
	router.POST("/news", CreateNews)
	router.PUT("/news/:id", UpdateNews)
	router.DELETE("/news/:id", DeleteNews)

	router.GET("/vnpay", CreateVNPayUrl)
	router.POST("/vnpay", ValidateVNPayResponse)

	router.GET("/momo", CreateMoMoUrl)
	router.POST("/momo", ValidateMoMoResponse)

	log.Println("Routes registered successfully")

	return router
}
