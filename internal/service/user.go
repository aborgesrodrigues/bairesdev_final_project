package service

import (
	"fmt"
	"strconv"

	"bairesdev_final_project/internal/dao"
	"bairesdev_final_project/internal/domain"

	"go.uber.org/zap"
)

// UserService struct
type UserService struct {
	Dao    dao.UserDAOInterface
	logger *zap.Logger
}

//go:generate mockgen -source=user.go -destination=user_mock.go -package=service

//UserServiceInterface interface
type UserServiceInterface interface {
	Create(domain.User) (*domain.User, error)
	FindByID(int) (*domain.User, error)
	GetAll() (*[]domain.User, error)
}

// NewUserService constructor of UserService struct
func NewUserService(logger *zap.Logger) *UserService {
	return &UserService{Dao: dao.NewUserDAO(logger), logger: logger}
}

// FindByID func
func (srv UserService) FindByID(ID int) (*domain.User, error) {
	// Entry log
	srv.logger.Info("Called FindByID",
		zap.String("ID", strconv.Itoa(ID)),
	)

	return srv.Dao.FindByID(ID)
}

// Create func
func (srv UserService) Create(user domain.User) (*domain.User, error) {
	// Entry log
	srv.logger.Info("Called Create",
		zap.String("user", fmt.Sprintf("%#v", user)),
	)

	return srv.Dao.Create(user)
}

// GetAll func
func (srv UserService) GetAll() (*[]domain.User, error) {
	// Entry log
	srv.logger.Info("Called GetAll")

	return srv.Dao.GetAll()
}
