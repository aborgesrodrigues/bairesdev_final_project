package main

import (
	"fmt"
	"os"

	"bairesdev_final_project/internal/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// CreateDatabase func
func main() {
	var host = os.Getenv("DATABASE_HOST")
	var user = os.Getenv("DATABASE_USER")
	var password = os.Getenv("DATABASE_PASSWORD")
	var dbName = os.Getenv("DATABASE_NAME")
	var port = os.Getenv("DATABASE_PORT")

	dsn := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%s", host, user, dbName, password, port)

	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&domain.User{}, &domain.Question{})
	//sqliteDatabase.Model(&domain.Question{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
}
