-- Create the database
CREATE DATABASE EcommercePlatformSendoDb;
GO

-- Use the newly created database
USE EcommercePlatformSendoDb;
GO

-- Create Categories table first as it is referenced by Products
CREATE TABLE Categories (
    CategoryID VARCHAR(10) PRIMARY KEY,
    CategoryName NVARCHAR(100) NOT NULL
);
GO

-- Create Users table
CREATE TABLE Users (
    UserID VARCHAR(10) PRIMARY KEY,
    Username NVARCHAR(50) NOT NULL UNIQUE,
    PasswordHash NVARCHAR(255) NOT NULL,
    Email NVARCHAR(100) UNIQUE NOT NULL,
    FullName NVARCHAR(100),
    PhoneNumber NVARCHAR(20),
    Address NVARCHAR(255),
    UserType NVARCHAR(20) CHECK (UserType IN ('Customer', 'Seller', 'Admin')),
    CreatedAt DATETIME DEFAULT GETDATE()
);
GO

-- Create Products table
CREATE TABLE Products (
    ProductID VARCHAR(10) PRIMARY KEY,
    SellerID VARCHAR(10) FOREIGN KEY REFERENCES Users(UserID),
    ProductName NVARCHAR(100) NOT NULL,
    Description NVARCHAR(MAX),
    Price DECIMAL(18, 2) NOT NULL CHECK (Price >= 0),
    Quantity INT NOT NULL CHECK (Quantity >= 0),
    CategoryID VARCHAR(10) FOREIGN KEY REFERENCES Categories(CategoryID),
    CreatedAt DATETIME DEFAULT GETDATE(),
    UpdatedAt DATETIME DEFAULT GETDATE()
);
GO

-- Create Vouchers table
CREATE TABLE Vouchers (
    VoucherID VARCHAR(10) PRIMARY KEY,
    VoucherCode NVARCHAR(50) UNIQUE NOT NULL,
    DiscountType NVARCHAR(20) CHECK (DiscountType IN ('Percentage', 'FixedAmount', 'FreeShipping')) NOT NULL,
    DiscountValue DECIMAL(18, 2) NOT NULL CHECK (DiscountValue >= 0),
    MinimumOrderAmount DECIMAL(18, 2) CHECK (MinimumOrderAmount >= 0),
    MaxDiscountAmount DECIMAL(18, 2) CHECK (MaxDiscountAmount >= 0),
    StartDate DATETIME NOT NULL,
    EndDate DATETIME NOT NULL,
    UsageLimit INT CHECK (UsageLimit >= 0),
    UsageCount INT DEFAULT 0 CHECK (UsageCount >= 0),
    IsActive BIT DEFAULT 1
);
go
CREATE TABLE Couriers (
    CourierID VARCHAR(10) PRIMARY KEY,
    Courier NVARCHAR(100),  -- Name of the courier, e.g., FedEx, UPS
);
GO
-- Create FreightRates table for managing dynamic shipping costs
CREATE TABLE FreightRates (
    RateID VARCHAR(10) PRIMARY KEY,
    CourierID VARCHAR(10) FOREIGN KEY REFERENCES Couriers(CourierID),  -- e.g., Standard, Express
    DistanceMinKM INT NOT NULL,  -- Minimum distance range in kilometers
    DistanceMaxKM INT NOT NULL,  -- Maximum distance range in kilometers
    CostPerKM DECIMAL(18, 2) NOT NULL CHECK (CostPerKM >= 0),  -- Cost per kilometer within this range
);

GO
-- Create Orders table
CREATE TABLE Orders (
    OrderID VARCHAR(10) PRIMARY KEY,
    CustomerID VARCHAR(10) FOREIGN KEY REFERENCES Users(UserID),
    OrderDate DATETIME DEFAULT GETDATE(),
    TotalAmount DECIMAL(18, 2) NOT NULL CHECK (TotalAmount >= 0),
    OrderStatus NVARCHAR(50) CHECK (OrderStatus IN ('Pending', 'Shipped', 'Delivered', 'Canceled')),
    ShippingAddress NVARCHAR(255) NOT NULL,
    CourierID VARCHAR(10) FOREIGN KEY REFERENCES Couriers(CourierID),
	FreightPrice DECIMAL(18, 2),
    EstimatedDeliveryDate DATETIME,  -- Estimated delivery date
    ActualDeliveryDate DATETIME,  -- Actual delivery date
    PaymentMethod NVARCHAR(50) NOT NULL,
    PaymentStatus NVARCHAR(50) CHECK (PaymentStatus IN ('Paid', 'Pending', 'Failed')),
    VoucherID VARCHAR(10) FOREIGN KEY REFERENCES Vouchers(VoucherID)
);
GO
-- Create OrderDetails table with composite primary key
CREATE TABLE OrderDetails (
    OrderID VARCHAR(10),
    ProductID VARCHAR(10),
    Quantity INT NOT NULL CHECK (Quantity > 0),
    UnitPrice DECIMAL(18, 2) NOT NULL CHECK (UnitPrice >= 0),
    PRIMARY KEY (OrderID, ProductID),
    FOREIGN KEY (OrderID) REFERENCES Orders(OrderID),
    FOREIGN KEY (ProductID) REFERENCES Products(ProductID)
);
GO

