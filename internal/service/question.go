package service

import (
	"errors"
	"fmt"
	"strconv"

	"bairesdev_final_project/internal/dao"
	"bairesdev_final_project/internal/domain"

	"go.uber.org/zap"
)

// QuestionService struct
type QuestionService struct {
	Dao    dao.QuestionDAOInterface
	logger *zap.Logger
}

//go:generate mockgen -source=question.go -destination=question_mock.go -package=service

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
	return &QuestionService{Dao: dao.NewQuestionDAO(logger), logger: logger}
}

// Create func
func (srv QuestionService) Create(question domain.Question) (*domain.Question, error) {
	// Entry log
	srv.logger.Info("Called Create",
		zap.String("question", fmt.Sprintf("%#v", question)),
	)

	if question.IsEmpty(false) {
		//log
		message := "Not all attributes are filled"
		srv.logger.Info(message,
			zap.String("question", fmt.Sprintf("%#v", question)),
		)
		return nil, errors.New(message)
	}

	return srv.Dao.Create(question)
}

// Update func
func (srv QuestionService) Update(question domain.Question) (*domain.Question, error) {
	// Entry log
	srv.logger.Info("Called Update",
		zap.String("question", fmt.Sprintf("%#v", question)),
	)

	if question.IsEmpty(true) {
		//log
		message := "Not all attributes are filled"
		srv.logger.Info(message,
			zap.String("question", fmt.Sprintf("%#v", question)),
		)

		return nil, errors.New(message)
	}

	return srv.Dao.Update(question)
}

// Delete func
func (srv QuestionService) Delete(ID int) error {
	// Entry log
	srv.logger.Info("Called Delete",
		zap.String("ID", strconv.Itoa(ID)),
	)

	return srv.Dao.Delete(ID)
}

// FindByID func
func (srv QuestionService) FindByID(ID int) (*domain.Question, error) {
	// Entry log
	srv.logger.Info("Called FindByID",
		zap.String("ID", strconv.Itoa(ID)),
	)

	return srv.Dao.FindByID(ID)
}

// FindByUser func
func (srv QuestionService) FindByUser(userID int) (*[]domain.Question, error) {
	// Entry log
	srv.logger.Info("Called FindByUser",
		zap.String("userID", strconv.Itoa(userID)),
	)

	return srv.Dao.FindByUser(userID)
}

// GetAll func
func (srv QuestionService) GetAll() (*[]domain.Question, error) {
	// Entry log
	srv.logger.Info("Called GetAll")

	return srv.Dao.GetAll()
}
