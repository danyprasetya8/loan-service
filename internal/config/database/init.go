package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New() (db *gorm.DB, err error) {
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")

	dsn := fmt.Sprintf("host=%s user=%s database=postgres password=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", host, user, password, port)
	db, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		return nil, err
	}

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

	dsn = fmt.Sprintf("host=%s user=%s database=%s password=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", host, user, dbName, password, port)
	return gorm.Open(postgres.Open(dsn))
}
