package database

import (
	"fmt"
	"log"

	"github.com/devlopersabbir/juan_don82-server/internal/migrations"
	"github.com/devlopersabbir/juan_don82-server/internal/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func OpenConnection(env *config.Env) {
	dsn := generateConnectionString(env)
	fmt.Println("dsn", dsn)
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Couldn't establish database connection: %s", err)
	}
	// audo migrations
	migrations.Automigrate(DB)
	// redis store if in future needs

	// Elastic search connection
	ESClientConnection(&env.ElasticConfig)
}

func generateConnectionString(env *config.Env) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%v sslmode=%s", env.DatabaseConfig.DBHost, env.DatabaseConfig.DBUser, env.DatabaseConfig.DBPass, env.DatabaseConfig.DBName, env.DatabaseConfig.DBPort, env.DatabaseConfig.DBSSLMode)
}
