-- Create the database
CREATE DATABASE "EcommercePlatformSendoDbV1"
WITH
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'English_United States.utf8'
    LC_CTYPE = 'English_United States.utf8'
    LOCALE_PROVIDER = 'libc'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1
    IS_TEMPLATE = False;

-- Create Categories table
CREATE TABLE Categories (
    CategoryID VARCHAR(10) PRIMARY KEY,
    CategoryName VARCHAR(100) UNIQUE NOT NULL,
	Status BOOLEAN DEFAULT TRUE
);

-- Create Users table
CREATE TABLE Users (
    UserID VARCHAR(10) PRIMARY KEY,
    Username VARCHAR(50) NOT NULL UNIQUE,
    PasswordHash VARCHAR(255) NOT NULL,
    Email VARCHAR(100) UNIQUE NOT NULL,
    FullName VARCHAR(100),
    PhoneNumber VARCHAR(20),
    Address VARCHAR(255),
    UserType VARCHAR(20) CHECK (UserType IN ('Customer', 'Seller', 'Admin')),
    CreatedAt TIMESTAMP DEFAULT NOW(),
	Status BOOLEAN DEFAULT TRUE
);

-- Create Products table
CREATE TABLE Products (
    ProductID VARCHAR(10) PRIMARY KEY,
    SellerID VARCHAR(10) REFERENCES Users(UserID),
    ProductName VARCHAR(100) NOT NULL,
    Description TEXT,
    Price DECIMAL(18, 2) NOT NULL CHECK (Price >= 0),
    Quantity INT NOT NULL CHECK (Quantity >= 0),
    CategoryID VARCHAR(10) REFERENCES Categories(CategoryID),
    CreatedAt TIMESTAMP DEFAULT NOW(),
    UpdatedAt TIMESTAMP DEFAULT NOW(),
	Status BOOLEAN DEFAULT TRUE
);

-- Create Vouchers table
CREATE TABLE Vouchers (
    VoucherID VARCHAR(10) PRIMARY KEY,
    VoucherCode VARCHAR(50) UNIQUE NOT NULL,
    DiscountType VARCHAR(20) CHECK (DiscountType IN ('Percentage', 'FixedAmount', 'FreeShipping')) NOT NULL,
    DiscountValue DECIMAL(18, 2) NOT NULL CHECK (DiscountValue >= 0),
    MinimumOrderAmount DECIMAL(18, 2) CHECK (MinimumOrderAmount >= 0),
    MaxDiscountAmount DECIMAL(18, 2) CHECK (MaxDiscountAmount >= 0),
    StartDate TIMESTAMP NOT NULL,
    EndDate TIMESTAMP NOT NULL,
    UsageLimit INT CHECK (UsageLimit >= 0),
    UsageCount INT DEFAULT 0 CHECK (UsageCount >= 0),
    Status BOOLEAN DEFAULT TRUE
);

-- Create Couriers table
CREATE TABLE Couriers (
    CourierID VARCHAR(10) PRIMARY KEY,
    Courier VARCHAR(100),
	Status BOOLEAN DEFAULT TRUE
);

-- Create FreightRates table
CREATE TABLE FreightRates (
    RateID VARCHAR(10) PRIMARY KEY,
    CourierID VARCHAR(10) REFERENCES Couriers(CourierID),
    DistanceMinKM INT NOT NULL,
    DistanceMaxKM INT NOT NULL,
    CostPerKM DECIMAL(18, 2) NOT NULL CHECK (CostPerKM >= 0),
	Status BOOLEAN DEFAULT TRUE
);

-- Create Orders table
CREATE TABLE Orders (
    OrderID VARCHAR(10) PRIMARY KEY,
    CustomerID VARCHAR(10) REFERENCES Users(UserID),
    OrderDate TIMESTAMP DEFAULT NOW(),
    TotalAmount DECIMAL(18, 2) NOT NULL CHECK (TotalAmount >= 0),
    OrderStatus VARCHAR(50) CHECK (OrderStatus IN ('Pending', 'Shipped', 'Delivered', 'Canceled')),
    ShippingAddress VARCHAR(255) NOT NULL,
    CourierID VARCHAR(10) REFERENCES Couriers(CourierID),
    FreightPrice DECIMAL(18, 2),
    EstimatedDeliveryDate TIMESTAMP,
    ActualDeliveryDate TIMESTAMP,
    PaymentMethod VARCHAR(50) NOT NULL,
    PaymentStatus VARCHAR(50) CHECK (PaymentStatus IN ('Paid', 'Pending', 'Failed')),
    VoucherID VARCHAR(10) REFERENCES Vouchers(VoucherID)
);

