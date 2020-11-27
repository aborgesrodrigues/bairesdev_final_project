package dao

import (
	"errors"
	"fmt"
	"os"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type dao struct {
	db     *gorm.DB
	logger *zap.Logger
}

// newDAO Constructor of dao struct
func newDAO(logger *zap.Logger) *dao {
	d := &dao{logger: logger}
	d.setConnection()
	return d
}

func (d *dao) setConnection() {
	var host = os.Getenv("DATABASE_HOST")
	var user = os.Getenv("DATABASE_USER")
	var password = os.Getenv("DATABASE_PASSWORD")
	var dbName = os.Getenv("DATABASE_NAME")
	var port = os.Getenv("DATABASE_PORT")

	dsn := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%s", host, user, dbName, password, port)

	// Entry log
	d.logger.Info("Called setConnection",
		zap.String("dsn", dsn),
	)

	var err error
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	d.db = db

	if err != nil {
		// log
		d.logger.Error(err.Error())
	}
}

func (d *dao) create(domainStruct interface{}) (interface{}, error) {
	// Entry log
	d.logger.Info("Called create",
		zap.String("domainStruct", fmt.Sprintf("%#v", domainStruct)),
	)

	var tx *gorm.DB

	tx = d.db.Omit("User").Create(domainStruct)

	if tx.Error != nil {
		// log
		d.logger.Error(tx.Error.Error(),
			zap.String("domainStruct", fmt.Sprintf("%#v", domainStruct)),
		)
		return domainStruct, tx.Error
	}

	return domainStruct, nil
}

func (d *dao) update(domainStruct interface{}) (interface{}, error) {
	// Entry log
	d.logger.Info("Called update",
		zap.String("domainStruct", fmt.Sprintf("%#v", domainStruct)),
	)

	var tx *gorm.DB

	tx = d.db.Updates(domainStruct)

	if tx.Error != nil {
		// log
		d.logger.Error(tx.Error.Error(),
			zap.String("domainStruct", fmt.Sprintf("%#v", domainStruct)),
		)
		return domainStruct, tx.Error
	} else if tx.RowsAffected == 0 {
		message := "ID not found"
		// log
		d.logger.Error(message,
			zap.String("domainStruct", fmt.Sprintf("%#v", domainStruct)),
		)
		return domainStruct, errors.New(message)
	}

	return domainStruct, nil
}

func (d *dao) delete(domainStruct interface{}, ID int) error {
	// Entry log
	d.logger.Info("Called delete",
		zap.String("domainStruct", fmt.Sprintf("%#v", domainStruct)),
	)

	var tx *gorm.DB

	tx = d.db.Delete(domainStruct, ID)

	if tx.Error != nil {
		// log
		d.logger.Error(tx.Error.Error(),
			zap.String("domainStruct", fmt.Sprintf("%#v", domainStruct)),
		)
		return tx.Error
	} else if tx.RowsAffected == 0 {
		// log
		message := "ID not found"
		d.logger.Error(message,
			zap.String("domainStruct", fmt.Sprintf("%#v", domainStruct)),
		)
		return errors.New(message)
	}

	return nil
}

func (d *dao) findByID(domainStruct interface{}, id int) (interface{}, error) {
	// Entry log
	d.logger.Info("Called findByID",
		zap.String("domainStruct", fmt.Sprintf("%#v", domainStruct)),
	)
	var tx *gorm.DB

	tx = d.db.First(domainStruct, id)

	if tx.Error != nil {
		// log
		d.logger.Error(tx.Error.Error(),
			zap.String("domainStruct", fmt.Sprintf("%#v", domainStruct)),
		)
		return domainStruct, tx.Error
	}

	return domainStruct, nil
}

func (d *dao) getAll(domainStruct interface{}) (interface{}, error) {
	// Entry log
	d.logger.Info("Called getAll",
		zap.String("domainStruct", fmt.Sprintf("%#v", domainStruct)),
	)

	var tx *gorm.DB

	tx = d.db.Find(domainStruct)

	if tx.Error != nil {
		// log
		d.logger.Error(tx.Error.Error(),
			zap.String("domainStruct", fmt.Sprintf("%#v", domainStruct)),
		)
		return domainStruct, tx.Error
	}

	return domainStruct, nil
}
