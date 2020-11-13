package service

import (
	"errors"
	"time"

	"../dao"
	"../domain"
)

// QuestionService struct
type QuestionService struct{}

// NewQuestionService constructor of QuestionService struct
func NewQuestionService() *QuestionService {
	return &QuestionService{}
}

// Check if all attributes are filled
func checkStruct(question domain.Question, checkID bool) bool {
	if (question.ID == 0 && checkID) ||
		question.Statement == "" ||
		question.UserID == 0 {

		return false
	}

	return true
}

// Create func
func (questionService QuestionService) Create(question domain.Question) (domain.Question, error) {
	if !checkStruct(question, false) {
		emptyQuestion := &domain.Question{}
		return *emptyQuestion, errors.New("Not all attributes are filled")
	}

	questionDAO := dao.NewQuestionDAO()

	return questionDAO.Create(question)
}

// Update func
func (questionService QuestionService) Update(question domain.Question) (domain.Question, error) {
	if !checkStruct(question, true) {
		emptyQuestion := &domain.Question{}
		return *emptyQuestion, errors.New("Not all attributes are filled")
	}

	questionDAO := dao.NewQuestionDAO()
	question.UpdatedAt = time.Now()

	return questionDAO.Update(question)
}

// Delete func
func (questionService QuestionService) Delete(ID int) error {
	questionDAO := dao.NewQuestionDAO()

	return questionDAO.Delete(ID)
}

// FindByID func
func (questionService QuestionService) FindByID(id int) (domain.Question, error) {
	questionDAO := dao.NewQuestionDAO()

	return questionDAO.FindByID(id)
}

// FindByUser func
func (questionService QuestionService) FindByUser(userID int) ([]domain.Question, error) {
	questionDAO := dao.NewQuestionDAO()

	return questionDAO.FindByUser(userID)
}

// GetAll func
func (questionService QuestionService) GetAll() ([]domain.Question, error) {
	questionDAO := dao.NewQuestionDAO()

	return questionDAO.GetAll()
}