-- Create OrderDetails table
CREATE TABLE OrderDetails (
    OrderID VARCHAR(10),
    ProductID VARCHAR(10),
    Quantity INT NOT NULL CHECK (Quantity > 0),
    UnitPrice DECIMAL(18, 2) NOT NULL CHECK (UnitPrice >= 0),
    PRIMARY KEY (OrderID, ProductID),
    FOREIGN KEY (OrderID) REFERENCES Orders(OrderID),
    FOREIGN KEY (ProductID) REFERENCES Products(ProductID)
);

-- Create VoucherRedemptions table
CREATE TABLE VoucherRedemptions (
    RedemptionID VARCHAR(10) PRIMARY KEY,
    VoucherID VARCHAR(10) REFERENCES Vouchers(VoucherID),
    UserID VARCHAR(10) REFERENCES Users(UserID),
    OrderID VARCHAR(10) REFERENCES Orders(OrderID),
    RedeemedAt TIMESTAMP DEFAULT NOW()
);

-- Create ShoppingCart table
CREATE TABLE ShoppingCart (
    CartID VARCHAR(10) PRIMARY KEY,
    UserID VARCHAR(10) REFERENCES Users(UserID),
    CreatedAt TIMESTAMP DEFAULT NOW(),
	Status BOOLEAN DEFAULT TRUE
);

-- Create CartItems table
CREATE TABLE CartItems (
    CartItemID VARCHAR(10) PRIMARY KEY,
    CartID VARCHAR(10) REFERENCES ShoppingCart(CartID),
    ProductID VARCHAR(10) REFERENCES Products(ProductID),
    Quantity INT NOT NULL CHECK (Quantity > 0),
);

-- Create Reviews table
CREATE TABLE Reviews (
    ReviewID VARCHAR(10) PRIMARY KEY,
    ProductID VARCHAR(10) REFERENCES Products(ProductID),
    UserID VARCHAR(10) REFERENCES Users(UserID),
    Rating INT CHECK (Rating BETWEEN 1 AND 5),
    Comment TEXT,
    CreatedAt TIMESTAMP DEFAULT NOW()
	Status BOOLEAN DEFAULT TRUE
);

-- Create Transactions table
CREATE TABLE Transactions (
    TransactionID VARCHAR(10) PRIMARY KEY,
    OrderID VARCHAR(10) REFERENCES Orders(OrderID),
    PaymentAmount DECIMAL(18, 2) NOT NULL CHECK (PaymentAmount >= 0),
    TransactionDate TIMESTAMP DEFAULT NOW(),
    PaymentMethod VARCHAR(50),
    PaymentStatus VARCHAR(50) CHECK (PaymentStatus IN ('Success', 'Failed'))
);

-- Create News table
CREATE TABLE News (
    NewsID VARCHAR(10) PRIMARY KEY,
    Title VARCHAR(255) NOT NULL,
    Content TEXT NOT NULL,
    PublishedDate TIMESTAMP DEFAULT NOW(),
    AuthorID VARCHAR(10) REFERENCES Users(UserID),
    Status VARCHAR(20) CHECK (Status IN ('Draft', 'Published', 'Archived')) DEFAULT 'Draft',
    ImageURL VARCHAR(255),
    Category VARCHAR(100),
);

-- Create Discounts table
CREATE TABLE Discounts (
    DiscountID VARCHAR(10) PRIMARY KEY,
    DiscountType VARCHAR(20) CHECK (DiscountType IN ('Percentage', 'FixedAmount')) NOT NULL,
    DiscountValue DECIMAL(18, 2) NOT NULL CHECK (DiscountValue >= 0),
    StartDate TIMESTAMP NOT NULL,
    EndDate TIMESTAMP NOT NULL
);

-- Create ProductDiscounts table
CREATE TABLE ProductDiscounts (
    ProductID VARCHAR(10) REFERENCES Products(ProductID),
    DiscountID VARCHAR(10) REFERENCES Discounts(DiscountID),
    PRIMARY KEY (ProductID, DiscountID)
);

