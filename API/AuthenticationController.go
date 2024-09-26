package API

import (
	"net/http"
	"th3y3m/e-commerce-platform/DependencyInjection"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type RegisterRequest struct {
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirmPassword" binding:"required"`
}

type RegisterResponse struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

// @Summary User login
// @Description Logs in a user by email and password.
// @Tags auth
// @Accept json
// @Produce json
// @Param user body API.LoginRequest true "User credentials"
// @Success 200 {object} API.LoginResponse
// @Failure 400 {object} ErrorResponse
// @Router /login [post]
func Login(c *gin.Context) {
	var request LoginRequest

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email and password are required"})
		return
	}
	service := DependencyInjection.NewNewAuthenticationServiceProvider()

	token, err := service.Login(request.Email, request.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// @Summary Register a new customer
// @Description Registers a new customer by providing user details.
// @Tags auth
// @Accept json
// @Produce json
// @Param customer body API.RegisterRequest true "Customer details"
// @Success 201 {object} API.RegisterResponse
// @Failure 400 {object} ErrorResponse
// @Router /register [post]
func RegisterCustomer(c *gin.Context) {
	var request RegisterRequest

	if request.Password != request.ConfirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{"error": "passwords do not match"})
		return
	}

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email and password are required"})
		return
	}
	service := DependencyInjection.NewNewAuthenticationServiceProvider()

	err := service.RegisterCustomer(request.Email, request.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}
