package migrate

import (
	"posts-api/internal/config"
	"posts-api/internal/models"
)

func MigrateDB() {
	config.ConnectDB()
	config.DB.AutoMigrate(&models.User{})
}
