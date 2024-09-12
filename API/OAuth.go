package API

import (
	"log"
	"net/http"
	"os"
	"th3y3m/e-commerce-platform/Services"

	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/facebook"
	"github.com/markbates/goth/providers/google"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Set up environment variables
	clientID := os.Getenv("GOOGLE_CLIENT_ID")
	clientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	facebookID, facebookSecret := os.Getenv("FACEBOOK_CLIENT_ID"), os.Getenv("FACEBOOK_CLIENT_SECRET")

	sessionSecret := os.Getenv("SESSION_SECRET")

	// Check if session secret is set
	if sessionSecret == "" {
		log.Fatal("SESSION_SECRET environment variable is not set")
	}

	// Set SESSION_SECRET for Gothic
	os.Setenv("SESSION_SECRET", sessionSecret)

	// Initialize Goth with the Google provider
	goth.UseProviders(
		google.New(clientID, clientSecret, "http://localhost:8080/auth/google/callback"),
		facebook.New(facebookID, facebookSecret, "http://localhost:8080/auth/facebook/callback"),
	)

	// Explicitly set Gothic store to use the session store from Gin
	key := []byte(sessionSecret)
	store := cookie.NewStore(key, key)
	gothic.Store = store
}

// GoogleLogin redirects to Google for authentication
func GoogleLogin(c *gin.Context) {
	c.Request.URL.RawQuery = "provider=google"
	gothic.BeginAuthHandler(c.Writer, c.Request)
}

// GoogleCallback handles Google OAuth callback
func GoogleCallback(c *gin.Context) {
	// Complete the user authentication with Gothic
	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		log.Printf("Error during authentication: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to authenticate with Google"})
		return
	}
	// Handle Google user and generate JWT token
	token, err := Services.HandleOAuthUser(user)
	if err != nil {
		log.Printf("Error handling Google user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to handle Google user"})
		return
	}

	// Respond with the generated JWT token
	c.JSON(http.StatusOK, gin.H{"token": token})
}

// GoogleLogout clears the session
func GoogleLogout(c *gin.Context) {
	if err := gothic.Logout(c.Writer, c.Request); err != nil {
		log.Printf("Error logging out: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to logout"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}

// FacebookLogin redirects to Facebook for authentication
func FacebookLogin(c *gin.Context) {
	c.Request.URL.RawQuery = "provider=facebook"
	gothic.BeginAuthHandler(c.Writer, c.Request)
}

// FacebookCallback handles Facebook OAuth callback
func FacebookCallback(c *gin.Context) {
	// Complete the user authentication with Gothic
	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		log.Printf("Error during authentication: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to authenticate with Facebook"})
		return
	}

	// Handle Facebook user and generate JWT token
	token, err := Services.HandleOAuthUser(user) // This can be renamed to a more generic handler
	if err != nil {
		log.Printf("Error handling Facebook user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to handle Facebook user"})
		return
	}

	// Respond with the generated JWT token
	c.JSON(http.StatusOK, gin.H{"token": token})
}

// FacebookLogout clears the session
func FacebookLogout(c *gin.Context) {
	if err := gothic.Logout(c.Writer, c.Request); err != nil {
		log.Printf("Error logging out: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to logout"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}
