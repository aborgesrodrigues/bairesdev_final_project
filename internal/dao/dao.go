package dao

import (
	"log"

	"../domain"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DaoInterface interface
type DaoInterface interface {
	create(domain interface{}) interface{}
	findById(domain interface{}) interface{}
}

type dao struct {
	db *gorm.DB
}

func (d *dao) setConnection() {
	var err error
	d.db, err = gorm.Open(sqlite.Open("sqlite-database.db"), &gorm.Config{})

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

func (d *dao) findByID(domainStruct interface{}, id int) (interface{}, error) {
	var tx *gorm.DB
	d.setConnection()

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

func (d *dao) create(domainStruct interface{}) (interface{}, error) {
	var tx *gorm.DB
	d.setConnection()

	// domainConverted := d.getType(domainStruct)
	// tx = d.db.Create(&domainConverted)

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

	if tx.Error != nil {
		log.Println(tx.Error.Error())
		return domainStruct, tx.Error
	}

	return domainStruct, nil
}
