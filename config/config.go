package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

const AppName = "JIMA"

type Config struct {
	// General
	Environment string `envconfig:"ENVIRONMENT" default:"development"`
	Port        int    `envconfig:"PORT" default:"8080"`

	// Database
	PostgresDBHost      string `envconfig:"POSTGRES_DB_HOST" default:"localhost"`
	PostgresDBPort      int    `envconfig:"POSTGRES_DB_PORT" default:"5432"`
	PostgresDBName      string `envconfig:"POSTGRES_DB_NAME" default:"app_db"`
	PostgresDBUser      string `envconfig:"POSTGRES_DB_USER" default:"user"`
	PostgresDBPassword  string `envconfig:"POSTGRES_DB_PASSWORD" default:"password"`
	PostgresDBDebugMode bool   `envconfig:"POSTGRES_DB_DEBUG_MODE" default:"false"`
}

// Get to get defined configuration
func Get() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := Config{}
	envconfig.MustProcess(AppName, &cfg)
	return cfg
}
