package service

import (
	"../dao"
	"../domain"
)

// UserService struct
type UserService struct{}

// FindByID func
func (userService UserService) FindByID(id int) (domain.User, error) {
	userDAO := &dao.UserDAO{}

	return userDAO.FindByID(id)
}

// Create func
func (userService UserService) Create(user domain.User) (domain.User, error) {
	userDAO := &dao.UserDAO{}

	return userDAO.Create(user)
}
