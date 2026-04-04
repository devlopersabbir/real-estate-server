package config

import (
	"os"
	"strconv"

	"github.com/devlopersabbir/juan_don82-server/internal/pkg/types"
	"github.com/joho/godotenv"
)

func LoadEnv() (*types.Env, error) {
	err := godotenv.Load()

	env := &types.Env{
		ServerConfig: types.ServerConfig{
			Port: GetEnv("PORT", 5000),
			Host: GetEnv("HOST", "0.0.0.0"),
		},
	}
	if err != nil {
		return nil, err
	}
	return env, nil
}

// Generic function for taking string & int both
func GetEnv[T string | int](key string, defaultValue T) T {
	if value := os.Getenv(key); value != "" {
		var zero T
		switch any(zero).(type) {
		case string:
			return any(value).(T)
		case int:
			if intVal, err := strconv.Atoi(value); err == nil {
				return any(intVal).(T)
			}
		}
	}
	return defaultValue
}
