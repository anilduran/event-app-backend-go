package routes

import (
	"net/http"

	"example.com/event-app-backend-go/db"
	"example.com/event-app-backend-go/models"
	"github.com/gin-gonic/gin"
)

func GetLocations(c *gin.Context) {

	var locations []models.Location

	result := db.DB.Find(&locations)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": locations,
	})

}

func GetLocationByID(c *gin.Context) {

	var location models.Location

	id := c.Param("id")

	result := db.DB.First(&location, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, location)
}

func CreateLocation(c *gin.Context) {

	type CreateLocationInput struct {
		Name        string  `form:"name" binding:"required"`
		Description string  `form:"description" binding:"required"`
		ImageUrl    string  `form:"image_url" binding:"required"`
		Latitude    float64 `form:"latitude" binding:"required"`
		Longitude   float64 `form:"longitude" binding:"required"`
	}

	var input CreateLocationInput

	err := c.ShouldBind(&input)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	userId := c.GetUint("userId")

	location := models.Location{
		Name:        input.Name,
		Description: input.Description,
		ImageUrl:    input.ImageUrl,
		Latitude:    input.Latitude,
		Longitude:   input.Longitude,
		CreatorID:   userId,
	}

	result := db.DB.Create(&location)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, location)
}

func UpdateLocation(c *gin.Context) {

	type UpdateLocationInput struct {
		Name        string  `form:"name"`
		Description string  `form:"description"`
		ImageUrl    string  `form:"image_url"`
		Latitude    float64 `form:"latitude"`
		Longitude   float64 `form:"longitude"`
	}

	var input UpdateLocationInput

	err := c.ShouldBind(&input)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	var location models.Location

	id := c.Param("id")

	result := db.DB.First(&location, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	if input.Name != "" {
		location.Name = input.Name
	}

	if input.Description != "" {
		location.Description = input.Description
	}

	if input.ImageUrl != "" {
		location.ImageUrl = input.ImageUrl
	}

	if input.Latitude != 0 {
		location.Latitude = input.Latitude
	}

	if input.Longitude != 0 {
		location.Longitude = input.Longitude
	}

	result = db.DB.Save(&location)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, location)

}

func DeleteLocation(c *gin.Context) {

	id := c.Param("id")

	var location models.Location

	result := db.DB.First(&location, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	result = db.DB.Delete(&location)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, location)

}

func GetLocationEvents(c *gin.Context) {

	id := c.Param("id")

	var location models.Location

	result := db.DB.First(&location, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	var events []models.Event

	err := db.DB.Model(&location).Association("Events").Find(&events)

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": events,
	})

}
