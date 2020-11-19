package cmd

import (
	"fmt"
	"os"

	"bairesdev_final_project/internal/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// CreateDatabase func
func CreateDatabase() {
	var host = os.Getenv("DATABASE_HOST")
	var user = os.Getenv("DATABASE_USER")
	var password = os.Getenv("DATABASE_PASSWORD")
	var dbName = os.Getenv("DATABASE_NAME")

	dsnPostgres := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", host, user, dbName, password)

	db, _ := gorm.Open(postgres.Open(dsnPostgres), &gorm.Config{})

	db.AutoMigrate(&domain.User{}, &domain.Question{})
	//sqliteDatabase.Model(&domain.Question{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
}
