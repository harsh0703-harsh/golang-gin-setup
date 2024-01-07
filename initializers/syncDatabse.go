package initializers

import "gin/models"

func SyncDatabases() {

	DB.AutoMigrate(&models.User{})
}
