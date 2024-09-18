package API

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"th3y3m/e-commerce-platform/Services"
	"th3y3m/e-commerce-platform/Util"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func GetUsers(c *gin.Context) {
	searchValue := c.DefaultQuery("searchValue", "")
	sortBy := c.DefaultQuery("sortBy", "")
	pageIndex, _ := strconv.Atoi(c.DefaultQuery("pageIndex", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	statusParam := c.DefaultQuery("status", "")
	var status *bool

	if statusParam != "" {
		statusValue := statusParam == "true"
		status = &statusValue
	}

	users, err := Services.GetPaginatedUserList(searchValue, sortBy, pageIndex, pageSize, status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	user, err := Services.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func UpdateProfile(c *gin.Context) {
	ID := c.Param("id")

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	firebase := os.Getenv("FIREBASE")

	var info struct {
		FullName    string `form:"full_name" binding:"required"`
		PhoneNumber string `form:"phone_number" binding:"required"`
		Address     string `form:"address" binding:"required"`
	}
	if err := c.ShouldBind(&info); err != nil {
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
	bucketName := firebase
	objectName := "products/" + file.Filename
	publicURL, err := Util.UploadFileToFireBase(bucketName, objectName, tempFilePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload file to Firebase"})
		return
	}
	err = Services.UpdateProfile(ID, info.FullName, info.PhoneNumber, info.Address, publicURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func BanUser(c *gin.Context) {
	id := c.Param("id")
	err := Services.BanUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User banned successfully"})
}

func UnBanUser(c *gin.Context) {
	id := c.Param("id")
	err := Services.UnBanUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User unbanned successfully"})
}
