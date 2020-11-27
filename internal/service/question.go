package service

import (
	"context"
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
	Create(context.Context, domain.Question) (*domain.Question, error)
	Update(context.Context, domain.Question) (*domain.Question, error)
	Delete(context.Context, int) error
	FindByID(context.Context, int) (*domain.Question, error)
	FindByUser(context.Context, int) (*[]domain.Question, error)
	GetAll(context.Context) (*[]domain.Question, error)
}

// NewQuestionService constructor of QuestionService struct
func NewQuestionService(logger *zap.Logger) *QuestionService {
	return &QuestionService{Dao: dao.NewQuestionDAO(logger), logger: logger}
}

// Create func
func (srv QuestionService) Create(_ context.Context, question domain.Question) (*domain.Question, error) {
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
func (srv QuestionService) Update(_ context.Context, question domain.Question) (*domain.Question, error) {
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
func (srv QuestionService) Delete(_ context.Context, ID int) error {
	// Entry log
	srv.logger.Info("Called Delete",
		zap.String("ID", strconv.Itoa(ID)),
	)

	return srv.Dao.Delete(ID)
}

// FindByID func
func (srv QuestionService) FindByID(_ context.Context, ID int) (*domain.Question, error) {
	// Entry log
	srv.logger.Info("Called FindByID",
		zap.String("ID", strconv.Itoa(ID)),
	)

	return srv.Dao.FindByID(ID)
}

// FindByUser func
func (srv QuestionService) FindByUser(_ context.Context, userID int) (*[]domain.Question, error) {
	// Entry log
	srv.logger.Info("Called FindByUser",
		zap.String("userID", strconv.Itoa(userID)),
	)

	return srv.Dao.FindByUser(userID)
}

// GetAll func
func (srv QuestionService) GetAll(_ context.Context) (*[]domain.Question, error) {
	// Entry log
	srv.logger.Info("Called GetAll")

	return srv.Dao.GetAll()
}
