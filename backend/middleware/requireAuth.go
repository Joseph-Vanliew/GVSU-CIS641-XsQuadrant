package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	"os"
	"time"
	"v/backend/initializers"
	"v/backend/models"
)

func RequireAuth(c *gin.Context) {
	//Get the cookie off request
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	//Decode/validate it
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		log.Println("Token parsing/validation error:", err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		//check the expiration
		exp, expOk := claims["exp"].(float64)
		if !expOk || exp == 0 {
			log.Println("Missing or invalid 'exp' claim in token")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if float64(time.Now().Unix()) > exp {
			log.Println("Token has expired")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// find the user with token sub
		var user models.User
		initializers.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			log.Println("User not found for token 'sub' claim")
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		//attach request
		c.Set("user", user)

		//continue
		c.Next()
	} else {
		log.Println("Token does not contain 'sub' claim")
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
