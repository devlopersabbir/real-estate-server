package wishlist

import (
	"github.com/devlopersabbir/juan_don82-server/arch/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(v1 *gin.RouterGroup) {
	wishlist := v1.Group("/v1/wishlist")
	wishlist.Use(middleware.AuthMiddleware())
	{
		wishlist.GET("/", List)
		wishlist.POST("/", Add)
		wishlist.DELETE("/", Remove)
	}
}
