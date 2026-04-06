package config

type DatabaseConfig struct {
	DBHost    string `validate:"required"`
	DBUser    string `validate:"required"`
	DBPass    string `validate:"required"`
	DBName    string `validate:"required"`
	DBPort    int    `validate:"required"`
	DBSSLMode string `validate:"required"`
}
type ElasticConfig struct {
	ESAddresses []string `validate:"required"`
	ESUsername  string
	ESPassword  string
}
type ServerConfig struct {
	Port int    `validate:"required"`
	Host string `validate:"required"`
}
type CORSConfig struct {
	AllowOrigins []string
	AllowHeaders []string
	AllowMethods []string
}
type JWTConfig struct {
	Secret        string `validate:"required"`
	RefreshSecret string `validate:"required"`
}

type Env struct {
	DatabaseConfig DatabaseConfig
	ServerConfig   ServerConfig
	ElasticConfig  ElasticConfig
	JWTConfig      JWTConfig
}
