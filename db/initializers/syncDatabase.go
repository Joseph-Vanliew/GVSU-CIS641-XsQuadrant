package initializers

import (
	"v/db/models"
)

func SyncDatabase() {

	DB.AutoMigrate(&models.User{})

}
