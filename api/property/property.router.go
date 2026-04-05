package property

import (
	"github.com/gin-gonic/gin"
)

// RegisterRoutes attaches all property-domain routes to the provided versioned router group.
// Expected base group: /api/v1
func RegisterRoutes(v1 *gin.RouterGroup) {
	// /api/v1/properties
	properties := v1.Group("/properties")
	{
		properties.GET("/", GetProperties)
		properties.GET("/:id", GetProperty)
		properties.POST("/", CreateProperty)
		properties.PUT("/:id", UpdateProperty)
		properties.DELETE("/:id", DeleteProperty)
	}
}
