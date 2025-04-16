package database

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New() (*gorm.DB, error) {
	dsn := os.Getenv("POSTGRES_DSN")
	return gorm.Open(postgres.Open(dsn))
}
