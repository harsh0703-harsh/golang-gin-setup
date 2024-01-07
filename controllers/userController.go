package controllers

import (
	"gin/initializers"
	"gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {

	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read body..",
		})

		return

	}

	user := models.User{Email: body.Email, Password: body.Password}

	result := initializers.DB.Create(&user)

	if result.Error != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to create the user record",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
