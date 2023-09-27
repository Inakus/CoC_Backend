package initializers

import "github.com/Inakus/CoC_Backend/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Game{})
	DB.AutoMigrate(&models.Player{})
	DB.AutoMigrate(&models.PlayerResoureces{})
	DB.AutoMigrate(&models.Note{})
	DB.AutoMigrate(&models.Book{})
	DB.AutoMigrate(&models.BookContent{})
	DB.AutoMigrate(&models.Spell{})
}
