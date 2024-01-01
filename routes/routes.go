package routes

import (
	"example.com/event-app-backend-go/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	auth := r.Group("/api/auth")
	auth.POST("/sign-in", SignIn)
	auth.POST("/sign-up", SignUp)

	me := r.Group("/api/me")
	me.Use(middlewares.Auth)
	me.GET("/", GetMyCredentials)
	me.PUT("/", UpdateMyCredentials)
	me.GET("/events", GetMyEvents)

	user := r.Group("/api/users")
	user.Use(middlewares.Auth)
	user.GET("/", GetUsers)
	user.GET("/:id", GetUserByID)
	user.POST("/", CreateUser)
	user.PUT("/:id", UpdateUser)
	user.DELETE("/:id", DeleteUser)

	category := r.Group("/api/categories")
	category.Use(middlewares.Auth)
	category.GET("/", GetCategories)
	category.GET("/:id", GetCategoryByID)
	category.POST("/", CreateCategory)
	category.PUT("/:id", UpdateCategory)
	category.DELETE("/:id", DeleteCategory)

	event := r.Group("/api/events")
	event.Use(middlewares.Auth)
	event.GET("/", GetEvents)
	event.GET("/:id", GetEventByID)
	event.POST("/", CreateEvent)
	event.PUT("/:id", UpdateEvent)
	event.DELETE("/:id", DeleteEvent)
	event.POST("/:id/attend", AttendEvent)
	event.DELETE("/:id/unattend", UnattendEvent)

}
