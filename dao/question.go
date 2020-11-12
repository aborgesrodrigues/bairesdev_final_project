package dao

import "../domain"

// QuestionDAO struct
type QuestionDAO struct{}

// Create func
func (questionDAO QuestionDAO) Create(question domain.Question) (domain.Question, error) {
	var d dao = dao{}

	createdQuestion, err := d.create(question)

	if err != nil {
		return createdQuestion.(domain.Question), err
	}

	return createdQuestion.(domain.Question), nil
}

// FindByID func
func (questionDAO QuestionDAO) FindByID(id int) (domain.Question, error) {
	var d dao = dao{}

	question, err := d.findByID(&domain.Question{}, id)

	if err != nil {
		return question.(domain.Question), err
	}

	return question.(domain.Question), nil
}
