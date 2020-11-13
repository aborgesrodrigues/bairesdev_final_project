package dao

import "../domain"

// UserDAO struct
type UserDAO struct {
	d *dao
}

// NewUserDAO Constructor of UserDAO struct
func NewUserDAO() *UserDAO {
	userDAO := &UserDAO{}
	userDAO.d = NewDAO()
	return userDAO
}

// Create func
func (userDAO UserDAO) Create(user domain.User) (domain.User, error) {
	createdUser, err := userDAO.d.create(user)

	if err != nil {
		return createdUser.(domain.User), err
	}

	return createdUser.(domain.User), nil
}

// FindByID func
func (userDAO UserDAO) FindByID(id int) (domain.User, error) {
	user, err := userDAO.d.findByID(domain.User{}, id)

	if err != nil {
		return user.(domain.User), err
	}

	return user.(domain.User), nil
}

// GetAll func
func (userDAO UserDAO) GetAll() ([]domain.User, error) {
	users, err := userDAO.d.getAll(make([]domain.User, 0))

	if err != nil {
		return nil, err
	}

	return users.([]domain.User), nil
}
