// Package routes routes/routes.go
package routes

import (
	"github.com/gin-gonic/gin"
	"v/backend/controllers"
	"v/backend/middleware"
)

func RegisterRoutes(router *gin.Engine) {

	// Instantiate controllers
	userController := controllers.NewUserController()
	meetingController := controllers.NewMeetingController()

	// User routes
	router.POST("api/signup", userController.Signup)
	router.POST("api/login", userController.Login)
	router.GET("api/validate", middleware.RequireAuth, userController.ValidateSession)
	router.POST("api/logout", controllers.Logout)

	// Meeting routes (using a route group for /meetings)
	meetingRoutes := router.Group("api/meetings")
	{
		meetingRoutes.POST("/", middleware.RequireAuth, meetingController.CreateMeeting)
		meetingRoutes.GET("/:id", middleware.RequireAuth, meetingController.GetMeeting)
		meetingRoutes.PUT("/:id", middleware.RequireAuth, meetingController.UpdateMeeting)
		meetingRoutes.DELETE("/:id", middleware.RequireAuth, meetingController.DeleteMeeting)
		meetingRoutes.GET("/", middleware.RequireAuth, meetingController.ListMeetings)
	}

}
