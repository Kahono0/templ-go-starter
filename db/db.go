package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(url string) {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	DB = db
}
