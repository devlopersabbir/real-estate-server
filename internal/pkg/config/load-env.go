package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func LoadEnv() (*Env, error) {
	_ = godotenv.Load()

	env := &Env{
		JWT_SECRET:         GetEnv("JWT_SECRET", "supersecretkey"),
		JWT_REFRESH_SECRET: GetEnv("JWT_REFRESH_SECRET", "superrefreshsecretkey"),
		ServerConfig: ServerConfig{
			Port: GetEnv("PORT", 9000),
			Host: GetEnv("HOST", "0.0.0.0"),
		},
		DatabaseConfig: DatabaseConfig{
			DBHost:    GetEnv("DB_HOST", "localhost"),
			DBPort:    GetEnv("DB_PORT", 5431),
			DBUser:    GetEnv("DB_USER", "postgres"),
			DBPass:    GetEnv("DB_PASSWORD", "postgres"),
			DBName:    GetEnv("DB_NAME", "juan"),
			DBSSLMode: GetEnv("DB_SSL_MODE", "disable"),
		},
		ElasticConfig: ElasticConfig{
			ESAddresses: []string{GetEnv("ES_ADDRESSES", "http://localhost:9200")},
			ESUsername:  GetEnv("ES_USERNAME", "elastic"),
			ESPassword:  GetEnv("ES_PASSWORD", "elastic"),
		},
		JWTConfig: JWTConfig{
			Secret:        GetEnv("JWT_SECRET", "supersecretkey"),
			RefreshSecret: GetEnv("JWT_REFRESH_SECRET", "superrefreshsecretkey"),
		},
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
