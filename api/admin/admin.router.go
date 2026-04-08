package admin

import (
	"github.com/devlopersabbir/juan_don82-server/arch/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(v1 *gin.RouterGroup) {
	admin := v1.Group("/v1/admin")
	admin.Use(middleware.AuthMiddleware())
	// Additional Admin check middleware should be added here
	{
		admin.GET("/users", ManageUsers)
		admin.GET("/properties", ManageProperties)
		admin.GET("/chats", ViewAllChats)
	}
}
