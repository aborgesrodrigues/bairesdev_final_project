package dao

import (
	"fmt"
	"strconv"

	"bairesdev_final_project/internal/domain"

	"go.uber.org/zap"
)

// UserDAO struct
type UserDAO struct {
	d      *dao
	logger *zap.Logger
}

//go:generate mockgen -source=user.go -destination=user_mock.go -package=dao

//UserDAOInterface interface
type UserDAOInterface interface {
	Create(*domain.User) (*domain.User, error)
	FindByID(int) (*domain.User, error)
	GetAll() ([]domain.User, error)
}

// NewUserDAO Constructor of UserDAO struct
func NewUserDAO(logger *zap.Logger) *UserDAO {
	return &UserDAO{d: newDAO(logger), logger: logger}
}

// Create func
func (userDAO UserDAO) Create(user *domain.User) (*domain.User, error) {
	// Entry log
	userDAO.logger.Info("Called Create",
		zap.String("user", fmt.Sprintf("%#v", user)),
	)

	createdUser, err := userDAO.d.create(user)

	if err != nil {
		// Entry log
		userDAO.logger.Error(err.Error(),
			zap.String("user", fmt.Sprintf("%#v", user)),
		)

		return createdUser.(*domain.User), err
	}

	return createdUser.(*domain.User), nil
}

// FindByID func
func (userDAO UserDAO) FindByID(ID int) (*domain.User, error) {
	// Entry log
	userDAO.logger.Info("Called FindByID",
		zap.String("ID", strconv.Itoa(ID)),
	)

	user, err := userDAO.d.findByID(&domain.User{}, ID)

	if err != nil {
		// Entry log
		userDAO.logger.Error(err.Error(),
			zap.String("user", fmt.Sprintf("%#v", user)),
		)

		return user.(*domain.User), err
	}

	return user.(*domain.User), nil
}

// GetAll func
func (userDAO UserDAO) GetAll() ([]domain.User, error) {
	// Entry log
	userDAO.logger.Info("Called GetAll")

	users, err := userDAO.d.getAll(make([]domain.User, 0))

	if err != nil {
		// Entry log
		userDAO.logger.Error(err.Error())

		return nil, err
	}

	return users.([]domain.User), nil
}
