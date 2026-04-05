package users

import (
	"github.com/gin-gonic/gin"
)

// RegisterRoutes attaches all user-domain routes to the provided versioned router group.
// Expected base group: /api/v1
func RegisterRoutes(v1 *gin.RouterGroup) {
	// /api/v1/users
	users := v1.Group("/v1/users")
	{
		users.GET("/", GetUsers)
		users.POST("/register", CreateUser)
	}

	// /api/v1/auth
	auth := v1.Group("/v1/auth")
	{
		auth.POST("/login", LoginUser)
	}
}
