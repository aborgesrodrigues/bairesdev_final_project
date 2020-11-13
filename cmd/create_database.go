package cmd

import (
	"log"
	"os"

	"../internal/domain"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// CreateDatabase func
func CreateDatabase() {
	os.Remove("../sqlite-database.db")

	file, err := os.Create("../sqlite-database.db") // Create SQLite file
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	file.Close()

	sqliteDatabase, openError := gorm.Open(sqlite.Open("../sqlite-database.db"), &gorm.Config{})

	if openError != nil {
		log.Fatal(openError.Error())
	}

	sqliteDatabase.AutoMigrate(&domain.User{})
	sqliteDatabase.AutoMigrate(&domain.Question{})
}
