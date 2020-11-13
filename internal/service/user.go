package service

import (
	"../dao"
	"../domain"
)

// UserService struct
type UserService struct{}

// NewUserService constructor of UserService struct
func NewUserService() *UserService {
	return &UserService{}
}

// FindByID func
func (userService UserService) FindByID(id int) (domain.User, error) {
	userDAO := dao.NewUserDAO()

	return userDAO.FindByID(id)
}

// Create func
func (userService UserService) Create(user domain.User) (domain.User, error) {
	userDAO := dao.NewUserDAO()

	return userDAO.Create(user)
}

// GetAll func
func (userService UserService) GetAll() ([]domain.User, error) {
	userDAO := dao.NewUserDAO()

	return userDAO.GetAll()
}
