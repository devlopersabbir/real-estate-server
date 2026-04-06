package migrations

import (
	"log"

	"github.com/devlopersabbir/juan_don82-server/api/users/core"
	"gorm.io/gorm"
)

func Automigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&core.Users{},
	)
	if err != nil {
		log.Fatalf("Fail to auto migrate: %s", err.Error())
	}
	log.Println("Auto migration completed successfully")
}
