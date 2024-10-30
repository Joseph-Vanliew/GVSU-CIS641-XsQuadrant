package initializers

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnvVariables() {
	err := godotenv.Load("app/xsface/backend/.env")
	if err != nil {
		log.Fatal("Error loading .env variables")
	} else {
		log.Println("Loaded .env variables")
	}
}
