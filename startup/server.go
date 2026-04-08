package startup

import (
	"net/http"

	"github.com/devlopersabbir/juan_don82-server/api/admin"
	"github.com/devlopersabbir/juan_don82-server/api/chat"
	"github.com/devlopersabbir/juan_don82-server/api/property"
	"github.com/devlopersabbir/juan_don82-server/api/subscriptions"
	"github.com/devlopersabbir/juan_don82-server/api/users"
	"github.com/devlopersabbir/juan_don82-server/api/wishlist"
	_ "github.com/devlopersabbir/juan_don82-server/docs"
	"github.com/devlopersabbir/juan_don82-server/internal/pkg/config"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

/*
@title Juan Don API
@version 1.0
@description Modern REST API for Juan Don project
@host localhost:9000
@BasePath /

@securityDefinitions.apikey BearerAuth
@in header
@name Authorization
*/
func Server(*config.Env) *gin.Engine {
	r := gin.Default()

	// Health-check (unversioned)
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong pong"})
	})

	// /api/v1 — all versioned domain routes live under here
	u := r.Group("/api")
	{
		users.RegisterRoutes(u)         // /api/v1/users, /api/v1/auth
		property.RegisterRoutes(u)      // /api/v1/properties
		subscriptions.RegisterRoutes(u) // /api/v1/subscriptions
		chat.RegisterRoutes(u)          // /api/v1/chats
		wishlist.RegisterRoutes(u)      // /api/v1/wishlist
		admin.RegisterRoutes(u)         // /api/v1/admin
	}

	// Swagger endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
