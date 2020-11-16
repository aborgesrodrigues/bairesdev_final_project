package service

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"../dao"
	"../domain"
	"go.uber.org/zap"
)

// QuestionService struct
type QuestionService struct {
	dao    dao.QuestionDAOInterface
	logger *zap.Logger
}

//QuestionServiceInterface interface
type QuestionServiceInterface interface {
	Create(domain.Question) (*domain.Question, error)
	Update(domain.Question) (*domain.Question, error)
	Delete(int) error
	FindByID(int) (*domain.Question, error)
	FindByUser(int) (*[]domain.Question, error)
	GetAll() (*[]domain.Question, error)
}

// NewQuestionService constructor of QuestionService struct
func NewQuestionService(logger *zap.Logger) *QuestionService {
	return &QuestionService{dao: dao.NewQuestionDAO(logger), logger: logger}
}

// Check if all attributes are filled
func (srv QuestionService) checkStruct(question domain.Question, checkID bool) bool {
	// Entry log
	srv.logger.Info("Called checkStruct",
		zap.String("question", fmt.Sprintf("%#v", question)),
		zap.String("checkID", strconv.FormatBool(checkID)),
	)

	if (question.ID == 0 && checkID) ||
		question.Statement == "" ||
		question.UserID == 0 {

		return false
	}

	return true
}

// Create func
func (srv QuestionService) Create(question domain.Question) (*domain.Question, error) {
	// Entry log
	srv.logger.Info("Called Create",
		zap.String("question", fmt.Sprintf("%#v", question)),
	)

	if !srv.checkStruct(question, false) {
		//log
		message := "Not all attributes are filled"
		srv.logger.Info(message,
			zap.String("question", fmt.Sprintf("%#v", question)),
		)
		return &domain.Question{}, errors.New(message)
	}

	return srv.dao.Create(&question)
}

// Update func
func (srv QuestionService) Update(question domain.Question) (*domain.Question, error) {
	// Entry log
	srv.logger.Info("Called Update",
		zap.String("question", fmt.Sprintf("%#v", question)),
	)

	if !srv.checkStruct(question, true) {
		//log
		message := "Not all attributes are filled"
		srv.logger.Info(message,
			zap.String("question", fmt.Sprintf("%#v", question)),
		)

		return &domain.Question{}, errors.New(message)
	}

	question.UpdatedAt = time.Now()

	return srv.dao.Update(&question)
}

// Delete func
func (srv QuestionService) Delete(ID int) error {
	// Entry log
	srv.logger.Info("Called Delete",
		zap.String("ID", strconv.Itoa(ID)),
	)

	return srv.dao.Delete(ID)
}

// FindByID func
func (srv QuestionService) FindByID(ID int) (*domain.Question, error) {
	// Entry log
	srv.logger.Info("Called FindByID",
		zap.String("ID", strconv.Itoa(ID)),
	)

	return srv.dao.FindByID(ID)
}

// FindByUser func
func (srv QuestionService) FindByUser(userID int) (*[]domain.Question, error) {
	// Entry log
	srv.logger.Info("Called FindByUser",
		zap.String("userID", strconv.Itoa(userID)),
	)

	return srv.dao.FindByUser(userID)
}

// GetAll func
func (srv QuestionService) GetAll() (*[]domain.Question, error) {
	// Entry log
	srv.logger.Info("Called GetAll")

	return srv.dao.GetAll()
}
