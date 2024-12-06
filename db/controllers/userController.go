package controllers

import (
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"time"
	"v/db/initializers"
	"v/db/middleware"
	"v/db/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	db *gorm.DB
}

func NewUserController() *UserController {
	return &UserController{
		db: initializers.DB, // Use the global DB connection from initializers
	}
}

// Signup creates a new user which is then stored in the db
func (uc *UserController) Signup(c *gin.Context) {
	//Get the email/pass off request body
	var body struct {
		Email     string
		Password  string
		Firstname string
		Lastname  string
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	//Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	// Create the user model
	user := models.User{
		Email:     body.Email,
		Password:  string(hash),
		Firstname: body.Firstname,
		Lastname:  body.Lastname,
	}

	// Insert the user into the database
	result := initializers.DB.Create(&user)
	if result.Error != nil {
		log.Printf("Error creating user in database: %v", result.Error)

		// Additional error handling for database-specific errors
		if result.Error.Error() == "duplicate key value violates unique constraint" {
			c.JSON(http.StatusConflict, gin.H{
				"error": "Email already exists",
			})
		} else {
			log.Printf("Error creating user in database: %v", result.Error)
			log.Printf("GORM Error Details: RowsAffected=%d, Error=%v", result.RowsAffected, result.Error)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to create user",
			})
		}
		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{})
}

// Login logs in the user and creates a JWT for session management
func (uc *UserController) Login(c *gin.Context) {
	// Get the email and password from the request body
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Bind JSON body to the struct
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	var user models.User
	// Look up requested user
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email",
		})
		return
	}

	// Compare sent password with saved user password hash
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid password",
		})
		return
	}

	// Generate a JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 2).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})
		return
	}

	// Set the cookie with the token
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
	})
}

// Logout Invalidates the JWT for user session
func Logout(c *gin.Context) {
	// Clear the cookie
	c.SetCookie("Authorization", "", -1, "/", "", false, true)

	// Respond with success
	c.JSON(http.StatusOK, gin.H{
		"message": "Logged out successfully",
	})
}

// ValidateSession returns the user info for the session, always use requireAuth in middleware when validating
func (uc *UserController) ValidateSession(c *gin.Context) {
	middleware.RequireAuth(c)

	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
