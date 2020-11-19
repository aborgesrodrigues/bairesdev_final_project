package dao

import (
	"fmt"
	"strconv"

	"bairesdev_final_project/internal/domain"

	"go.uber.org/zap"
)

// QuestionDAO struct
type QuestionDAO struct {
	d      *dao
	logger *zap.Logger
}

//go:generate mockgen -source=question.go -destination=question_mock.go -package=dao

//QuestionDAOInterface interface
type QuestionDAOInterface interface {
	Create(domain.Question) (*domain.Question, error)
	Update(domain.Question) (*domain.Question, error)
	Delete(int) error
	FindByID(int) (*domain.Question, error)
	FindByUser(int) (*[]domain.Question, error)
	GetAll() (*[]domain.Question, error)
}

// NewQuestionDAO Constructor of QuestionDAO struct
func NewQuestionDAO(logger *zap.Logger) *QuestionDAO {
	return &QuestionDAO{d: newDAO(logger), logger: logger}
}

// Create func
func (questionDAO QuestionDAO) Create(question domain.Question) (*domain.Question, error) {
	// Entry log
	questionDAO.logger.Info("Called Create",
		zap.String("question", fmt.Sprintf("%#v", question)),
	)

	createdQuestion, err := questionDAO.d.create(&question)

	if err != nil {
		// log
		questionDAO.logger.Error(err.Error(),
			zap.String("question", fmt.Sprintf("%#v", question)),
		)

		return nil, err
	}

	return createdQuestion.(*domain.Question), nil
}

// Update func
func (questionDAO QuestionDAO) Update(question domain.Question) (*domain.Question, error) {
	// Entry log
	questionDAO.logger.Info("Called Update",
		zap.String("question", fmt.Sprintf("%#v", question)),
	)

	updatedQuestion, err := questionDAO.d.update(&question)

	if err != nil {
		// log
		questionDAO.logger.Error(err.Error(),
			zap.String("question", fmt.Sprintf("%#v", question)),
		)

		return nil, err
	}

	return updatedQuestion.(*domain.Question), nil
}

// Delete func
func (questionDAO QuestionDAO) Delete(ID int) error {
	// Entry log
	questionDAO.logger.Info("Called Delete",
		zap.String("ID", strconv.Itoa(ID)),
	)

	err := questionDAO.d.delete(&domain.Question{}, ID)

	if err != nil {
		// log
		questionDAO.logger.Error(err.Error(),
			zap.String("ID", strconv.Itoa(ID)),
		)

		return err
	}

	return nil
}

// FindByID func
func (questionDAO QuestionDAO) FindByID(ID int) (*domain.Question, error) {
	// Entry log
	questionDAO.logger.Info("Called FindByID",
		zap.String("ID", strconv.Itoa(ID)),
	)

	question, err := questionDAO.d.findByID(&domain.Question{}, ID)

	if err != nil {
		// log
		questionDAO.logger.Error(err.Error(),
			zap.String("question", fmt.Sprintf("%#v", question)),
		)

		return nil, err
	}

	return question.(*domain.Question), nil
}

// FindByUser func
func (questionDAO QuestionDAO) FindByUser(userID int) (*[]domain.Question, error) {
	// Entry log
	questionDAO.logger.Info("Called FindByUser",
		zap.String("userID", strconv.Itoa(userID)),
	)

	var questions = make([]domain.Question, 0)
	tx := questionDAO.d.db.Where(&domain.Question{UserID: uint(userID)}).Find(&questions)

	if tx.Error != nil {
		// log
		questionDAO.logger.Error(tx.Error.Error())
		return nil, tx.Error
	}

	return &questions, nil
}

// GetAll func
func (questionDAO QuestionDAO) GetAll() (*[]domain.Question, error) {
	// Entry log
	questionDAO.logger.Info("Called GetAll")

	arr := make([]domain.Question, 0)
	questions, err := questionDAO.d.getAll(&arr)

	if err != nil {
		// log
		questionDAO.logger.Error(err.Error())
		return nil, err
	}

	return questions.(*[]domain.Question), nil
}
