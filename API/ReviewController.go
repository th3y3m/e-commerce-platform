package API

import (
	"net/http"
	"strconv"
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Services"

	"github.com/gin-gonic/gin"
)

func GetPaginatedReviewList(c *gin.Context) {
	sortBy := c.DefaultQuery("sortBy", "")
	reviewID := c.DefaultQuery("reviewID", "")
	userID := c.DefaultQuery("userID", "")
	productID := c.DefaultQuery("productID", "")
	pageIndex, _ := strconv.Atoi(c.DefaultQuery("pageIndex", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	minRating, _ := strconv.Atoi(c.DefaultQuery("minRating", "0"))
	maxRating, _ := strconv.Atoi(c.DefaultQuery("maxRating", "5"))
	status, _ := strconv.ParseBool(c.DefaultQuery("status", ""))

	reviews, err := Services.GetPaginatedReviewList(sortBy, reviewID, userID, productID, pageIndex, pageSize, &minRating, &maxRating, &status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, reviews)
}

func GetAllReviews(c *gin.Context) {
	reviews, err := Services.GetAllReviews()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, reviews)
}

func GetReviewByID(c *gin.Context) {

	id := c.Param("id")
	review, err := Services.GetReviewByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, review)
}

func CreateReview(c *gin.Context) {
	var review BusinessObjects.NewReview
	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := Services.CreateReview(review)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Review created successfully"})
}

func UpdateReview(c *gin.Context) {
	id := c.Param("id")

	var review struct {
		Comment string `json:"comment" binding:"required"`
		Rating  int    `json:"rating" binding:"required"`
	}
	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := Services.UpdateReview(id, review.Comment, review.Rating)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Review updated successfully"})
}

func DeleteReview(c *gin.Context) {
	id := c.Param("id")
	err := Services.DeleteReview(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Review deleted successfully"})
}
