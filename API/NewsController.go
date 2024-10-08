package API

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"th3y3m/e-commerce-platform/DependencyInjection"
	"th3y3m/e-commerce-platform/Util"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func GetPaginatedNewsList(c *gin.Context) {
	module := DependencyInjection.NewNewsServiceProvider()
	searchValue := c.DefaultQuery("searchValue", "")
	sortBy := c.DefaultQuery("sortBy", "")
	pageIndex, _ := strconv.Atoi(c.DefaultQuery("pageIndex", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	newID := c.DefaultQuery("newID", "")
	authorID := c.DefaultQuery("authorID", "")
	status, _ := strconv.ParseBool(c.DefaultQuery("status", ""))

	news, err := module.GetPaginatedNewsList(c, searchValue, sortBy, newID, authorID, pageIndex, pageSize, &status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, news)
}

func GetAllNews(c *gin.Context) {
	module := DependencyInjection.NewNewsServiceProvider()
	news, err := module.GetAllNews(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, news)
}

func GetNewsByID(c *gin.Context) {
	module := DependencyInjection.NewNewsServiceProvider()

	id := c.Param("id")
	news, err := module.GetNewsByID(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, news)
}

func CreateNews(c *gin.Context) {
	module := DependencyInjection.NewNewsServiceProvider()
	// Load environment variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	firebase := os.Getenv("FIREBASE")

	// Define the news struct
	var news struct {
		Title    string `json:"title" binding:"required"`
		Content  string `json:"content" binding:"required"`
		AuthorID string `json:"author_id" binding:"required"`
		Category string `json:"category" binding:"required"`
	}

	// Bind JSON data
	if err := c.ShouldBindJSON(&news); err != nil {
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
	objectName := "news/" + file.Filename
	publicURL, err := Util.UploadFileToFireBase(bucketName, objectName, tempFilePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload file to Firebase"})
		return
	}

	// Call the service to create the news entry
	err = module.CreateNews(c, news.Title, news.Content, news.AuthorID, news.Category, publicURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "News created successfully", "file_url": publicURL})
}

func UpdateNews(c *gin.Context) {
	module := DependencyInjection.NewNewsServiceProvider()
	var news struct {
		Title    string `json:"title" binding:"required"`
		Content  string `json:"content" binding:"required"`
		AuthorID string `json:"author_id" binding:"required"`
		Category string `json:"category" binding:"required"`
		ImageURL string `json:"image_url" binding:"required"`
	}
	if err := c.ShouldBindJSON(&news); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	err := module.UpdateNews(c, id, news.Title, news.Content, news.AuthorID, news.Category, news.ImageURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"message": "News updated successfully"})
}

func DeleteNews(c *gin.Context) {
	module := DependencyInjection.NewNewsServiceProvider()
	id := c.Param("id")
	err := module.DeleteNews(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"message": "News deleted successfully"})
}
