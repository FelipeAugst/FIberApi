package db

import (
	"api/config"

	"gorm.io/gorm"

	"gorm.io/driver/sqlite"
)

func Connect() (*gorm.DB, error) {
	connectString := config.DbName
	db, err := gorm.Open(sqlite.Open(connectString), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
