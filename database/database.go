package database

import (
	"fmt"
	"jima/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDB(cfg config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		cfg.PostgresDBHost,
		cfg.PostgresDBUser,
		cfg.PostgresDBPassword,
		cfg.PostgresDBName,
		cfg.PostgresDBPort,
		cfg.PostgresDBSSLMode,
		cfg.PostgresDBTimezone,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to postgres database")
	}
	fmt.Println("Connected to Postgres database successfully")

	if cfg.PostgresDBDebugMode {
		db = db.Debug()
		fmt.Println("Database debug mode is enabled")
	}

	return db
}
