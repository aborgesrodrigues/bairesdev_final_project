package dao_test

import (
	"errors"
	"testing"

	"bairesdev_final_project/internal/dao"
	"bairesdev_final_project/internal/domain"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

// TestCreateQuestion tests the CreateQuestion function of the handler
func TestCreateQuestion(t *testing.T) {
	// Data
	argument := domain.Question{
		Statement: "Statement 1",
		UserID:    1,
	}

	logger, _ := zap.NewProduction()

	dao := dao.NewQuestionDAO(logger)

	// call the service
	question, error := dao.Create(argument)

	assert.NoError(t, error)
	assert.Equal(t, question.Statement, argument.Statement, "they should be equal")
	assert.Equal(t, question.Answer, argument.Answer, "they should be equal")
	assert.Equal(t, question.UserID, argument.UserID, "they should be equal")
}

func TestUpdateQuestion(t *testing.T) {
	logger, _ := zap.NewProduction()

	// get a question to change
	dao := dao.NewQuestionDAO(logger)
	allQuestions, _ := dao.GetAll()
	argument := (*allQuestions)[0]
	argument.Statement = "Statement"
	argument.UserID = 2

	// call the service
	question, error := dao.Update(argument)

	assert.NoError(t, error)
	assert.Equal(t, question, &argument, "they should be equal")
}

func TestDeleteQuestion(t *testing.T) {
	logger, _ := zap.NewProduction()

	dao := dao.NewQuestionDAO(logger)
	allQuestions, _ := dao.GetAll()
	argument := (*allQuestions)[0]

	// call the service
	error := dao.Delete(int(argument.ID))

	assert.NoError(t, error)
}

func TestFailDeleteQuestion(t *testing.T) {
	logger, _ := zap.NewProduction()

	dao := dao.NewQuestionDAO(logger)

	// call the service
	error := dao.Delete(111)

	assert.Equal(t, error, errors.New("ID not found"), "they should be equal")
}
