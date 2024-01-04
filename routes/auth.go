package routes

import (
	"net/http"

	"example.com/event-app-backend-go/db"
	"example.com/event-app-backend-go/models"
	"example.com/event-app-backend-go/utils"
	"github.com/gin-gonic/gin"
)

func SignIn(c *gin.Context) {

	type SignInInput struct {
		Email    string `form:"email" binding:"required"`
		Password string `form:"password" binding:"required"`
	}

	var input SignInInput

	err := c.ShouldBind(&input)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	var user models.User

	result := db.DB.Where("email = ?", input.Email).First(&user)

	if result.Error == nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	comparePasswords := utils.ComparePasswords(user.Password, input.Password)

	if !comparePasswords {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid credentials",
		})
		return
	}

	var token string

	token, err = utils.GenerateToken(user.ID, user.Username)

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}

func SignUp(c *gin.Context) {

	type SignUpInput struct {
		Username string `form:"username" binding:"required"`
		Email    string `form:"email" binding:"required"`
		Password string `form:"password" binding:"required"`
	}

	var input SignUpInput

	err := c.ShouldBind(&input)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	var user models.User

	result := db.DB.Where("email = ? OR username = ?", input.Email, input.Username).First(&user)

	if result.Error == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User already exists",
		})
		return
	}

	user.Username = input.Username
	user.Email = input.Email
	hashedPassword := utils.HashPassword(input.Password)
	user.Password = hashedPassword

	result = db.DB.Create(&user)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	var token string

	token, err = utils.GenerateToken(user.ID, user.Username)

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}
