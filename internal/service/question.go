package service

import (
	"errors"
	"time"

	"../dao"
	"../domain"
)

// QuestionService struct
type QuestionService struct {
	dao *dao.QuestionDAO
}

// NewQuestionService constructor of QuestionService struct
func NewQuestionService() *QuestionService {
	return &QuestionService{dao: dao.NewQuestionDAO()}
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
func (srv QuestionService) Create(question domain.Question) (domain.Question, error) {
	if !checkStruct(question, false) {
		return domain.Question{}, errors.New("Not all attributes are filled")
	}

	return srv.dao.Create(question)
}

// Update func
func (srv QuestionService) Update(question domain.Question) (domain.Question, error) {
	if !checkStruct(question, true) {
		return domain.Question{}, errors.New("Not all attributes are filled")
	}

	question.UpdatedAt = time.Now()

	return srv.dao.Update(question)
}

// Delete func
func (srv QuestionService) Delete(ID int) error {
	return srv.dao.Delete(ID)
}

// FindByID func
func (srv QuestionService) FindByID(id int) (domain.Question, error) {
	return srv.dao.FindByID(id)
}

// FindByUser func
func (srv QuestionService) FindByUser(userID int) ([]domain.Question, error) {
	return srv.dao.FindByUser(userID)
}

// GetAll func
func (srv QuestionService) GetAll() ([]domain.Question, error) {
	return srv.dao.GetAll()
}
