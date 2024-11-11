// Package routes routes/routes.go
package routes

import (
	"github.com/gin-gonic/gin"
	"xsface/controllers"
	"xsface/middleware"
)

func RegisterRoutes(router *gin.Engine) {

	// Instantiate controllers
	userController := controllers.NewUserController()
	meetingController := controllers.NewMeetingController()

	// User routes
	router.POST("/signup", userController.Signup)
	router.POST("/login", userController.Login)
	router.GET("/validate", middleware.RequireAuth, userController.Validate)

	// Meeting routes (using a route group for /meetings)
	meetingRoutes := router.Group("/meetings")
	{
		meetingRoutes.POST("/", middleware.RequireAuth, meetingController.CreateMeeting)
		meetingRoutes.GET("/:id", middleware.RequireAuth, meetingController.GetMeeting)
		meetingRoutes.PUT("/:id", middleware.RequireAuth, meetingController.UpdateMeeting)
		meetingRoutes.DELETE("/:id", middleware.RequireAuth, meetingController.DeleteMeeting)
		meetingRoutes.GET("/", middleware.RequireAuth, meetingController.ListMeetings)
	}

}
