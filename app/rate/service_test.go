package rate

import (
	"errors"
	"testing"
	"time"

	"github.com/deryrahman/foreign-currency/app"
)

type RateRepoMock struct {
	Fail               bool
	FetchFn            bool
	FetchBetweenDateFn bool
	StoreFn            bool
}
type CurrencyRepoMock struct {
	Fail       bool
	FetchFn    bool
	FetchOneFn bool
	StoreFn    bool
}

func (repo *CurrencyRepoMock) Fetch() ([]*app.Currency, error) {
	repo.FetchFn = true
	return nil, nil
}
func (repo *CurrencyRepoMock) FetchOne(from, to string, lastNRates int) (*app.Currency, error) {
	if repo.Fail {
		return nil, errors.New("")
	}
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

func assertBool(t *testing.T, got, want bool) {
	t.Helper()
	if got != want {
		t.Errorf("got '%v' want '%v'", got, want)
	}
}

func assertFloat(t *testing.T, got, want float32) {
	t.Helper()
	if got != want {
		t.Errorf("got '%f' want '%f'", got, want)
	}
}
func TestCurrencyRates(t *testing.T) {
	rateRepo := &RateRepoMock{false, false, false, false}
	currencyRepo := &CurrencyRepoMock{false, false, false, false}
	rateService := CreateService(rateRepo, currencyRepo)

	rateService.CurrencyRates("USD", "SGD", 7)
	assertBool(t, currencyRepo.FetchOneFn, true)
}

func TestCurrencyRates_fail(t *testing.T) {
	rateRepo := &RateRepoMock{true, false, false, false}
	currencyRepo := &CurrencyRepoMock{true, false, false, false}
	rateService := CreateService(rateRepo, currencyRepo)

	currencyResponse, err := rateService.CurrencyRates("USD", "SGD", 7)
	assertBool(t, currencyRepo.FetchOneFn, false)
	if currencyResponse != nil {
		t.Errorf("should nil, got '%v'", currencyResponse)
	}
	if err == nil {
		t.Errorf("wanted an error")
	}
}

func TestCalculateAvg(t *testing.T) {
	rateRepo := &RateRepoMock{false, false, false, false}
	currencyRepo := &CurrencyRepoMock{false, false, false, false}
	rateService := CreateService(rateRepo, currencyRepo)
	rates := []app.Rate{
		app.Rate{ID: 1, RateValue: 1},
		app.Rate{ID: 2, RateValue: 4},
	}
	got := rateService.calculateAvg(rates)
	want := float32(1+4) / float32(2)
	assertFloat(t, got, want)
}

func TestCalculateAvg_zeroRates(t *testing.T) {
	rateRepo := &RateRepoMock{false, false, false, false}
	currencyRepo := &CurrencyRepoMock{false, false, false, false}
	rateService := CreateService(rateRepo, currencyRepo)
	rates := []app.Rate{}
	got := rateService.calculateAvg(rates)
	want := float32(-1)
	assertFloat(t, got, want)
}

func TestCalculateVar(t *testing.T) {
	rateRepo := &RateRepoMock{false, false, false, false}
	currencyRepo := &CurrencyRepoMock{false, false, false, false}
	rateService := CreateService(rateRepo, currencyRepo)
	rates := []app.Rate{
		app.Rate{ID: 1, RateValue: 1},
		app.Rate{ID: 2, RateValue: 4},
	}
	got := rateService.calculateVar(rates)
	want := float32(4 - 1)
	assertFloat(t, got, want)
}

func TestCalculateVar_zeroRates(t *testing.T) {
	rateRepo := &RateRepoMock{false, false, false, false}
	currencyRepo := &CurrencyRepoMock{false, false, false, false}
	rateService := CreateService(rateRepo, currencyRepo)
	rates := []app.Rate{}
	got := rateService.calculateVar(rates)
	want := float32(-1)
	assertFloat(t, got, want)
}

func TestCreateRate_dontHaveCurrencyBefore(t *testing.T) {
	rateRepo := &RateRepoMock{false, false, false, false}
	currencyRepo := &CurrencyRepoMock{false, false, false, false}
	rateService := CreateService(rateRepo, currencyRepo)
	ti := time.Now()
	rateReq := app.RateRequest{
		Date:      &ti,
		From:      "USD",
		To:        "SGD",
		RateValue: 0.8,
	}
	rateService.CreateRate(&rateReq)
	assertBool(t, currencyRepo.FetchOneFn, true)
	assertBool(t, rateRepo.StoreFn, true)
}
