package dao

import (
	"log"

	"../domain"
)

// QuestionDAO struct
type QuestionDAO struct {
	d *dao
}

//QuestionDAOInterface interface
type QuestionDAOInterface interface {
	Create(domain.Question) (domain.Question, error)
	Update(domain.Question) (domain.Question, error)
	Delete(int) error
	FindByID(int) (domain.Question, error)
	FindByUser(int) ([]domain.Question, error)
	GetAll() ([]domain.Question, error)
}

// NewQuestionDAO Constructor of QuestionDAO struct
func NewQuestionDAO() *QuestionDAO {
	questionDAO := &QuestionDAO{}
	questionDAO.d = newDAO()
	return questionDAO
}

// Create func
func (questionDAO QuestionDAO) Create(question domain.Question) (domain.Question, error) {
	createdQuestion, err := questionDAO.d.create(question)

	if err != nil {
		return createdQuestion.(domain.Question), err
	}

	return createdQuestion.(domain.Question), nil
}

// Update func
func (questionDAO QuestionDAO) Update(question domain.Question) (domain.Question, error) {
	updatedQuestion, err := questionDAO.d.update(question)

	if err != nil {
		return updatedQuestion.(domain.Question), err
	}

	return updatedQuestion.(domain.Question), nil
}

// Delete func
func (questionDAO QuestionDAO) Delete(ID int) error {
	err := questionDAO.d.delete(&domain.Question{}, ID)

	if err != nil {
		return err
	}

	return nil
}

// FindByID func
func (questionDAO QuestionDAO) FindByID(id int) (domain.Question, error) {
	question, err := questionDAO.d.findByID(&domain.Question{}, id)

	if err != nil {
		return question.(domain.Question), err
	}

	return question.(domain.Question), nil
}

// FindByUser func
func (questionDAO QuestionDAO) FindByUser(userID int) ([]domain.Question, error) {
	//var user domain.User
	//userReturned, err := d.findByID(&domain.User{}, userID)
	// /user := userReturned.(domain.User)

	//if err != nil {
	//	log.Println(err.Error())
	//	return []domain.Question{}, err
	//}

	var questions = make([]domain.Question, 0)
	tx := questionDAO.d.db.Where(&domain.Question{UserID: userID}).Find(&questions)

	if tx.Error != nil {
		log.Println(tx.Error.Error())
		return make([]domain.Question, 0), tx.Error
	}

	return questions, nil
}

// GetAll func
func (questionDAO QuestionDAO) GetAll() ([]domain.Question, error) {
	questions, err := questionDAO.d.getAll(make([]domain.Question, 0))

	if err != nil {
		return make([]domain.Question, 0), err
	}

	return questions.([]domain.Question), nil
}
