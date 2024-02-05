package routes

import (
	"net/http"

	"example.com/event-app-backend-go/db"
	"example.com/event-app-backend-go/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetEvents(c *gin.Context) {

	var events []models.Event

	result := db.DB.Find(&events)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": events,
	})
}

func GetEventByID(c *gin.Context) {

	var event models.Event

	id := c.Param("id")

	result := db.DB.First(&event, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, event)

}

func CreateEvent(c *gin.Context) {

	type CreateEventInput struct {
		Name        string `form:"name" binding:"required"`
		Description string `form:"description" binding:"required"`
		Capacity    uint   `form:"capacity" binding:"required"`
	}

	var input CreateEventInput

	err := c.ShouldBind(&input)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	userId, _ := uuid.Parse(c.GetString("userId"))

	event := models.Event{
		Name:        input.Name,
		Description: input.Description,
		Capacity:    input.Capacity,
		CreatorID:   userId,
	}

	result := db.DB.Create(&event)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, event)

}

func UpdateEvent(c *gin.Context) {

	type UpdateEvent struct {
		Name        string `form:"name"`
		Description string `form:"description"`
		Capacity    uint   `form:"capacity"`
	}

	var input UpdateEvent

	err := c.ShouldBind(&input)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	id := c.Param("id")

	var event models.Event

	result := db.DB.First(&event, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	if event.Name != "" {
		event.Name = input.Name
	}

	if event.Description != "" {
		event.Description = input.Description
	}

	if event.Capacity != 0 {
		event.Capacity = input.Capacity
	}

	result = db.DB.Save(&event)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, event)
}

func DeleteEvent(c *gin.Context) {

	id := c.Param("id")

	var event models.Event

	result := db.DB.First(&event, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	result = db.DB.Delete(&event)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, event)

}

func AttendEvent(c *gin.Context) {

	id := c.Param("id")

	var event models.Event

	result := db.DB.First(&event, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	userId := c.GetUint("userId")

	var user models.User

	result = db.DB.First(&user, userId)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	err := db.DB.Model(&event).Association("Participants").Append(&user)

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, event)

}

func UnattendEvent(c *gin.Context) {

	id := c.Param("id")

	var event models.Event

	result := db.DB.First(&event, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	userId := c.GetUint("userId")

	var user models.User

	result = db.DB.First(&user, userId)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	err := db.DB.Model(&event).Association("Participants").Delete(&user)

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, event)

}

func GetEventParticipants(c *gin.Context) {

	id := c.Param("id")

	var event models.Event

	result := db.DB.First(&event, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	var participants []models.User

	err := db.DB.Model(&event).Association("Participants").Find(&participants)

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, participants)

}

func GetEventComments(c *gin.Context) {

	id := c.Param("id")

	var event models.Event

	result := db.DB.First(&event, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	var comments []models.Comment

	err := db.DB.Model(&event).Association("Comments").Find(&comments)

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, comments)

}