-- Create VoucherRedemptions table
CREATE TABLE VoucherRedemptions (
    RedemptionID VARCHAR(10) PRIMARY KEY,
    VoucherID VARCHAR(10) FOREIGN KEY REFERENCES Vouchers(VoucherID),
    UserID VARCHAR(10) FOREIGN KEY REFERENCES Users(UserID),
    OrderID VARCHAR(10) FOREIGN KEY REFERENCES Orders(OrderID),
    RedeemedAt DATETIME DEFAULT GETDATE()
);
GO

-- Create ShoppingCart table
CREATE TABLE ShoppingCart (
    CartID VARCHAR(10) PRIMARY KEY,
    UserID VARCHAR(10) FOREIGN KEY REFERENCES Users(UserID),
    CreatedAt DATETIME DEFAULT GETDATE()
);
GO

-- Create CartItems table
CREATE TABLE CartItems (
    CartItemID VARCHAR(10) PRIMARY KEY,
    CartID VARCHAR(10) FOREIGN KEY REFERENCES ShoppingCart(CartID),
    ProductID VARCHAR(10) FOREIGN KEY REFERENCES Products(ProductID),
    Quantity INT NOT NULL CHECK (Quantity > 0)
);
GO

-- Create Reviews table
CREATE TABLE Reviews (
    ReviewID VARCHAR(10) PRIMARY KEY,
    ProductID VARCHAR(10) FOREIGN KEY REFERENCES Products(ProductID),
    UserID VARCHAR(10) FOREIGN KEY REFERENCES Users(UserID),
    Rating INT CHECK (Rating BETWEEN 1 AND 5),
    Comment NVARCHAR(MAX),
    CreatedAt DATETIME DEFAULT GETDATE()
);
GO

-- Create Transactions table
CREATE TABLE Transactions (
    TransactionID VARCHAR(10) PRIMARY KEY,
    OrderID VARCHAR(10) FOREIGN KEY REFERENCES Orders(OrderID),
    PaymentAmount DECIMAL(18, 2) NOT NULL CHECK (PaymentAmount >= 0),
    TransactionDate DATETIME DEFAULT GETDATE(),
    PaymentMethod NVARCHAR(50),
    PaymentStatus NVARCHAR(50) CHECK (PaymentStatus IN ('Success', 'Failed'))
);
GO
-- Create News table
CREATE TABLE News (
    NewsID VARCHAR(10) PRIMARY KEY,
    Title NVARCHAR(255) NOT NULL,
    Content NVARCHAR(MAX) NOT NULL,
    PublishedDate DATETIME DEFAULT GETDATE(),
    AuthorID VARCHAR(10) FOREIGN KEY REFERENCES Users(UserID),
    Status NVARCHAR(20) CHECK (Status IN ('Draft', 'Published', 'Archived')) DEFAULT 'Draft',
    ImageURL NVARCHAR(255),
    Category NVARCHAR(100)
);
GO

-- Create Discounts table
CREATE TABLE Discounts (
    DiscountID VARCHAR(10) PRIMARY KEY,
    DiscountType NVARCHAR(20) CHECK (DiscountType IN ('Percentage', 'FixedAmount')) NOT NULL,
    DiscountValue DECIMAL(18, 2) NOT NULL CHECK (DiscountValue >= 0),
    StartDate DATETIME NOT NULL,
    EndDate DATETIME NOT NULL
);
GO

-- Create ProductDiscounts table
CREATE TABLE ProductDiscounts (
    ProductID VARCHAR(10) FOREIGN KEY REFERENCES Products(ProductID),
    DiscountID VARCHAR(10) FOREIGN KEY REFERENCES Discounts(DiscountID)
	PRIMARY KEY (ProductID, DiscountID),
);
GO

