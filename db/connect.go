package db

import (
	"gorm.io/gorm"

	"gorm.io/driver/sqlite"
)

func Connect() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("db/estoque.db"), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
