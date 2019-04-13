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
	currencies := []app.Rate{}
	repo.DB.Find(&currencies)
	result := []*app.Rate{}
	for i := range currencies {
		result = append(result, &currencies[i])
	}
	return result, nil
}

// Store is a method to store new daily rate into database
func (repo *RDBMSRepo) Store(rate *app.Rate) error {
	repo.DB.Create(rate)
	return nil
}
