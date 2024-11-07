// Package routes routes/routes.go
package routes

import (
	"github.com/gin-gonic/gin"
	"xsface/controllers"
	"xsface/middleware"
)

func RegisterRoutes(router *gin.Engine) {

	userController := controllers.NewUserController()
	meetingController := controllers.NewMeetingController()

	// Register routes
	router.POST("/signup", userController.Signup)
	router.POST("/login", userController.Login)
	router.GET("/validate", middleware.RequireAuth, userController.Validate)
	// Add additional route groups or endpoints as needed
}
