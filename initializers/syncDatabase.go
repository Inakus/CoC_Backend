package initializers

import "github.com/Inakus/CoC_Backend/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
