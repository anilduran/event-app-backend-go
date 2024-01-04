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
	me.GET("/events/attended", GetMyAttendedEvents)
	me.GET("/comments", GetMyComments)
	me.GET("/locations", GetMyLocations)

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
	category.GET("/:id/events", GetCategoryEvents)

	event := r.Group("/api/events")
	event.Use(middlewares.Auth)
	event.GET("/", GetEvents)
	event.GET("/:id", GetEventByID)
	event.POST("/", CreateEvent)
	event.PUT("/:id", UpdateEvent)
	event.DELETE("/:id", DeleteEvent)
	event.GET("/:id/participants", GetEventParticipants)
	event.GET("/:id/comments", GetEventComments)
	event.POST("/:id/attend", AttendEvent)
	event.DELETE("/:id/unattend", UnattendEvent)

	comment := r.Group("/api/comments")
	comment.Use(middlewares.Auth)
	comment.GET("/", GetComments)
	comment.GET("/:id", GetCommentByID)
	comment.POST("/", CreateComment)
	comment.PUT("/:id", UpdateComment)
	comment.DELETE("/:id", DeleteComment)

	location := r.Group("/api/locations")
	location.Use(middlewares.Auth)
	location.GET("/", GetLocations)
	location.GET("/:id", GetLocationByID)
	location.POST("/", CreateLocation)
	location.PUT("/:id", UpdateLocation)
	location.DELETE("/:id", DeleteLocation)
	location.GET("/:id/events", GetLocationEvents)

}
