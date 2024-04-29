package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tjalle/link_shortener/helpers"
)

type LoginRequestBody struct {
	Username string
	Password string
}

func Login(c *gin.Context) {
	var requestBody LoginRequestBody

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized.",
		})
		return
	}

	// TODO: Login Verification
	hasCorrectCredentials := requestBody.Username == "test" && requestBody.Password == "1234"

	if !hasCorrectCredentials {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized.",
		})
		return
	}

	tokenString, err := helpers.CreateToken(requestBody.Username)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
