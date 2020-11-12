package handler

import (
	"../dao"
	"../domain"
)

// QuestionService struct
type QuestionService struct{}

// FindByID func
func (questionService QuestionService) FindByID(id int) (domain.Question, error) {
	questionDAO := &dao.QuestionDAO{}

	return questionDAO.FindByID(id)
}

// Create func
func (questionService QuestionService) Create(question domain.Question) (domain.Question, error) {
	questionDAO := &dao.QuestionDAO{}

	return questionDAO.Create(question)
}
