USE [master]
GO
/****** Object:  Database [EcommercePlatformSendoDb]    Script Date: 9/9/2024 11:31:13 AM ******/
CREATE DATABASE [EcommercePlatformSendoDb]
 CONTAINMENT = NONE
 ON  PRIMARY 
( NAME = N'EcommercePlatformSendoDb', FILENAME = N'c:\Program Files\Microsoft SQL Server\MSSQL11.SQLEXPRESS\MSSQL\DATA\EcommercePlatformSendoDb.mdf' , SIZE = 3136KB , MAXSIZE = UNLIMITED, FILEGROWTH = 1024KB )
 LOG ON 
( NAME = N'EcommercePlatformSendoDb_log', FILENAME = N'c:\Program Files\Microsoft SQL Server\MSSQL11.SQLEXPRESS\MSSQL\DATA\EcommercePlatformSendoDb_log.ldf' , SIZE = 784KB , MAXSIZE = 2048GB , FILEGROWTH = 10%)
GO
ALTER DATABASE [EcommercePlatformSendoDb] SET COMPATIBILITY_LEVEL = 110
GO
IF (1 = FULLTEXTSERVICEPROPERTY('IsFullTextInstalled'))
begin
EXEC [EcommercePlatformSendoDb].[dbo].[sp_fulltext_database] @action = 'enable'
end
GO
ALTER DATABASE [EcommercePlatformSendoDb] SET ANSI_NULL_DEFAULT OFF 
GO
ALTER DATABASE [EcommercePlatformSendoDb] SET ANSI_NULLS OFF 
GO
ALTER DATABASE [EcommercePlatformSendoDb] SET ANSI_PADDING OFF 
GO
ALTER DATABASE [EcommercePlatformSendoDb] SET ANSI_WARNINGS OFF 
GO
ALTER DATABASE [EcommercePlatformSendoDb] SET ARITHABORT OFF 
GO
ALTER DATABASE [EcommercePlatformSendoDb] SET AUTO_CLOSE ON 
GO
ALTER DATABASE [EcommercePlatformSendoDb] SET AUTO_CREATE_STATISTICS ON 
GO
ALTER DATABASE [EcommercePlatformSendoDb] SET AUTO_SHRINK OFF 
GO
ALTER DATABASE [EcommercePlatformSendoDb] SET AUTO_UPDATE_STATISTICS ON 
GO
ALTER DATABASE [EcommercePlatformSendoDb] SET CURSOR_CLOSE_ON_COMMIT OFF 
GO
ALTER DATABASE [EcommercePlatformSendoDb] SET CURSOR_DEFAULT  GLOBAL 
GO
ALTER DATABASE [EcommercePlatformSendoDb] SET CONCAT_NULL_YIELDS_NULL OFF 
GO
ALTER DATABASE [EcommercePlatformSendoDb] SET NUMERIC_ROUNDABORT OFF 
GO
ALTER DATABASE [EcommercePlatformSendoDb] SET QUOTED_IDENTIFIER OFF 
GO
ALTER DATABASE [EcommercePlatformSendoDb] SET RECURSIVE_TRIGGERS OFF 
GO
ALTER DATABASE [EcommercePlatformSendoDb] SET  ENABLE_BROKER 
GO
ALTER DATABASE [EcommercePlatformSendoDb] SET AUTO_UPDATE_STATISTICS_ASYNC OFF 
GO
ALTER DATABASE [EcommercePlatformSendoDb] SET DATE_CORRELATION_OPTIMIZATION OFF 
GO
ALTER DATABASE [EcommercePlatformSendoDb] SET TRUSTWORTHY OFF 
GO
ALTER DATABASE [EcommercePlatformSendoDb] SET ALLOW_SNAPSHOT_ISOLATION OFF 
GO
ALTER DATABASE [EcommercePlatformSendoDb] SET PARAMETERIZATION SIMPLE 
GO
ALTER DATABASE [EcommercePlatformSendoDb] SET READ_COMMITTED_SNAPSHOT OFF 
GO
ALTER DATABASE [EcommercePlatformSendoDb] SET HONOR_BROKER_PRIORITY OFF 
GO
ALTER DATABASE [EcommercePlatformSendoDb] SET RECOVERY SIMPLE 
GO
ALTER DATABASE [EcommercePlatformSendoDb] SET  MULTI_USER 
GO
ALTER DATABASE [EcommercePlatformSendoDb] SET PAGE_VERIFY CHECKSUM  
GO
ALTER DATABASE [EcommercePlatformSendoDb] SET DB_CHAINING OFF 
GO
ALTER DATABASE [EcommercePlatformSendoDb] SET FILESTREAM( NON_TRANSACTED_ACCESS = OFF ) 
GO
ALTER DATABASE [EcommercePlatformSendoDb] SET TARGET_RECOVERY_TIME = 0 SECONDS 
GO
USE [EcommercePlatformSendoDb]
GO
/****** Object:  Table [dbo].[CartItems]    Script Date: 9/9/2024 11:31:13 AM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
SET ANSI_PADDING ON
GO
CREATE TABLE [dbo].[CartItems](
	[CartItemID] [varchar](10) NOT NULL,
	[CartID] [varchar](10) NULL,
	[ProductID] [varchar](10) NULL,
	[Quantity] [int] NOT NULL,
PRIMARY KEY CLUSTERED 
(
	[CartItemID] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY]

GO
SET ANSI_PADDING OFF
GO
/****** Object:  Table [dbo].[Categories]    Script Date: 9/9/2024 11:31:13 AM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
SET ANSI_PADDING ON
GO
CREATE TABLE [dbo].[Categories](
	[CategoryID] [varchar](10) NOT NULL,
	[CategoryName] [nvarchar](100) NOT NULL,
PRIMARY KEY CLUSTERED 
(
	[CategoryID] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY]

GO
SET ANSI_PADDING OFF
GO
/****** Object:  Table [dbo].[Discounts]    Script Date: 9/9/2024 11:31:13 AM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
SET ANSI_PADDING ON
GO
CREATE TABLE [dbo].[Discounts](
	[DiscountID] [varchar](10) NOT NULL,
	[DiscountType] [nvarchar](20) NOT NULL,
	[DiscountValue] [decimal](18, 2) NOT NULL,
	[StartDate] [datetime] NOT NULL,
	[EndDate] [datetime] NOT NULL,
PRIMARY KEY CLUSTERED 
(
	[DiscountID] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY]

GO
SET ANSI_PADDING OFF
GO
/****** Object:  Table [dbo].[FreightRates]    Script Date: 9/9/2024 11:31:13 AM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
SET ANSI_PADDING ON
GO
CREATE TABLE [dbo].[FreightRates](
	[RateID] [varchar](10) NOT NULL,
	[Courier] [nvarchar](100) NULL,
	[ShippingMethod] [nvarchar](50) NULL,
	[DistanceMinKM] [int] NOT NULL,
	[DistanceMaxKM] [int] NOT NULL,
	[CostPerKM] [decimal](18, 2) NOT NULL,
PRIMARY KEY CLUSTERED 
(
	[RateID] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY]

GO
SET ANSI_PADDING OFF
GO
/****** Object:  Table [dbo].[News]    Script Date: 9/9/2024 11:31:13 AM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
SET ANSI_PADDING ON
GO
CREATE TABLE [dbo].[News](
	[NewsID] [varchar](10) NOT NULL,
	[Title] [nvarchar](255) NOT NULL,
	[Content] [nvarchar](max) NOT NULL,
	[PublishedDate] [datetime] NULL,
	[AuthorID] [varchar](10) NULL,
	[Status] [nvarchar](20) NULL,
	[ImageURL] [nvarchar](255) NULL,
	[Category] [nvarchar](100) NULL,
PRIMARY KEY CLUSTERED 
(
	[NewsID] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY] TEXTIMAGE_ON [PRIMARY]

GO
SET ANSI_PADDING OFF
GO
/****** Object:  Table [dbo].[OrderDetails]    Script Date: 9/9/2024 11:31:13 AM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
SET ANSI_PADDING ON
GO
CREATE TABLE [dbo].[OrderDetails](
	[OrderID] [varchar](10) NOT NULL,
	[ProductID] [varchar](10) NOT NULL,
	[Quantity] [int] NOT NULL,
	[UnitPrice] [decimal](18, 2) NOT NULL,
PRIMARY KEY CLUSTERED 
(
	[OrderID] ASC,
	[ProductID] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY]

GO
SET ANSI_PADDING OFF
GO
/****** Object:  Table [dbo].[Orders]    Script Date: 9/9/2024 11:31:13 AM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
SET ANSI_PADDING ON
GO
CREATE TABLE [dbo].[Orders](
	[OrderID] [varchar](10) NOT NULL,
	[CustomerID] [varchar](10) NULL,
	[OrderDate] [datetime] NULL,
	[TotalAmount] [decimal](18, 2) NOT NULL,
	[OrderStatus] [nvarchar](50) NULL,
	[ShippingAddress] [nvarchar](255) NOT NULL,
	[RateID] [varchar](10) NULL,
	[EstimatedDeliveryDate] [datetime] NULL,
	[ActualDeliveryDate] [datetime] NULL,
	[PaymentMethod] [nvarchar](50) NOT NULL,
	[PaymentStatus] [nvarchar](50) NULL,
	[VoucherID] [varchar](10) NULL,
PRIMARY KEY CLUSTERED 
(
	[OrderID] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY]

GO
SET ANSI_PADDING OFF
GO
/****** Object:  Table [dbo].[ProductDiscounts]    Script Date: 9/9/2024 11:31:13 AM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
SET ANSI_PADDING ON
GO
CREATE TABLE [dbo].[ProductDiscounts](
	[ProductDiscountID] [varchar](10) NOT NULL,
	[ProductID] [varchar](10) NULL,
	[DiscountID] [varchar](10) NULL,
	[StartDate] [datetime] NOT NULL,
	[EndDate] [datetime] NOT NULL,
PRIMARY KEY CLUSTERED 
(
	[ProductDiscountID] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY]

GO
SET ANSI_PADDING OFF
GO
/****** Object:  Table [dbo].[Products]    Script Date: 9/9/2024 11:31:13 AM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
SET ANSI_PADDING ON
GO
CREATE TABLE [dbo].[Products](
	[ProductID] [varchar](10) NOT NULL,
	[SellerID] [varchar](10) NULL,
	[ProductName] [nvarchar](100) NOT NULL,
	[Description] [nvarchar](max) NULL,
	[Price] [decimal](18, 2) NOT NULL,
	[Quantity] [int] NOT NULL,
	[CategoryID] [varchar](10) NULL,
	[CreatedAt] [datetime] NULL,
	[UpdatedAt] [datetime] NULL,
PRIMARY KEY CLUSTERED 
(
	[ProductID] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY] TEXTIMAGE_ON [PRIMARY]

GO
SET ANSI_PADDING OFF
GO
/****** Object:  Table [dbo].[Reviews]    Script Date: 9/9/2024 11:31:13 AM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
SET ANSI_PADDING ON
GO
CREATE TABLE [dbo].[Reviews](
	[ReviewID] [varchar](10) NOT NULL,
	[ProductID] [varchar](10) NULL,
	[UserID] [varchar](10) NULL,
	[Rating] [int] NULL,
	[Comment] [nvarchar](max) NULL,
	[CreatedAt] [datetime] NULL,
PRIMARY KEY CLUSTERED 
(
	[ReviewID] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY] TEXTIMAGE_ON [PRIMARY]

GO
SET ANSI_PADDING OFF
GO
/****** Object:  Table [dbo].[ShoppingCart]    Script Date: 9/9/2024 11:31:13 AM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
SET ANSI_PADDING ON
GO
CREATE TABLE [dbo].[ShoppingCart](
	[CartID] [varchar](10) NOT NULL,
	[UserID] [varchar](10) NULL,
	[CreatedAt] [datetime] NULL,
PRIMARY KEY CLUSTERED 
(
	[CartID] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY]

GO
SET ANSI_PADDING OFF
GO
/****** Object:  Table [dbo].[Transactions]    Script Date: 9/9/2024 11:31:13 AM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
SET ANSI_PADDING ON
GO
CREATE TABLE [dbo].[Transactions](
	[TransactionID] [varchar](10) NOT NULL,
	[OrderID] [varchar](10) NULL,
	[PaymentAmount] [decimal](18, 2) NOT NULL,
	[TransactionDate] [datetime] NULL,
	[PaymentMethod] [nvarchar](50) NULL,
	[PaymentStatus] [nvarchar](50) NULL,
PRIMARY KEY CLUSTERED 
(
	[TransactionID] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY]

GO
SET ANSI_PADDING OFF
GO
/****** Object:  Table [dbo].[Users]    Script Date: 9/9/2024 11:31:13 AM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
SET ANSI_PADDING ON
GO
CREATE TABLE [dbo].[Users](
	[UserID] [varchar](10) NOT NULL,
	[Username] [nvarchar](50) NOT NULL,
	[PasswordHash] [nvarchar](255) NOT NULL,
	[Email] [nvarchar](100) NOT NULL,
	[FullName] [nvarchar](100) NULL,
	[PhoneNumber] [nvarchar](20) NULL,
	[Address] [nvarchar](255) NULL,
	[UserType] [nvarchar](20) NULL,
	[CreatedAt] [datetime] NULL,
PRIMARY KEY CLUSTERED 
(
	[UserID] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY],
UNIQUE NONCLUSTERED 
(
	[Username] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY],
UNIQUE NONCLUSTERED 
(
	[Email] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY]

GO
SET ANSI_PADDING OFF
GO
/****** Object:  Table [dbo].[VoucherRedemptions]    Script Date: 9/9/2024 11:31:13 AM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
SET ANSI_PADDING ON
GO
CREATE TABLE [dbo].[VoucherRedemptions](
	[RedemptionID] [varchar](10) NOT NULL,
	[VoucherID] [varchar](10) NULL,
	[UserID] [varchar](10) NULL,
	[OrderID] [varchar](10) NULL,
	[RedeemedAt] [datetime] NULL,
PRIMARY KEY CLUSTERED 
(
	[RedemptionID] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY]

GO
SET ANSI_PADDING OFF
GO
/****** Object:  Table [dbo].[Vouchers]    Script Date: 9/9/2024 11:31:13 AM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
SET ANSI_PADDING ON
GO
CREATE TABLE [dbo].[Vouchers](
	[VoucherID] [varchar](10) NOT NULL,
	[VoucherCode] [nvarchar](50) NOT NULL,
	[DiscountType] [nvarchar](20) NOT NULL,
	[DiscountValue] [decimal](18, 2) NOT NULL,
	[MinimumOrderAmount] [decimal](18, 2) NULL,
	[MaxDiscountAmount] [decimal](18, 2) NULL,
	[StartDate] [datetime] NOT NULL,
	[EndDate] [datetime] NOT NULL,
	[UsageLimit] [int] NULL,
	[UsageCount] [int] NULL,
	[IsActive] [bit] NULL,
PRIMARY KEY CLUSTERED 
(
	[VoucherID] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY],
UNIQUE NONCLUSTERED 
(
	[VoucherCode] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY]

GO
SET ANSI_PADDING OFF
GO
ALTER TABLE [dbo].[News] ADD  DEFAULT (getdate()) FOR [PublishedDate]
GO
ALTER TABLE [dbo].[News] ADD  DEFAULT ('Draft') FOR [Status]
GO
ALTER TABLE [dbo].[Orders] ADD  DEFAULT (getdate()) FOR [OrderDate]
GO
ALTER TABLE [dbo].[Products] ADD  DEFAULT (getdate()) FOR [CreatedAt]
GO
ALTER TABLE [dbo].[Products] ADD  DEFAULT (getdate()) FOR [UpdatedAt]
GO
ALTER TABLE [dbo].[Reviews] ADD  DEFAULT (getdate()) FOR [CreatedAt]
GO
ALTER TABLE [dbo].[ShoppingCart] ADD  DEFAULT (getdate()) FOR [CreatedAt]
GO
ALTER TABLE [dbo].[Transactions] ADD  DEFAULT (getdate()) FOR [TransactionDate]
GO
ALTER TABLE [dbo].[Users] ADD  DEFAULT (getdate()) FOR [CreatedAt]
GO
ALTER TABLE [dbo].[VoucherRedemptions] ADD  DEFAULT (getdate()) FOR [RedeemedAt]
GO
ALTER TABLE [dbo].[Vouchers] ADD  DEFAULT ((0)) FOR [UsageCount]
GO
ALTER TABLE [dbo].[Vouchers] ADD  DEFAULT ((1)) FOR [IsActive]
GO
ALTER TABLE [dbo].[CartItems]  WITH CHECK ADD FOREIGN KEY([CartID])
REFERENCES [dbo].[ShoppingCart] ([CartID])
GO
ALTER TABLE [dbo].[CartItems]  WITH CHECK ADD FOREIGN KEY([ProductID])
REFERENCES [dbo].[Products] ([ProductID])
GO
ALTER TABLE [dbo].[News]  WITH CHECK ADD FOREIGN KEY([AuthorID])
REFERENCES [dbo].[Users] ([UserID])
GO
ALTER TABLE [dbo].[OrderDetails]  WITH CHECK ADD FOREIGN KEY([OrderID])
REFERENCES [dbo].[Orders] ([OrderID])
GO
ALTER TABLE [dbo].[OrderDetails]  WITH CHECK ADD FOREIGN KEY([ProductID])
REFERENCES [dbo].[Products] ([ProductID])
GO
ALTER TABLE [dbo].[Orders]  WITH CHECK ADD FOREIGN KEY([CustomerID])
REFERENCES [dbo].[Users] ([UserID])
GO
ALTER TABLE [dbo].[Orders]  WITH CHECK ADD FOREIGN KEY([RateID])
REFERENCES [dbo].[FreightRates] ([RateID])
GO
ALTER TABLE [dbo].[Orders]  WITH CHECK ADD FOREIGN KEY([VoucherID])
REFERENCES [dbo].[Vouchers] ([VoucherID])
GO
ALTER TABLE [dbo].[ProductDiscounts]  WITH CHECK ADD FOREIGN KEY([DiscountID])
REFERENCES [dbo].[Discounts] ([DiscountID])
GO
ALTER TABLE [dbo].[ProductDiscounts]  WITH CHECK ADD FOREIGN KEY([ProductID])
REFERENCES [dbo].[Products] ([ProductID])
GO
ALTER TABLE [dbo].[Products]  WITH CHECK ADD FOREIGN KEY([CategoryID])
REFERENCES [dbo].[Categories] ([CategoryID])
GO
ALTER TABLE [dbo].[Products]  WITH CHECK ADD FOREIGN KEY([SellerID])
REFERENCES [dbo].[Users] ([UserID])
GO
ALTER TABLE [dbo].[Reviews]  WITH CHECK ADD FOREIGN KEY([ProductID])
REFERENCES [dbo].[Products] ([ProductID])
GO
ALTER TABLE [dbo].[Reviews]  WITH CHECK ADD FOREIGN KEY([UserID])
REFERENCES [dbo].[Users] ([UserID])
GO
ALTER TABLE [dbo].[ShoppingCart]  WITH CHECK ADD FOREIGN KEY([UserID])
REFERENCES [dbo].[Users] ([UserID])
GO
ALTER TABLE [dbo].[Transactions]  WITH CHECK ADD FOREIGN KEY([OrderID])
REFERENCES [dbo].[Orders] ([OrderID])
GO
ALTER TABLE [dbo].[VoucherRedemptions]  WITH CHECK ADD FOREIGN KEY([OrderID])
REFERENCES [dbo].[Orders] ([OrderID])
GO
ALTER TABLE [dbo].[VoucherRedemptions]  WITH CHECK ADD FOREIGN KEY([UserID])
REFERENCES [dbo].[Users] ([UserID])
GO
ALTER TABLE [dbo].[VoucherRedemptions]  WITH CHECK ADD FOREIGN KEY([VoucherID])
REFERENCES [dbo].[Vouchers] ([VoucherID])
GO
ALTER TABLE [dbo].[CartItems]  WITH CHECK ADD CHECK  (([Quantity]>(0)))
GO
ALTER TABLE [dbo].[Discounts]  WITH CHECK ADD CHECK  (([DiscountType]='FixedAmount' OR [DiscountType]='Percentage'))
GO
ALTER TABLE [dbo].[Discounts]  WITH CHECK ADD CHECK  (([DiscountValue]>=(0)))
GO
ALTER TABLE [dbo].[FreightRates]  WITH CHECK ADD CHECK  (([CostPerKM]>=(0)))
GO
ALTER TABLE [dbo].[News]  WITH CHECK ADD CHECK  (([Status]='Archived' OR [Status]='Published' OR [Status]='Draft'))
GO
ALTER TABLE [dbo].[OrderDetails]  WITH CHECK ADD CHECK  (([Quantity]>(0)))
GO
ALTER TABLE [dbo].[OrderDetails]  WITH CHECK ADD CHECK  (([UnitPrice]>=(0)))
GO
ALTER TABLE [dbo].[Orders]  WITH CHECK ADD CHECK  (([OrderStatus]='Canceled' OR [OrderStatus]='Delivered' OR [OrderStatus]='Shipped' OR [OrderStatus]='Pending'))
GO
ALTER TABLE [dbo].[Orders]  WITH CHECK ADD CHECK  (([PaymentStatus]='Failed' OR [PaymentStatus]='Pending' OR [PaymentStatus]='Paid'))
GO
ALTER TABLE [dbo].[Orders]  WITH CHECK ADD CHECK  (([TotalAmount]>=(0)))
GO
ALTER TABLE [dbo].[Products]  WITH CHECK ADD CHECK  (([Price]>=(0)))
GO
ALTER TABLE [dbo].[Products]  WITH CHECK ADD CHECK  (([Quantity]>=(0)))
GO
ALTER TABLE [dbo].[Reviews]  WITH CHECK ADD CHECK  (([Rating]>=(1) AND [Rating]<=(5)))
GO
ALTER TABLE [dbo].[Transactions]  WITH CHECK ADD CHECK  (([PaymentAmount]>=(0)))
GO
ALTER TABLE [dbo].[Transactions]  WITH CHECK ADD CHECK  (([PaymentStatus]='Failed' OR [PaymentStatus]='Success'))
GO
ALTER TABLE [dbo].[Users]  WITH CHECK ADD CHECK  (([UserType]='Admin' OR [UserType]='Seller' OR [UserType]='Customer'))
GO
ALTER TABLE [dbo].[Vouchers]  WITH CHECK ADD CHECK  (([DiscountType]='FreeShipping' OR [DiscountType]='FixedAmount' OR [DiscountType]='Percentage'))
GO
ALTER TABLE [dbo].[Vouchers]  WITH CHECK ADD CHECK  (([DiscountValue]>=(0)))
GO
ALTER TABLE [dbo].[Vouchers]  WITH CHECK ADD CHECK  (([MaxDiscountAmount]>=(0)))
GO
ALTER TABLE [dbo].[Vouchers]  WITH CHECK ADD CHECK  (([MinimumOrderAmount]>=(0)))
GO
ALTER TABLE [dbo].[Vouchers]  WITH CHECK ADD CHECK  (([UsageCount]>=(0)))
GO
ALTER TABLE [dbo].[Vouchers]  WITH CHECK ADD CHECK  (([UsageLimit]>=(0)))
GO
USE [master]
GO
ALTER DATABASE [EcommercePlatformSendoDb] SET  READ_WRITE 
GO
