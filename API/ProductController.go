package API

import (
	"net/http"
	"strconv"
	"th3y3m/e-commerce-platform/Services"
	"th3y3m/e-commerce-platform/Util"

	"github.com/gin-gonic/gin"
)

func GetPaginatedProductList(c *gin.Context) {

	searchValue := c.DefaultQuery("searchValue", "")
	sortBy := c.DefaultQuery("sortBy", "")
	productID := c.DefaultQuery("productID", "")
	sellerID := c.DefaultQuery("sellerID", "")
	categoryID := c.DefaultQuery("categoryID", "")
	pageIndex, _ := strconv.Atoi(c.DefaultQuery("pageIndex", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	statusParam := c.DefaultQuery("status", "")
	var status *bool

	if statusParam != "" {
		statusValue := statusParam == "true"
		status = &statusValue
	}

	products, err := Services.GetPaginatedProductList(searchValue, sortBy, productID, sellerID, categoryID, pageIndex, pageSize, status)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}

func GetProductByID(c *gin.Context) {
	id := c.Param("id")
	product, err := Services.GetProductByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, product)
}

func CreateProduct(c *gin.Context) {
	var product struct {
		SellerID    string  `form:"seller_id" binding:"required"`
		ProductName string  `form:"product_name" binding:"required"`
		Description string  `form:"description" binding:"required"`
		CategoryID  string  `form:"category_id" binding:"required"`
		Price       float64 `form:"price" binding:"required"`
		Quantity    int     `form:"quantity" binding:"required"`
	}

	// Bind form data
	if err := c.ShouldBind(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Retrieve the file from the form data
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File upload failed"})
		return
	}

	// Save the file temporarily
	tempFilePath := "./uploads/" + file.Filename
	if err := c.SaveUploadedFile(file, tempFilePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	// Upload the file to Firebase
	bucketName := "sendo-a5204.appspot.com"
	objectName := "products/" + file.Filename
	publicURL, err := Util.UploadFileToFireBase(bucketName, objectName, tempFilePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload file to Firebase"})
		return
	}

	// Call the service to create the product
	err = Services.CreateProduct(product.SellerID, product.ProductName, product.Description, product.CategoryID, publicURL, product.Price, product.Quantity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product created successfully", "file_url": publicURL})
}

func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product struct {
		SellerID    string  `json:"seller_id" binding:"required"`
		ProductName string  `json:"product_name" binding:"required"`
		Description string  `json:"description" binding:"required"`
		CategoryID  string  `json:"category_id" binding:"required"`
		Price       float64 `json:"price" binding:"required"`
		Quantity    int     `json:"quantity" binding:"required"`
	}
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := Services.UpdateProduct(id, product.SellerID, product.ProductName, product.Description, product.CategoryID, product.Price, product.Quantity)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	err := Services.DeleteProduct(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}

func GetProductPriceAfterDiscount(c *gin.Context) {
	id := c.Param("id")
	price, err := Services.GetProductPriceAfterDiscount(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"price": price})
}
