package main

import (
	"posts-api/internal/config"
	"posts-api/internal/handlers"
	"posts-api/internal/migrate"
)

func main() {
	config.ConnectDB()
	migrate.MigrateDB()
	handlers.Handlers()

}
