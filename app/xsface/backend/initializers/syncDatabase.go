package initializers

import (
	"xsface/models"
)

func SyncDatabase() {

	DB.AutoMigrate(&models.User{})

}
