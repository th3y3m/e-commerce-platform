package API

import (
	"log"
	"os"
	"path/filepath"
	"th3y3m/e-commerce-platform/Middleware"

	_ "th3y3m/e-commerce-platform/docs" // Import generated docs

	"github.com/casbin/casbin/v2"
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

	// Swagger route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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

	router.GET("/couriers", GetAllCouriers)
	router.GET("/users", GetUsers)

	// Protected routes with JWT and Casbin middleware
	protected := router.Group("/auth")
	protected.Use(Middleware.AuthMiddleware(enforcer))
	{
		protected.GET("/users/:id", GetUserByID)
		protected.PUT("/users/UpdateProfile/:id", UpdateProfile)
		protected.PUT("/users/Ban/:id", BanUser)
		protected.PUT("/users/UnBan/:id", UnBanUser)

		protected.GET("/products", GetPaginatedProductList)
		protected.GET("/products/:id", GetProductByID)
		protected.POST("/products", CreateProduct)
		protected.PUT("/products/:id", UpdateProduct)
		protected.DELETE("/products/:id", DeleteProduct)

		protected.GET("/categories", GetAllCategories)
		protected.GET("/categories/:id", GetCategoryByID)
		protected.POST("/categories", CreateCategory)
		protected.PUT("/categories/:id", UpdateCategory)
		protected.DELETE("/categories/:id", DeleteCategory)

		protected.GET("/couriers/:id", GetCourierByID)
		protected.POST("/couriers", CreateCourier)
		protected.PUT("/couriers/:id", UpdateCourier)
		protected.DELETE("/couriers/:id", DeleteCourier)

		protected.GET("/discounts", GetAllDiscounts)
		protected.GET("/discounts/:id", GetDiscountByID)
		protected.POST("/discounts", CreateDiscount)
		protected.PUT("/discounts/:id", UpdateDiscount)
		protected.DELETE("/discounts/:id", DeleteDiscount)

		protected.GET("/productDiscounts/:id", GetProductDiscountByID)

		protected.GET("/cartItems/:id", GetCartItemByCartID)

		protected.GET("/orders", GetPaginatedOrderList)
		protected.GET("/orders/:id", GetOrderById)
		protected.POST("/orders", PlaceOrder)

		protected.GET("/orderDetails/:id", GetOrderDetailOfAOrder)

		protected.GET("/freightRates", GetAllFreightRates)
		protected.GET("/freightRates/:id", GetFreightRateByID)
		protected.POST("/freightRates", CreateFreightRate)
		protected.PUT("/freightRates/:id", UpdateFreightRate)
		protected.DELETE("/freightRates/:id", DeleteFreightRate)

		protected.GET("/vouchers", GetAllVouchers)
		protected.GET("/vouchers/:id", GetVoucherByID)
		protected.POST("/vouchers", CreateVoucher)
		protected.PUT("/vouchers/:id", UpdateVoucher)
		protected.DELETE("/vouchers/:id", DeleteVoucher)

		protected.GET("/reviews", GetAllReviews)
		protected.GET("/reviews/:id", GetReviewByID)
		protected.POST("/reviews", CreateReview)
		protected.PUT("/reviews/:id", UpdateReview)
		protected.DELETE("/reviews/:id", DeleteReview)

		protected.GET("/transactions", GetPaginatedTransactionList)
		protected.GET("/transactions/:id", GetTransactionByID)
		protected.POST("/transactions", CreateTransaction)
		protected.PUT("/transactions/:id", UpdateTransaction)

		protected.GET("/shoppingCart/:id", GetShoppingCartByID)
		protected.GET("/shoppingCart", GetUserShoppingCart)
		protected.POST("/shoppingCart", AddProductToCart)
		protected.PUT("/shoppingCart", RemoveProductFromCart)
		protected.DELETE("/shoppingCart/:id", ClearShoppingCart)
		protected.GET("/shoppingCart/numberofitems/:id", NumberOfItemsInCart)

		protected.PUT("/cookie/deleteUnitItem", DeleteUnitItem)
		protected.PUT("/cookie/removeFromCart", RemoveFromCart)
		protected.GET("/cookie/getCartItems/:id", GetCartItems)
		protected.PUT("/cookie/deleteCartInCookie/:id", DeleteCartInCookie)
		protected.GET("/cookie/numberOfItemsInCartCookie/:id", NumberOfItemsInCartCookie)
		protected.POST("/cookie/saveCartToCookieHandler", SaveCartToCookieHandler)

		protected.GET("/news", GetAllNews)
		protected.GET("/news/:id", GetNewsByID)
		protected.POST("/news", CreateNews)
		protected.PUT("/news/:id", UpdateNews)
		protected.DELETE("/news/:id", DeleteNews)

		protected.GET("/vnpay", CreateVNPayUrl)
		protected.POST("/vnpay", ValidateVNPayResponse)

		protected.GET("/momo", CreateMoMoUrl)
		protected.POST("/momo", ValidateMoMoResponse)
	}

	log.Println("Routes registered successfully")

	return router
}
