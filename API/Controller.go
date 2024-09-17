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
		protected.GET("/auth/users/:id", GetUserByID)
		protected.PUT("/auth/users/UpdateProfile/:id", UpdateProfile)
		protected.PUT("/auth/users/Ban/:id", BanUser)
		protected.PUT("/auth/users/UnBan/:id", UnBanUser)

		protected.GET("/auth/products", GetPaginatedProductList)
		protected.GET("/auth/products/:id", GetProductByID)
		protected.POST("/auth/products", CreateProduct)
		protected.PUT("/auth/products/:id", UpdateProduct)
		protected.DELETE("/auth/products/:id", DeleteProduct)

		protected.GET("/auth/categories", GetAllCategories)
		protected.GET("/auth/categories/:id", GetCategoryByID)
		protected.POST("/auth/categories", CreateCategory)
		protected.PUT("/auth/categories/:id", UpdateCategory)
		protected.DELETE("/auth/categories/:id", DeleteCategory)

		protected.GET("/auth/couriers", GetAllCouriers)
		protected.GET("/auth/couriers/:id", GetCourierByID)
		protected.POST("/auth/couriers", CreateCourier)
		protected.PUT("/auth/couriers/:id", UpdateCourier)
		protected.DELETE("/auth/couriers/:id", DeleteCourier)

		protected.GET("/auth/discounts", GetAllDiscounts)
		protected.GET("/auth/discounts/:id", GetDiscountByID)
		protected.POST("/auth/discounts", CreateDiscount)
		protected.PUT("/auth/discounts/:id", UpdateDiscount)
		protected.DELETE("/auth/discounts/:id", DeleteDiscount)

		protected.GET("/auth/productDiscounts/:id", GetProductDiscountByID)

		protected.GET("/auth/orders", GetPaginatedOrderList)
		protected.GET("/auth/orders/:id", GetOrderById)
		protected.POST("/auth/orders", PlaceOrder)

		protected.GET("/auth/vouchers", GetAllVouchers)
		protected.GET("/auth/vouchers/:id", GetVoucherByID)
		protected.POST("/auth/vouchers", CreateVoucher)
		protected.PUT("/auth/vouchers/:id", UpdateVoucher)
		protected.DELETE("/auth/vouchers/:id", DeleteVoucher)

		protected.GET("/auth/reviews", GetAllReviews)
		protected.GET("/auth/reviews/:id", GetReviewByID)
		protected.POST("/auth/reviews", CreateReview)
		protected.PUT("/auth/reviews/:id", UpdateReview)
		protected.DELETE("/auth/reviews/:id", DeleteReview)

		protected.GET("/auth/transactions", GetPaginatedTransactionList)
		protected.GET("/auth/transactions/:id", GetTransactionByID)
		protected.POST("/auth/transactions", CreateTransaction)
		protected.PUT("/auth/transactions/:id", UpdateTransaction)

		protected.GET("/auth/shoppingcart/:id", GetShoppingCartByID)
		protected.GET("/auth/shoppingcart", GetUserShoppingCart)
		protected.POST("/auth/shoppingcart", AddProductToCart)
		protected.PUT("/auth/shoppingcart", RemoveProductFromCart)
		protected.DELETE("/auth/shoppingcart/:id", ClearShoppingCart)
		protected.GET("/auth/shoppingcart/numberofitems", NumberOfItemsInCart)

		protected.GET("/auth/cookie/:id", DeleteUnitItem)
		protected.GET("/auth/cookie", RemoveFromCart)
		protected.POST("/auth/cookie/GetCartItems", GetCartItems)
		protected.PUT("/auth/cookie/DeleteCartInCookie", DeleteCartInCookie)
		protected.DELETE("/auth/cookie/NumberOfItemsInCartCookie", NumberOfItemsInCartCookie)
		protected.GET("/auth/cookie/NumberOfItemsInCartCookie", NumberOfItemsInCartCookie)

		protected.GET("/auth/news", GetAllNews)
		protected.GET("/auth/news/:id", GetNewsByID)
		protected.POST("/auth/news", CreateNews)
		protected.PUT("/auth/news/:id", UpdateNews)
		protected.DELETE("/auth/news/:id", DeleteNews)

		protected.GET("/auth/vnpay", CreateVNPayUrl)
		protected.POST("/auth/vnpay", ValidateVNPayResponse)

		protected.GET("/auth/momo", CreateMoMoUrl)
		protected.POST("/auth/momo", ValidateMoMoResponse)
	}

	log.Println("Routes registered successfully")

	return router
}
