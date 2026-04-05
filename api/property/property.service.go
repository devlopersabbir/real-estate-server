package property

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProperties(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "List of properties"})
}

func GetProperty(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "Get property", "id": id})
}

func CreateProperty(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "Property created successfully"})
}

func UpdateProperty(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "Property updated successfully", "id": id})
}

func DeleteProperty(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "Property deleted successfully", "id": id})
}
