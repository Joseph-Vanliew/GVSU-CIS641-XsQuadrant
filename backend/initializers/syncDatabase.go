package initializers

import (
	"v/backend/models"
)

func SyncDatabase() {

	DB.AutoMigrate(&models.User{})

}
