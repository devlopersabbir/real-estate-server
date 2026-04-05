package types

type DatabaseConfig struct {
	DBHost string `validate:"required"`
	DBUser string `validate:"required"`
	DBPass string `validate:"required"`
	DBName string `validate:"required"`
	DBPort int    `validate:"required"`
}
type ServerConfig struct {
	Port int    `validate:"required,gte=5000,lte=9000"` // port need to be within 5000 to 9000
	Host string `validate:"required"`
}
type CORSConfig struct {
	AllowOrigins []string
	AllowHeaders []string
	AllowMethods []string
}
type Env struct {
	DatabaseConfig DatabaseConfig
	ServerConfig   ServerConfig
}
