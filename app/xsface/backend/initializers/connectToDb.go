package initializers

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectToDb() {
	dsn := os.Getenv("DB")

	// Open the database connection with PreferSimpleProtocol
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // Disable prepared statements
	}),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Info), // keep this
		})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	} else {
		log.Println("Database connection established successfully with PreferSimpleProtocol enabled")
	}

	//reassigning our local variable to the global
	DB = db
}
