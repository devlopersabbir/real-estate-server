package startup

import (
	"net/http"

	"github.com/devlopersabbir/juan_don82-server/api/property"
	"github.com/devlopersabbir/juan_don82-server/api/users"
	"github.com/devlopersabbir/juan_don82-server/internal/pkg/types"
	"github.com/gin-gonic/gin"
)

func Server(*types.Env) *gin.Engine {
	r := gin.Default()

	// Health-check (unversioned)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// /api/v1 — all versioned domain routes live under here
	v1 := r.Group("/api")
	{
		users.RegisterRoutes(v1)    // /api/v1/users, /api/v1/auth
		property.RegisterRoutes(v1) // /api/v1/properties
	}

	return r
}
