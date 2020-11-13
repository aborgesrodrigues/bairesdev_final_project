package dao

import (
	"log"

	"../domain"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DaoInterface interface
// type DaoInterface interface {
// 	create(domain interface{}) interface{}
// 	findById(domain interface{}) interface{}
// }

type dao struct {
	db *gorm.DB
}

// newDAO Constructor of dao struct
func newDAO() *dao {
	d := &dao{}
	d.setConnection()
	return d
}

func (d *dao) setConnection() {
	var err error
	d.db, err = gorm.Open(sqlite.Open("../sqlite-database.db"), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}
}

func (d *dao) getType(domainStruct interface{}) interface{} {
	switch domainStruct.(type) {
	case domain.User:
		//var user domain.User
		return domainStruct.(domain.User)
		//tx = d.db.Create(&user)
	case domain.Question:
		return domainStruct.(domain.Question)
		//tx = d.db.Create(&question)
	}
	return nil
}

func (d *dao) create(domainStruct interface{}) (interface{}, error) {
	var tx *gorm.DB

	switch domainStruct.(type) {
	case domain.User:
		user := domainStruct.(domain.User)
		tx = d.db.Create(&user)
		domainStruct = user
	case domain.Question:
		question := domainStruct.(domain.Question)
		tx = d.db.Create(&question)
		domainStruct = question
	}

	//tx = d.db.Create(domainStruct)

	if tx.Error != nil {
		log.Println(tx.Error.Error())
		return domainStruct, tx.Error
	}

	return domainStruct, nil
}

func (d *dao) update(domainStruct interface{}) (interface{}, error) {
	var tx *gorm.DB

	switch domainStruct.(type) {
	case domain.User:
		user := domainStruct.(domain.User)
		tx = d.db.Updates(&user)
		domainStruct = user
	case domain.Question:
		question := domainStruct.(domain.Question)
		tx = d.db.Updates(&question)
		domainStruct = question
	}

	if tx.Error != nil {
		log.Println(tx.Error.Error())
		return domainStruct, tx.Error
	}

	return domainStruct, nil
}

func (d *dao) delete(domainStruct interface{}, ID int) error {
	var tx *gorm.DB

	switch domainStruct.(type) {
	case *domain.User:
		user := domainStruct.(*domain.User)
		tx = d.db.Delete(&user, ID)
	case *domain.Question:
		question := domainStruct.(*domain.Question)
		tx = d.db.Delete(&question, ID)
	}

	if tx.Error != nil {
		log.Println(tx.Error.Error())
		return tx.Error
	}

	return nil
}

func (d *dao) findByID(domainStruct interface{}, id int) (interface{}, error) {
	var tx *gorm.DB
	//d.setConnection()

	switch domainStruct.(type) {
	case domain.User:
		user := domainStruct.(domain.User)
		tx = d.db.First(&user, id)
		domainStruct = user
	case domain.Question:
		question := domainStruct.(domain.Question)
		tx = d.db.First(&question, id)
		domainStruct = question
	}

	//tx := d.db.First(&domain, id)

	if tx.Error != nil {
		log.Println(tx.Error.Error())
		return domainStruct, tx.Error
	}

	return domainStruct, nil
}

func (d *dao) getAll(domainStruct interface{}) (interface{}, error) {
	var tx *gorm.DB
	//d.setConnection()

	switch domainStruct.(type) {
	case []domain.User:
		users := domainStruct.([]domain.User)
		tx = d.db.Find(&users)
		domainStruct = users
	case []domain.Question:
		questions := domainStruct.([]domain.Question)
		tx = d.db.Find(&questions)
		domainStruct = questions
	}

	if tx.Error != nil {
		log.Println(tx.Error.Error())
		return domainStruct, tx.Error
	}

	return domainStruct, nil
}
