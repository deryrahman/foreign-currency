package currency

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func newDB() (sqlmock.Sqlmock, *gorm.DB) {
	db, mock, _ := sqlmock.New()
	gormDB, _ := gorm.Open("mysql", db)
	return mock, gormDB
}
