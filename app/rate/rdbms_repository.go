package rate

import (
	"time"

	"github.com/deryrahman/foreign-currency/app"
	"github.com/jinzhu/gorm"
)

// RDBMSRepo is a struct to wrap its DB
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
	repo.DB.Find(&rates).Order("rates.date DESC")
	result := []*app.Rate{}
	for i := range rates {
		result = append(result, &rates[i])
	}
	return result, nil
}

// FetchBetweenDate is a method to fetch all rates within date
func (repo *RDBMSRepo) FetchBetweenDate(currencyID uint, from *time.Time, to *time.Time) ([]*app.Rate, error) {
	rates := []app.Rate{}
	repo.DB.Find(&rates, "rates.currency_id = ? AND rates.date BETWEEN ? AND ?", currencyID, from, to).Order("rates.date DESC")
	result := []*app.Rate{}
	for i := range rates {
		result = append(result, &rates[i])
	}
	return result, nil
}

// Store is a method to store new daily rate into database
// If there's existing rate (same date), this method will throw error
func (repo *RDBMSRepo) Store(rate *app.Rate) error {
	repo.DB.First(rate, "rates.date = ? AND rates.currency_id = ?", rate.Date, rate.CurrencyID)
	ok := repo.DB.NewRecord(rate)
	if !ok {
		return app.ErrExist
	}
	repo.DB.Create(rate)
	return nil
}
