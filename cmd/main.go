package main

import (
	"strconv"

	"github.com/devlopersabbir/juan_don82-server/internal/pkg/config"
	"github.com/devlopersabbir/juan_don82-server/startup"
)

func main() {
	env, err := config.LoadEnv()

	if err != nil {
		panic("Failed to load environment variables: " + err.Error())
	}
	r := startup.Server(env)

	port := strconv.Itoa(env.ServerConfig.Port)
	r.Run(":" + port)
}
