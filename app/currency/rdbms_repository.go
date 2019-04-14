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
	for i := range currencies {
		result[i] = &currencies[i]
	}
	return result, nil
}

// FetchTracked is a method to fetch all currency with tracked checked,
// either TrackedRev or Tracked
func (repo *RDBMSRepo) FetchTracked() ([]*app.Currency, error) {
	currencies := []app.Currency{}
	repo.DB.Find(&currencies, "currencies.tracked = ? OR currencies.tracked_rev = ?", true, true)
	result := make([]*app.Currency, len(currencies))
	for i := range currencies {
		result[i] = &currencies[i]
	}
	return result, nil
}

// FetchOne is a method to fetch one currency pair
// It should pass parameter "from", "to", and "lastNRates"
// "from" is always less than "to" lexicographically
// If lastNRates is negative, return all Rates, get latest N rates otherwise
func (repo *RDBMSRepo) FetchOne(from, to string, lastNRates int) (*app.Currency, error) {
	currency := app.Currency{}
	rates := []app.Rate{}
	if lastNRates == 0 {
		repo.DB.First(&currency, "currencies.from = ? AND currencies.to = ?", from, to)
	} else if lastNRates < 0 {
		repo.DB.First(&currency, "currencies.from = ? AND currencies.to = ?", from, to).Order("rates.date DESC").Related(&rates)
	} else {
		repo.DB.First(&currency, "currencies.from = ? AND currencies.to = ?", from, to).Order("rates.date DESC").Limit(lastNRates).Related(&rates)
	}
	if currency.ID == 0 {
		return nil, app.ErrNotFound
	}
	currency.Rates = rates
	return &currency, nil
}

// Update is a method to toggle tracked and trackedRev currency
// It will throw an error if currency not found
func (repo *RDBMSRepo) Update(id uint, currencyNew *app.Currency) (*app.Currency, error) {
	currency := &app.Currency{}
	repo.DB.First(currency, "id = ?", id)
	if currency.ID == 0 {
		return nil, app.ErrNotFound
	}
	currencyNew.ID = id
	repo.DB.Model(currency).Updates(app.Currency{
		Tracked:    currencyNew.Tracked,
		TrackedRev: currencyNew.TrackedRev,
	})
	return currencyNew, nil
}

// Store is a method to store new currency into database
// If there's existing currency (same from and date), this method will throw error
func (repo *RDBMSRepo) Store(currency *app.Currency) error {
	repo.DB.First(currency, "currencies.from = ? AND currencies.to = ?", currency.From, currency.To)
	ok := repo.DB.NewRecord(currency)
	if !ok {
		return app.ErrExist
	}
	repo.DB.Create(currency)
	return nil
}
