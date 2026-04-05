package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "User created successfully",
	})
}

func LoginUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "User logged in successfully",
	})
}

func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "User logged in successfully",
	})
}
