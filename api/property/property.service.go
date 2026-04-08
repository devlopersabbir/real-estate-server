package property

import (
	"fmt"
	"strconv"

	"github.com/devlopersabbir/juan_don82-server/api/property/core"
	"github.com/devlopersabbir/juan_don82-server/api/property/domain"
	"github.com/devlopersabbir/juan_don82-server/api/subscriptions"
	"github.com/devlopersabbir/juan_don82-server/arch/networks"
	v "github.com/devlopersabbir/juan_don82-server/internal/pkg/validator"
	"github.com/gin-gonic/gin"
)

// GetProperties handles fetching all properties
//
//	@Summary		Get all properties
//	@Description	Fetches a list of all properties
//	@Tags			Properties
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}	"List of properties"
//	@Router			/api/v1/properties [get]
func GetProperties(c *gin.Context) {
	res := networks.Send(c)

	var filter PropertyFilter
	if err := c.ShouldBindQuery(&filter); err != nil {
		res.BadRequestError("Invalid filter parameters", err)
		return
	}

	properties, err := SearchPropertiesElastic(c, filter)
	if err != nil {
		res.InternalServerError("Failed to fetch properties from search index", err)
		return
	}
	res.SuccessDataResponse("List of properties from search index", properties)
}

// GetProperty handles fetching a single property by ID
//
//	@Summary		Get a property
//	@Description	Fetches a single property by its ID
//	@Tags			Properties
//	@Produce		json
//	@Param			id	path		int						true	"Property ID"
//	@Success		200	{object}	map[string]interface{}	"Property details"
//	@Router			/api/v1/properties/{id} [get]
func GetProperty(c *gin.Context) {
	res := networks.Send(c)
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		res.BadRequestError("Invalid property ID", err)
		return
	}

	property, err := FindByID(id)
	if err != nil {
		res.NotFoundError("Property not found", err)
		return
	}

	res.SuccessDataResponse("Property details", property)
}

// CreateProperty handles creating a new property
//
//	@Summary		Create a property
//	@Description	Creates a new property (Only for AGENT and SYSTEM_ADMIN)
//	@Tags			Properties
//	@Security		BearerAuth
//	@Accept			json
//	@Produce		json
//	@Param			body	body		domain.PropertyRequest	true	"Property Details"
//	@Success		201		{object}	map[string]interface{}	"Property created successfully"
//	@Router			/api/v1/properties [post]
func CreateProperty(c *gin.Context) {
	res := networks.Send(c)

	role, exists := c.Get("role")
	fmt.Println(role)
	if !exists || role == "user" || role == "USER" {
		res.UnauthorizedError("You are not authorized to create properties", nil)
		return
	}

	userIDVal, exists := c.Get("userID")
	if !exists {
		res.UnauthorizedError("User ID not found in context", nil)
		return
	}
	userID := userIDVal.(uint)
	if role == "agent" || role == "AGENT" {
		_, err := subscriptions.FindActiveSubscription(userID)
		if err != nil {
			res.ForbiddenError("Active subscription required to list properties", err)
			return
		}
	}

	var body domain.PropertyRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		res.BadRequestError("Invalid request body", err)
		return
	}

	if errs := v.Validate(body); errs != nil {
		res.ValidationError("Validation failed", errs)
		return
	}

	property := &core.Property{
		UserID:        userID,
		Name:          body.Name,
		Description:   body.Description,
		Address:       body.Address,
		City:          body.City,
		State:         body.State,
		ZipCode:       body.ZipCode,
		Country:       body.Country,
		Price:         body.Price,
		DiscountPrice: body.DiscountPrice,
		Bedrooms:      body.Bedrooms,
		Bathrooms:     body.Bathrooms,
		SquareFeet:    body.SquareFeet,
		PropertyType:  body.PropertyType,
		RentPeriod:    body.RentPeriod,
		Status:        body.Status,
		Images:        body.Images,
		Videos:        body.Videos,
	}

	if err := Store(property); err != nil {
		res.InternalServerError("Failed to create property", err)
		return
	}

	if err := StoreElastic(c, property); err != nil {
		res.InternalServerError("Failed to store property in search index", err)
		return
	}

	res.SuccessDataResponse("Property created successfully", property)
}

