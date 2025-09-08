package database

import (
	"fmt"
	"jima/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDB(cfg config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		cfg.PostgresDBHost,
		cfg.PostgresDBUser,
		cfg.PostgresDBPassword,
		cfg.PostgresDBName,
		fmt.Sprint(cfg.PostgresDBPort),
		"UTC",
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to postgres database")
	}
	fmt.Println("Connected to Postgres database successfully")

	return db
}
