package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() {
	d, err := gorm.Open(sqlite.Open("book.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}