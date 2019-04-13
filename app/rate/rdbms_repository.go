package rate

import (
	"errors"
	"time"

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
	rates := []app.Rate{}
	repo.DB.Find(&rates)
	result := []*app.Rate{}
	for i := range rates {
		result = append(result, &rates[i])
	}
	return result, nil
}

// FetchBetweenDate is a method to fetch all rates within date
func (repo *RDBMSRepo) FetchBetweenDate(from *time.Time, to *time.Time) ([]*app.Rate, error) {
	rates := []app.Rate{}
	repo.DB.Find(&rates, "rates.date BETWEEN ? AND ?", from, to)
	result := []*app.Rate{}
	for i := range rates {
		result = append(result, &rates[i])
	}
	return result, nil
}

// Store is a method to store new daily rate into database
func (repo *RDBMSRepo) Store(rate *app.Rate) error {
	repo.DB.First(rate, "rates.date = ?", rate.Date)
	ok := repo.DB.NewRecord(rate)
	if !ok {
		return errors.New("rate exist")
	}
	repo.DB.Create(rate)
	return nil
}
