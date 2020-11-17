package cmd

import (
	"log"
	"os"

	"bairesdev_final_project/internal/domain"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// CreateDatabase func
func CreateDatabase() {
	var sqlitePash = os.Getenv("SQLITEPATH") + "sqlite-database.db"

	os.Remove(sqlitePash)

	file, err := os.Create(sqlitePash) // Create SQLite file
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	file.Close()

	sqliteDatabase, openError := gorm.Open(sqlite.Open(sqlitePash), &gorm.Config{})

	if openError != nil {
		log.Fatal(openError.Error())
	}

	sqliteDatabase.AutoMigrate(&domain.User{})
	sqliteDatabase.AutoMigrate(&domain.Question{})
}
