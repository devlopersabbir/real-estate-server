package validator

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

// singleton instance – created once, reused everywhere
var validate = validator.New()

// ValidationError holds a human-readable field-level error.
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// Validate parses the struct tags on T and returns a slice of ValidationErrors.
// Returns nil when the struct is valid.
//
// Usage:
//
//	errs := validator.Validate(myStruct)
//	if errs != nil {
//	    c.JSON(http.StatusBadRequest, gin.H{"errors": errs})
//	    return
//	}
func Validate[T any](payload T) []ValidationError {
	err := validate.Struct(payload)
	if err == nil {
		return nil
	}

	var errs []ValidationError
	for _, e := range err.(validator.ValidationErrors) {
		errs = append(errs, ValidationError{
			Field:   strings.ToLower(e.Field()),
			Message: fieldMessage(e),
		})
	}
	return errs
}

// fieldMessage converts a validator.FieldError into a readable sentence.
func fieldMessage(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", e.Field())
	case "email":
		return fmt.Sprintf("%s must be a valid email address", e.Field())
	case "min":
		return fmt.Sprintf("%s must be at least %s characters long", e.Field(), e.Param())
	case "max":
		return fmt.Sprintf("%s must be at most %s characters long", e.Field(), e.Param())
	case "gte":
		return fmt.Sprintf("%s must be greater than or equal to %s", e.Field(), e.Param())
	case "lte":
		return fmt.Sprintf("%s must be less than or equal to %s", e.Field(), e.Param())
	default:
		return fmt.Sprintf("%s failed validation: %s", e.Field(), e.Tag())
	}
}
