package startup

import (
	"net/http"

	"github.com/devlopersabbir/juan_don82-server/api/property"
	"github.com/devlopersabbir/juan_don82-server/api/users"
	"github.com/devlopersabbir/juan_don82-server/internal/pkg/config"
	"github.com/gin-gonic/gin"
)

func Server(*config.Env) *gin.Engine {
	r := gin.Default()

	// Health-check (unversioned)
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong pong"})
	})

	// /api/v1 — all versioned domain routes live under here
	u := r.Group("/api")
	{
		users.RegisterRoutes(u)    // /api/v1/users, /api/v1/auth
		property.RegisterRoutes(u) // /api/v1/properties
	}

	return r
}
