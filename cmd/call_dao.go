package cmd

import (
	"fmt"

	"../internal/dao"
)

//CallDao func
func CallDao() {
	// user1Birthday, _ := time.Parse("2006-01-02", "2000-01-01")

	// user1 := &domain.User{
	// 	Username: "username1",
	// 	Name:     "User 1",
	// 	Birthday: user1Birthday,
	// }

	userDao := &dao.UserDAO{}
	// userDao.Create(*user1)

	fmt.Println(userDao.FindByID(1))
}
