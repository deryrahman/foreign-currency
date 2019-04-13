package currency

import (
	"github.com/deryrahman/foreign-currency/app"
	"github.com/jinzhu/gorm"
)

// RDBMSRepo is a struct to wrap its DB
type RDBMSRepo struct {
	DB *gorm.DB
}

// CreateRDBMSRepo is used to create new Mysql repository
func CreateRDBMSRepo(db *gorm.DB) *RDBMSRepo {
	return &RDBMSRepo{db}
}

// Fetch is a method to fetch all currency that match with query
// It should return ErrNotFound if currency didn't found
func (repo *RDBMSRepo) Fetch() ([]*app.Currency, error) {
	currencies := []app.Currency{}
	repo.DB.Find(&currencies)
	result := make([]*app.Currency, len(currencies))
	for i, v := range currencies {
		result[i] = &v
	}
	return result, nil
}

// FetchOne is a method to fetch one currency pair
// It should pass parameter "from", "to", and "lastNRates"
// "from" is always less than "to" lexicographically
// If lastNRates is negative, return all Rates, get latest N rates otherwise
// It should return ErrNotFound if currency didn't found
func (repo *RDBMSRepo) FetchOne(from, to string, lastNRates int) (*app.Currency, error) {
	return nil, nil
}

// Store is a method to store new currency into database
func (repo *RDBMSRepo) Store(currency *app.Currency) error {
	return nil
}
