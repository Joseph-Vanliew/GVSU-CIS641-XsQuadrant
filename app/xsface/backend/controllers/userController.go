package controllers

import (
	"log"
	"net/http"
	"os"
	"time"
	"xsface/initializers"
	"xsface/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
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

func Login(c *gin.Context) {
	// Get the email and pass of req body
	var body struct {
		Email    string
		Password string
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	var user models.User
	//Look up requested user
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email",
		})
		return
	}

	//Compare sent in pass with saved user pass hash
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid password",
		})
		return
	}

	//generate a jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	//sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid to create token",
		})
		return
	}

	//send back
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{})

}

func Validate(c *gin.Context) {

	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}
