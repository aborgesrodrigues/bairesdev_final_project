package dao

import (
	"fmt"

	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
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
	// Entry log
	d.logger.Info("Called setConnection",
		zap.String("path_database", "./sqlite-database.db"),
	)

	var err error
	d.db, err = gorm.Open(sqlite.Open("./sqlite-database.db"), &gorm.Config{})

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

	tx = d.db.Create(domainStruct)

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
	}

	return domainStruct, nil
}

func (d *dao) delete(domainStruct interface{}, ID int) error {
	// Entry log
	d.logger.Info("Called delete",
		zap.String("domainStruct", fmt.Sprintf("%#v", domainStruct)),
	)

	var tx *gorm.DB

	d.db.Delete(domainStruct, ID)

	if tx.Error != nil {
		// log
		d.logger.Error(tx.Error.Error(),
			zap.String("domainStruct", fmt.Sprintf("%#v", domainStruct)),
		)
		return tx.Error
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
