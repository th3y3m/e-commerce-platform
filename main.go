package main

import (
	"log"
	"th3y3m/e-commerce-platform/API"
)

func main() {
	// Load environment variables from .env file
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatalf("Error loading .env file")
	// }

	// redisAddr := os.Getenv("REDIS_ADDR")
	// redisPassword := os.Getenv("REDIS_PASSWORD")
	// redisDBStr := os.Getenv("REDIS_DB")
	// redisDB, err := strconv.Atoi(redisDBStr)
	// if err != nil {
	// 	log.Fatalf("Invalid REDIS_DB value: %v", err)
	// }

	// srv := asynq.NewServer(
	// 	asynq.RedisClientOpt{
	// 		Addr:         redisAddr,
	// 		Password:     redisPassword,
	// 		DB:           redisDB,
	// 		DialTimeout:  30 * time.Second, // Increase dial timeout
	// 		ReadTimeout:  30 * time.Second, // Increase read timeout
	// 		WriteTimeout: 30 * time.Second, // Increase write timeout
	// 	},
	// 	asynq.Config{
	// 		Concurrency: 10,
	// 	},
	// )

	// mux := asynq.NewServeMux()
	// mux.HandleFunc(Services.TypeDeleteUser, Services.HandleDeleteUserTask)

	// // Run the asynq server in a separate goroutine
	// go func() {
	// 	if err := srv.Run(mux); err != nil {
	// 		log.Fatalf("could not run asynq server: %v", err)
	// 	}
	// }()

	// Start the Gin server
	router := API.Controller()
	if err := router.Run("localhost:8080"); err != nil {
		log.Fatalf("could not run Gin server: %v", err)
	}
}

// import (
// 	"log"

// 	"th3y3m/e-commerce-platform/BusinessObjects"

// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// var db *gorm.DB

// func main() {
// 	dsn := "host=localhost user=postgres password=12345 dbname=SendoDb port=5432 sslmode=disable TimeZone=Asia/Ho_Chi_Minh"

// 	var err error
// 	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatalf("Failed to connect to the database: %v", err)
// 	}

// 	// Auto-migrate models in the correct order to handle foreign key dependencies
// 	err = db.AutoMigrate(
// 		&BusinessObjects.User{},            // Users table
// 		&BusinessObjects.Courier{},         // Couriers table
// 		&BusinessObjects.Voucher{},         // Vouchers table
// 		&BusinessObjects.Category{},        // Categories
// 		&BusinessObjects.Product{},         // Products (has foreign key to Category)
// 		&BusinessObjects.Discount{},        // Discounts
// 		&BusinessObjects.News{},            // News (foreign key to User)
// 		&BusinessObjects.Order{},           // Orders (foreign key to User, Courier, Voucher)
// 		&BusinessObjects.OrderDetail{},     // OrderDetails (foreign key to Order, Product)
// 		&BusinessObjects.CartItem{},        // CartItems (foreign key to Product)
// 		&BusinessObjects.ShoppingCart{},    // ShoppingCarts (foreign key to User)
// 		&BusinessObjects.FreightRate{},     // FreightRates (foreign key to Courier)
// 		&BusinessObjects.ProductDiscount{}, // ProductDiscounts (foreign key to Product, Discount)
// 		&BusinessObjects.Review{},          // Reviews (foreign key to Product, User)
// 		&BusinessObjects.Transaction{},     // Transactions (foreign key to Order)
// 	)
// 	if err != nil {
// 		log.Fatalf("Failed to migrate the database: %v", err)
// 	}

// 	log.Println("Database migration completed successfully!")
// }
