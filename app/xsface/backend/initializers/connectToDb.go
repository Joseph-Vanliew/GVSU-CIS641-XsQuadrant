package initializers

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func ConnectToDb() {
	var err error
	dsn := os.Getenv("DB")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}) //reassign DB and err

	if err != nil {
		panic("Failed to connect to database")
	}
}
