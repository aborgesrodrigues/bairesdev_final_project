package service_test

import (
	"context"
	"errors"
	"testing"

	"bairesdev_final_project/internal/dao"
	"bairesdev_final_project/internal/domain"
	"bairesdev_final_project/internal/service"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

// TestCreateQuestion tests the CreateQuestion function of the handler
func TestCreateQuestion(t *testing.T) {
	logger, _ := zap.NewProduction()

	service := service.NewQuestionService(logger)

	// create the question mock interface
	ctrl := gomock.NewController(t)
	service.Dao = dao.NewMockQuestionDAOInterface(ctrl)

	ctx := context.Background()

	t.Run("Sucess", func(t *testing.T) {
		// Data
		argument := domain.Question{
			Statement: "Statement 1",
			UserID:    1,
		}

		service.Dao.(*dao.MockQuestionDAOInterface).EXPECT().Create(argument).Return(
			&domain.Question{
				ID:        1,
				UserID:    1,
				Statement: "Statement 1",
			}, nil)

		// call the service
		question, error := service.Create(ctx, argument)

		//included to facilitate the check with the returned value
		argument.ID = 1

		assert.NoError(t, error)
		assert.Equal(t, question, &argument, "they should be equal")
	})

	t.Run("Fail", func(t *testing.T) {
		// Data
		argument := domain.Question{
			UserID: 1,
		}

		service.Dao.(*dao.MockQuestionDAOInterface).EXPECT().Create(argument).Return(
			nil, nil)

		// call the service
		question, error := service.Create(ctx, argument)

		assert.Equal(t, error.Error(), "Not all attributes are filled", "they should be equal")
		assert.Nil(t, question)
	})
}

func TestUpdateQuestion(t *testing.T) {
	logger, _ := zap.NewProduction()

	service := service.NewQuestionService(logger)

	// create the question mock interface
	ctrl := gomock.NewController(t)
	service.Dao = dao.NewMockQuestionDAOInterface(ctrl)

	ctx := context.Background()

	t.Run("Sucess", func(t *testing.T) {
		// Data
		argument := domain.Question{
			ID:        1,
			Statement: "Statement 1",
			UserID:    1,
		}

		service.Dao.(*dao.MockQuestionDAOInterface).EXPECT().Update(argument).Return(
			&argument, nil)

		// call the service
		question, error := service.Update(ctx, argument)

		assert.NoError(t, error)
		assert.Equal(t, question, &argument, "they should be equal")
	})

	t.Run("Fail by imcomplete struct", func(t *testing.T) {
		// Data
		argument := domain.Question{
			UserID: 1,
		}

		service.Dao.(*dao.MockQuestionDAOInterface).EXPECT().Update(argument).Return(nil, nil)

		// call the service
		question, error := service.Update(ctx, argument)

		assert.Equal(t, error.Error(), "Not all attributes are filled", "they should be equal")
		assert.Nil(t, question)
	})

	t.Run("Fail by not found", func(t *testing.T) {
		// Data
		argument := domain.Question{
			ID:        1,
			Statement: "Statement 1",
			UserID:    1,
		}

		service.Dao.(*dao.MockQuestionDAOInterface).EXPECT().Update(argument).Return(
			nil, errors.New("ID not found"))

		// call the service
		question, error := service.Update(ctx, argument)

		assert.Equal(t, error.Error(), "ID not found", "they should be equal")
		assert.Nil(t, question)
	})
}

func TestDeleteQuestion(t *testing.T) {
	logger, _ := zap.NewProduction()

	service := service.NewQuestionService(logger)

	// create the question mock interface
	ctrl := gomock.NewController(t)
	service.Dao = dao.NewMockQuestionDAOInterface(ctrl)

	ctx := context.Background()

	t.Run("Sucess", func(t *testing.T) {
		service.Dao.(*dao.MockQuestionDAOInterface).EXPECT().Delete(1).Return(nil)

		// call the service
		error := service.Delete(ctx, 1)

		assert.NoError(t, error)
	})

	t.Run("Fail", func(t *testing.T) {
		service.Dao.(*dao.MockQuestionDAOInterface).EXPECT().Delete(1).Return(
			errors.New("ID not found"),
		)

		// call the service
		error := service.Delete(ctx, 1)

		assert.Equal(t, error.Error(), "ID not found", "they should be equal")
	})
}

func TestFindQuestionByID(t *testing.T) {
	foundQuestion := domain.Question{
		ID:        1,
		Statement: "Statement 1",
		UserID:    1,
	}
	logger, _ := zap.NewProduction()

	service := service.NewQuestionService(logger)
	ctx := context.Background()

	// create the question mock interface
	ctrl := gomock.NewController(t)
	service.Dao = dao.NewMockQuestionDAOInterface(ctrl)
	service.Dao.(*dao.MockQuestionDAOInterface).EXPECT().FindByID(1).Return(
		&foundQuestion, nil,
	)

	// call the service
	question, error := service.FindByID(ctx, 1)

	assert.NoError(t, error)
	assert.Equal(t, question, &foundQuestion, "they should be equal")
}

func TestFindQuestionsByUser(t *testing.T) {
	userQuestions := make([]domain.Question, 0)
	userQuestions = append(userQuestions, domain.Question{
		ID:        1,
		Statement: "statement 1",
		UserID:    1,
	})
	userQuestions = append(userQuestions, domain.Question{
		ID:        2,
		Statement: "statement 2",
		UserID:    1,
	})
	logger, _ := zap.NewProduction()

	service := service.NewQuestionService(logger)

	ctx := context.Background()

	// create the question mock interface
	ctrl := gomock.NewController(t)
	service.Dao = dao.NewMockQuestionDAOInterface(ctrl)
	service.Dao.(*dao.MockQuestionDAOInterface).EXPECT().FindByUser(1).Return(
		&userQuestions, nil,
	)

	// call the service
	questions, error := service.FindByUser(ctx, 1)

	assert.NoError(t, error)
	assert.Equal(t, questions, &userQuestions, "they should be equal")
}

func TestGetAllQuestions(t *testing.T) {
	allQuestions := make([]domain.Question, 0)
	allQuestions = append(allQuestions, domain.Question{
		ID:        1,
		Statement: "statement 1",
		UserID:    1,
	})
	allQuestions = append(allQuestions, domain.Question{
		ID:        2,
		Statement: "statement 2",
		UserID:    2,
	})
	logger, _ := zap.NewProduction()

	service := service.NewQuestionService(logger)

	ctx := context.Background()

	// create the question mock interface
	ctrl := gomock.NewController(t)
	service.Dao = dao.NewMockQuestionDAOInterface(ctrl)
	service.Dao.(*dao.MockQuestionDAOInterface).EXPECT().GetAll().Return(
		&allQuestions, nil,
	)

	// call the service
	questions, error := service.GetAll(ctx)

	assert.NoError(t, error)
	assert.Equal(t, questions, &allQuestions, "they should be equal")
}
