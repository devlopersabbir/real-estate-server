package startup

import (
	"net/http"
	"strconv"

	"github.com/devlopersabbir/juan_don82-server/internal/pkg/config"
	"github.com/devlopersabbir/juan_don82-server/internal/pkg/types"
	"github.com/gin-gonic/gin"
)

func Server() {
	env, err := config.LoadEnv()

	if err != nil {
		panic("Failed to load environment variables: " + err.Error())
	}
	r := create(env)

	port := strconv.Itoa(env.ServerConfig.Port)
	r.Run(":" + port)
}

func create(*types.Env) *gin.Engine {
	// Create a Gin router with default middleware (logger and recovery)
	r := gin.Default()

	// all routes group gose here...
	r.GET("/ping", func(c *gin.Context) {
		// Return JSON response
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	return r
}