// UpdateProperty handles updating an existing property
//
//	@Summary		Update a property
//	@Description	Updates an existing property by its ID
//	@Tags			Properties
//	@Security		BearerAuth
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int						true	"Property ID"
//	@Param			body	body		domain.PropertyRequest	true	"Property Details"
//	@Success		200		{object}	map[string]interface{}	"Property updated successfully"
//	@Router			/api/v1/properties/{id} [put]
func UpdateProperty(c *gin.Context) {
	res := networks.Send(c)

	role, exists := c.Get("role")
	if !exists || role == "user" || role == "USER" {
		res.UnauthorizedError("You are not authorized to update properties", nil)
		return
	}

	userIDVal, _ := c.Get("userID")
	userID := uint(userIDVal.(float64))

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		res.BadRequestError("Invalid property ID", err)
		return
	}

	property, err := FindByID(id)
	if err != nil {
		res.NotFoundError("Property not found", err)
		return
	}

	// Check if owner or admin
	if property.UserID != userID && role != "system_admin" && role != "SYSTEM_ADMIN" {
		res.UnauthorizedError("You do not own this property", nil)
		return
	}

	var body domain.PropertyRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		res.BadRequestError("Invalid request body", err)
		return
	}

	if errs := v.Validate(body); errs != nil {
		res.ValidationError("Validation failed", errs)
		return
	}

	property.Name = body.Name
	property.Description = body.Description
	property.Address = body.Address
	property.City = body.City
	property.State = body.State
	property.ZipCode = body.ZipCode
	property.Country = body.Country
	property.Price = body.Price
	property.DiscountPrice = body.DiscountPrice
	property.Bedrooms = body.Bedrooms
	property.Bathrooms = body.Bathrooms
	property.SquareFeet = body.SquareFeet
	property.PropertyType = body.PropertyType
	property.RentPeriod = body.RentPeriod
	property.Status = body.Status
	property.Images = body.Images
	property.Videos = body.Videos

	if err := Update(property); err != nil {
		res.InternalServerError("Failed to update property", err)
		return
	}

	if err := UpdateElastic(c, strconv.Itoa(int(property.ID)), property); err != nil {
		res.InternalServerError("Failed to update property in search index", err)
		return
	}

	res.SuccessDataResponse("Property updated successfully", property)
}

// DeleteProperty handles deleting a property
//
//	@Summary		Delete a property
//	@Description	Deletes a property by its ID
//	@Tags			Properties
//	@Security		BearerAuth
//	@Param			id	path		int						true	"Property ID"
//	@Success		200	{object}	map[string]interface{}	"Property deleted successfully"
//	@Router			/api/v1/properties/{id} [delete]
func DeleteProperty(c *gin.Context) {
	res := networks.Send(c)

	role, exists := c.Get("role")
	if !exists || role == "user" || role == "USER" {
		res.UnauthorizedError("You are not authorized to delete properties", nil)
		return
	}

	userIDVal, _ := c.Get("userID")
	userID := uint(userIDVal.(float64))

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		res.BadRequestError("Invalid property ID", err)
		return
	}

	property, err := FindByID(id)
	if err != nil {
		res.NotFoundError("Property not found", err)
		return
	}

	if property.UserID != userID && role != "system_admin" && role != "SYSTEM_ADMIN" {
		res.UnauthorizedError("You do not own this property", nil)
		return
	}

	if err := Delete(id); err != nil {
		res.InternalServerError("Failed to delete property", err)
		return
	}

	if err := DeleteElastic(c, strconv.Itoa(id)); err != nil {
		res.InternalServerError("Failed to delete property from search index", err)
		return
	}

	res.SuccessMsgResponse("Property deleted successfully")
}
