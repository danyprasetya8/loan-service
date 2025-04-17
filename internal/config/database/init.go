package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New() (db *gorm.DB, err error) {
	dsn := os.Getenv("POSTGRES_DSN")
	db, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		return nil, err
	}

	dbName := os.Getenv("POSTGRES_DB")
	checkQuery := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = '%s')", dbName)

	dbExist := false
	err = db.Raw(checkQuery).
		Scan(&dbExist).
		Error

	if err != nil {
		return nil, err
	}

	if !dbExist {
		if err = db.Exec("CREATE DATABASE " + dbName).Error; err != nil {
			return nil, err
		}
	}

	dsn = dsn + " dbname=" + dbName
	return gorm.Open(postgres.Open(dsn))
}
