package property

import (
	"github.com/devlopersabbir/juan_don82-server/arch/middleware"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes attaches all property-domain routes to the provided versioned router group.
// Expected base group: /api/v1
func RegisterRoutes(v1 *gin.RouterGroup) {
	// /api/v1/properties
	properties := v1.Group("/v1/properties")
	{
		properties.GET("/", GetProperties)
		properties.GET("/:id", GetProperty)

		// Protected routes
		protected := properties.Group("")
		protected.Use(middleware.AuthMiddleware())
		{
			protected.POST("/", CreateProperty)
			protected.PUT("/:id", UpdateProperty)
			protected.DELETE("/:id", DeleteProperty)
		}
	}
}
