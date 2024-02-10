package routes

import (
	"net/http"

	"example.com/event-app-backend-go/db"
	"example.com/event-app-backend-go/models"
	"example.com/event-app-backend-go/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetUsers(c *gin.Context) {

	var users []models.User

	result := db.DB.Find(&users)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})

}

func GetUserByID(c *gin.Context) {

	var user models.User

	id, _ := uuid.Parse(c.Param("id"))

	result := db.DB.First(&user, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, user)

}

func CreateUser(c *gin.Context) {

	type CreateUserInput struct {
		Username        string `form:"username" binding:"required"`
		Email           string `form:"email" binding:"required"`
		Password        string `form:"password" binding:"required"`
		ProfilePhotoUrl string `form:"profilePhotoUrl"`
	}

	var input CreateUserInput

	err := c.ShouldBind(&input)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	hashedPassword := utils.HashPassword(input.Password)

	user := models.User{
		Username: input.Username,
		Email:    input.Email,
		Password: hashedPassword,
	}

	if input.ProfilePhotoUrl != "" {
		user.ProfilePhotoUrl = input.ProfilePhotoUrl
	}

	result := db.DB.Create(&user)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, user)

}

func UpdateUser(c *gin.Context) {

	type UpdateUserInput struct {
		Username        string `form:"username"`
		Email           string `form:"email"`
		Password        string `form:"password"`
		ProfilePhotoUrl string `form:"profilePhotoUrl"`
	}

	var input UpdateUserInput

	err := c.ShouldBind(&input)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	id, _ := uuid.Parse(c.Param("id"))

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

	if input.ProfilePhotoUrl != "" {
		user.ProfilePhotoUrl = input.ProfilePhotoUrl
	}

	result = db.DB.Save(&user)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, user)

}

func DeleteUser(c *gin.Context) {

	id, _ := uuid.Parse(c.Param("id"))

	var user models.User

	result := db.DB.First(&user, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	result = db.DB.Delete(&user)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, user)

}
