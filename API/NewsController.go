package API

import (
	"net/http"
	"strconv"
	"th3y3m/e-commerce-platform/Services"

	"github.com/gin-gonic/gin"
)

func GetPaginatedNewsList(c *gin.Context) {
	searchValue := c.DefaultQuery("searchValue", "")
	sortBy := c.DefaultQuery("sortBy", "")
	pageIndex, _ := strconv.Atoi(c.DefaultQuery("pageIndex", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	newID := c.DefaultQuery("newID", "")
	authorID := c.DefaultQuery("authorID", "")
	status, _ := strconv.ParseBool(c.DefaultQuery("status", ""))

	news, err := Services.GetPaginatedNewsList(searchValue, sortBy, newID, authorID, pageIndex, pageSize, &status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"news": news})
}

func GetAllNews(c *gin.Context) {
	news, err := Services.GetAllNews()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"news": news})
}

func GetNewsByID(c *gin.Context) {

	id := c.Param("id")
	news, err := Services.GetNewsByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"news": news})
}

func CreateNews(c *gin.Context) {
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

	err := Services.CreateNews(news.Title, news.Content, news.AuthorID, news.Category, news.ImageURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

func UpdateNews(c *gin.Context) {
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
	err := Services.UpdateNews(id, news.Title, news.Content, news.AuthorID, news.Category, news.ImageURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

func DeleteNews(c *gin.Context) {
	id := c.Param("id")
	err := Services.DeleteNews(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
