package users

import (
	"net/http"

	"github.com/devlopersabbir/juan_don82-server/api/users/domain"
	v "github.com/devlopersabbir/juan_don82-server/internal/pkg/validator"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var body domain.CreateUserRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	// Validate struct fields
	if errs := v.Validate(body); errs != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"errors": errs,
		})
		return
	}

	// TODO: pass body to repository / service layer
	c.JSON(http.StatusCreated, gin.H{
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
		"message": "Users fetched successfully",
	})
}
