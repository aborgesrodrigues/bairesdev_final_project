package dao

import "../domain"

// UserDAO struct
type UserDAO struct{}

// Create func
func (UserDAO UserDAO) Create(user domain.User) (domain.User, error) {
	var d dao = dao{}

	createdUser, err := d.create(user)

	if err != nil {
		return createdUser.(domain.User), err
	}

	return createdUser.(domain.User), nil
}

// FindByID func
func (UserDAO UserDAO) FindByID(id int) (domain.User, error) {
	var d dao = dao{}

	user, err := d.findByID(domain.User{}, id)

	if err != nil {
		return user.(domain.User), err
	}

	return user.(domain.User), nil
}
