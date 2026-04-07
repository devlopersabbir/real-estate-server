package property

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetProperties handles fetching all properties
//
//	@Summary		Get all properties
//	@Description	Fetches a list of all properties
//	@Tags			Properties
//	@Produce		json
//	@Success		200	{object}	map[string]string	"List of properties"
//	@Router			/api/v1/properties [get]
func GetProperties(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "List of properties"})
}

// GetProperty handles fetching a single property by ID
//
//	@Summary		Get a property
//	@Description	Fetches a single property by its ID
//	@Tags			Properties
//	@Produce		json
//	@Param			id	path		string				true	"Property ID"
//	@Success		200	{object}	map[string]string	"Property details"
//	@Router			/api/v1/properties/{id} [get]
func GetProperty(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "Get property", "id": id})
}

// CreateProperty handles creating a new property
//
//	@Summary		Create a property
//	@Description	Creates a new property
//	@Tags			Properties
//	@Accept			json
//	@Produce		json
//	@Success		201	{object}	map[string]string	"Property created successfully"
//	@Router			/api/v1/properties [post]
func CreateProperty(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "Property created successfully"})
}

// UpdateProperty handles updating an existing property
//
//	@Summary		Update a property
//	@Description	Updates an existing property by its ID
//	@Tags			Properties
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string				true	"Property ID"
//	@Success		200	{object}	map[string]string	"Property updated successfully"
//	@Router			/api/v1/properties/{id} [put]
func UpdateProperty(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "Property updated successfully", "id": id})
}

// DeleteProperty handles deleting a property
//
//	@Summary		Delete a property
//	@Description	Deletes a property by its ID
//	@Tags			Properties
//	@Param			id	path		string				true	"Property ID"
//	@Success		200	{object}	map[string]string	"Property deleted successfully"
//	@Router			/api/v1/properties/{id} [delete]
func DeleteProperty(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "Property deleted successfully", "id": id})
}
