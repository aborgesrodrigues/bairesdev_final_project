package service

import (
	"fmt"
	"strconv"

	"final_project/internal/dao"
	"final_project/internal/domain"

	"go.uber.org/zap"
)

// UserService struct
type UserService struct {
	dao    dao.UserDAOInterface
	logger *zap.Logger
}

//UserServiceInterface interface
type UserServiceInterface interface {
	Create(domain.User) (*domain.User, error)
	FindByID(int) (*domain.User, error)
	GetAll() ([]domain.User, error)
}

// NewUserService constructor of UserService struct
func NewUserService(logger *zap.Logger) *UserService {
	return &UserService{dao: dao.NewUserDAO(logger), logger: logger}
}

// FindByID func
func (srv UserService) FindByID(ID int) (*domain.User, error) {
	// Entry log
	srv.logger.Info("Called FindByID",
		zap.String("ID", strconv.Itoa(ID)),
	)

	return srv.dao.FindByID(ID)
}

// Create func
func (srv UserService) Create(user domain.User) (*domain.User, error) {
	// Entry log
	srv.logger.Info("Called Create",
		zap.String("user", fmt.Sprintf("%#v", user)),
	)

	return srv.dao.Create(&user)
}

// GetAll func
func (srv UserService) GetAll() ([]domain.User, error) {
	// Entry log
	srv.logger.Info("Called GetAll")

	return srv.dao.GetAll()
}
