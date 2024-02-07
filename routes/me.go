package routes

import (
	"net/http"

	"example.com/event-app-backend-go/db"
	"example.com/event-app-backend-go/models"
	"example.com/event-app-backend-go/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetMyCredentials(c *gin.Context) {

	id, _ := uuid.Parse(c.GetString("userId"))

	var user models.User

	result := db.DB.First(&user, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, user)

}

func UpdateMyCredentials(c *gin.Context) {

	type UpdateMyCredentialsInput struct {
		Username string `form:"username"`
		Email    string `form:"email"`
		Password string `form:"password"`
	}

	var input UpdateMyCredentialsInput

	err := c.ShouldBind(&input)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	id, _ := uuid.Parse(c.GetString("userId"))

	var user models.User

	result := db.DB.First(&user, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	if input.Username != "" {
		user.Username = input.Username
	}

	if input.Email != "" {
		user.Email = input.Email
	}

	if input.Password != "" {
		hashedPassword := utils.HashPassword(input.Password)
		user.Password = hashedPassword
	}

	result = db.DB.Save(&user)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, user)

}

func GetMyEvents(c *gin.Context) {

	id, _ := uuid.Parse(c.GetString("userId"))

	var events []models.Event

	result := db.DB.Where("creator_id = ?", id).Find(&events)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, events)

}

func GetMyAttendedEvents(c *gin.Context) {

	id, _ := uuid.Parse(c.GetString("userId"))

	var user models.User

	result := db.DB.First(&user, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	var events []models.Event

	err := db.DB.Model(&user).Association("AttendedEvents").Find(&events)

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, events)

}

func GetMyComments(c *gin.Context) {

	id, _ := uuid.Parse(c.GetString("userId"))

	var user models.User

	result := db.DB.First(&user, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	var comments []models.Comment

	err := db.DB.Model(&user).Association("Comments").Find(&comments)

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, comments)

}

func GetMyLocations(c *gin.Context) {

	id, _ := uuid.Parse(c.GetString("userId"))

	var user models.User

	result := db.DB.First(&user, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	var locations []models.Location

	err := db.DB.Model(&user).Association("Locations").Find(&locations)

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": locations,
	})

}
