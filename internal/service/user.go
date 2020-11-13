package service

import (
	"../dao"
	"../domain"
)

// UserService struct
type UserService struct {
	dao dao.UserDAOInterface
}

//UserServiceInterface interface
type UserServiceInterface interface {
	Create(domain.User) (*domain.User, error)
	FindByID(int) (domain.User, error)
	GetAll() ([]domain.User, error)
}

// NewUserService constructor of UserService struct
func NewUserService() *UserService {
	return &UserService{dao: dao.NewUserDAO()}
}

// FindByID func
func (srv UserService) FindByID(id int) (domain.User, error) {
	return srv.dao.FindByID(id)
}

// Create func
func (srv UserService) Create(user domain.User) (*domain.User, error) {
	return srv.dao.Create(&user)
}

// GetAll func
func (srv UserService) GetAll() ([]domain.User, error) {
	return srv.dao.GetAll()
}
