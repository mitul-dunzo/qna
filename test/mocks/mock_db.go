package mocks

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
)

func GetMockDB() (*gorm.DB, sqlmock.Sqlmock) {

	db, mock, err := sqlmock.New()
	if err != nil {
		panic("Panic")
	}
	txn, err := gorm.Open("postgres", db)
	if err != nil {
		panic("Panic")
	}

	err = txn.LogMode(true).Error
	if err != nil {
		panic("Panic")
	}
	return txn, mock
}
