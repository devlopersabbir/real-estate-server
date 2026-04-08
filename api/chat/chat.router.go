package chat

import (
	"github.com/devlopersabbir/juan_don82-server/arch/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(v1 *gin.RouterGroup) {
	chats := v1.Group("/v1/chats")
	chats.Use(middleware.AuthMiddleware())
	{
		chats.POST("/start", StartChat)
		chats.POST("/message", SendMsg)
		chats.GET("/", GetRooms)
		chats.GET("/:id/messages", GetRoomMessages)
	}
}
