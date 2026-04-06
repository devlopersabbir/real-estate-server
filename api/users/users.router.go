package users

import (
	"github.com/devlopersabbir/juan_don82-server/arch/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(v1 *gin.RouterGroup) {
	// /api/v1/auth
	auth := v1.Group("/v1/auth")
	{
		auth.POST("/login", LoginUser)
		auth.POST("/register", CreateUser)
		auth.POST("/refresh", RefreshUserToken)
	}

	// /api/v1/users
	users := v1.Group("/v1/users")
	users.Use(middleware.AuthMiddleware())
	{
		users.GET("/", GetUsers)
	}
}
