package rate

import (
	"github.com/deryrahman/foreign-currency/app"
	"github.com/jinzhu/gorm"
)

type RDBMSRepo struct {
	DB *gorm.DB
}

// CreateRDBMSRepo is used to create new RDBMS repository
func CreateRDBMSRepo(db *gorm.DB) *RDBMSRepo {
	return &RDBMSRepo{db}
}

// Fetch is a method to fetch all rates
func (repo *RDBMSRepo) Fetch() ([]*app.Rate, error) {
	return nil, nil
}

// Store is a method to store new daily rate into database
func (repo *RDBMSRepo) Store(rate *app.Rate) error {
	return nil
}
