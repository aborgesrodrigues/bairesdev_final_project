package main

import (
	"fmt"
	"time"

	"bairesdev_final_project/internal/domain"
	"bairesdev_final_project/internal/service"
)

// CallService func
func CallService() {
	user1Birthday, _ := time.Parse("2006-01-02", "2000-01-01")

	user1 := &domain.User{
		Username: "username1",
		Name:     "User 1",
		Birthday: user1Birthday,
	}

	userDao := &service.UserService{}
	userDao.Create(*user1)

	fmt.Println(userDao.FindByID(1))
}
