package cmd

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"../domain"
)

// PopulateUser func
func PopulateUser() {

	user1Birthday, _ := time.Parse("2006-01-02", "2000-01-01")

	user1 := &domain.User{
		Username: "username1",
		Name:     "User 1",
		Birthday: user1Birthday,
	}

	user2Birthday, _ := time.Parse("2006-01-02", "1999-05-03")
	user2 := &domain.User{
		Username: "username2",
		Name:     "User 2",
		Birthday: user2Birthday,
	}

	user3Birthday, _ := time.Parse("2006-01-02", "1999-07-15")
	user3 := &domain.User{
		Username: "username3",
		Name:     "User 3",
		Birthday: user3Birthday,
	}

	user4Birthday, _ := time.Parse("2006-01-02", "2002-08-19")
	user4 := &domain.User{
		Username: "username4",
		Name:     "User 4",
		Birthday: user4Birthday,
	}

	db, _ := gorm.Open(sqlite.Open("sqlite-database.db"), &gorm.Config{})

	db.Create(&user1)
	db.Create(&user2)
	db.Create(&user3)
	db.Create(&user4)
}
