package rate

import (
	"testing"
	"time"

	"github.com/deryrahman/foreign-currency/app"
)

type RateRepoMock struct {
	FetchFn            bool
	FetchBetweenDateFn bool
	StoreFn            bool
}
type CurrencyRepoMock struct {
	FetchFn    bool
	FetchOneFn bool
	StoreFn    bool
}

func (repo *CurrencyRepoMock) Fetch() ([]*app.Currency, error) {
	repo.FetchFn = true
	return nil, nil
}
func (repo *CurrencyRepoMock) FetchOne(from, to string, lastNRates int) (*app.Currency, error) {
	repo.FetchOneFn = true
	return &app.Currency{
		ID:    1,
		From:  "SGD",
		To:    "USD",
		Rates: []app.Rate{},
	}, nil
}
func (repo *CurrencyRepoMock) Store(*app.Currency) error {
	repo.StoreFn = true
	return nil
}

func (repo *RateRepoMock) Fetch() ([]*app.Rate, error) {
	repo.FetchFn = true
	return nil, nil
}
func (repo *RateRepoMock) FetchBetweenDate(*time.Time, *time.Time) ([]*app.Rate, error) {
	repo.FetchBetweenDateFn = true
	return nil, nil
}
func (repo *RateRepoMock) Store(*app.Rate) error {
	repo.StoreFn = true
	return nil
}
func TestCurrencyRates(t *testing.T) {
	rateRepo := &RateRepoMock{false, false, false}
	currencyRepo := &CurrencyRepoMock{false, false, false}
	rateService := CreateService(rateRepo, currencyRepo)

	rateService.CurrencyRates("USD", "SGD", 7)
	got := currencyRepo.FetchOneFn
	want := true
	if got != want {
		t.Errorf("got '%v' want '%v'", got, want)
	}
}
