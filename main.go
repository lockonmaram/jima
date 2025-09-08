package main

import (
	"jima/config"
	"jima/database"
)

func main() {
	// Application entry point
	config := config.Get()
	database.NewPostgresDB(config)
}
