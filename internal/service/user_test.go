package service_test

import (
	"bairesdev_final_project/internal/dao"
	"bairesdev_final_project/internal/domain"
	"bairesdev_final_project/internal/service"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestCreateUser(t *testing.T) {
	// Data
	argument := domain.User{
		Username: "username1",
		Name:     "User Name 1",
	}

	logger, _ := zap.NewProduction()

	service := service.NewUserService(logger)

	// create the question mock interface
	ctrl := gomock.NewController(t)
	service.Dao = dao.NewMockUserDAOInterface(ctrl)
	service.Dao.(*dao.MockUserDAOInterface).EXPECT().Create(argument).Return(
		&domain.User{
			ID:       1,
			Username: "username1",
			Name:     "User Name 1",
		}, nil)

	// call the service
	question, error := service.Create(argument)

	//included to facilitate the check with the returned value
	argument.ID = 1

	assert.NoError(t, error)
	assert.Equal(t, question, &argument, "they should be equal")
}

func TestFindUserByID(t *testing.T) {
	foundUser := domain.User{
		ID:       1,
		Username: "username1",
		Name:     "User Name 1",
	}
	logger, _ := zap.NewProduction()

	service := service.NewUserService(logger)

	// create the question mock interface
	ctrl := gomock.NewController(t)
	service.Dao = dao.NewMockUserDAOInterface(ctrl)
	service.Dao.(*dao.MockUserDAOInterface).EXPECT().FindByID(1).Return(
		&foundUser, nil,
	)

	// call the service
	question, error := service.FindByID(1)

	assert.NoError(t, error)
	assert.Equal(t, question, &foundUser, "they should be equal")
}

func TestGetAllUsers(t *testing.T) {
	allUsers := make([]domain.User, 0)
	allUsers = append(allUsers, domain.User{
		ID:       1,
		Username: "username1",
		Name:     "User Name 1",
	})
	allUsers = append(allUsers, domain.User{
		ID:       2,
		Username: "username2",
		Name:     "User Name 2",
	})
	logger, _ := zap.NewProduction()

	service := service.NewUserService(logger)

	// create the question mock interface
	ctrl := gomock.NewController(t)
	service.Dao = dao.NewMockUserDAOInterface(ctrl)
	service.Dao.(*dao.MockUserDAOInterface).EXPECT().GetAll().Return(
		&allUsers, nil,
	)

	// call the service
	questions, error := service.GetAll()

	assert.NoError(t, error)
	assert.Equal(t, questions, &allUsers, "they should be equal")
}
