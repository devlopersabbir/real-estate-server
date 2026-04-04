package types

type Database struct {
	DBHost string
	DBUser string
	DBPass string
	DBName string
	DBPort string
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
	Database     Database
	ServerConfig ServerConfig
}
