package track

import (
	"fmt"
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
	FetchFn        bool
	FetchOneFn     bool
	FetchTrackedFn bool
	UpdateFn       bool
	StoreFn        bool
}

func (repo *CurrencyRepoMock) Fetch() ([]*app.Currency, error) {
	repo.FetchFn = true
	return nil, nil
}
func (repo *CurrencyRepoMock) FetchOne(from, to string, lastNRates int) (*app.Currency, error) {
	repo.FetchOneFn = true
	return nil, nil
}
func (repo *CurrencyRepoMock) FetchTracked() ([]*app.Currency, error) {
	repo.FetchTrackedFn = true
	return []*app.Currency{
		&app.Currency{ID: 1, From: "USD", To: "SGD", Tracked: true, TrackedRev: false},
		&app.Currency{ID: 2, From: "USD", To: "IDR", Tracked: true, TrackedRev: true},
	}, nil
}
func (repo *CurrencyRepoMock) Update(uint, *app.Currency) (*app.Currency, error) {
	repo.UpdateFn = true
	return nil, nil
}
func (repo *CurrencyRepoMock) Store(*app.Currency) error {
	repo.StoreFn = true
	return nil
}

func (repo *RateRepoMock) Fetch() ([]*app.Rate, error) {
	repo.FetchFn = true
	return nil, nil
}
func (repo *RateRepoMock) FetchBetweenDate(uint, *time.Time, *time.Time) ([]*app.Rate, error) {
	repo.FetchBetweenDateFn = true
	dates := []time.Time{}
	for i := 0; i < 10; i++ {
		ti, _ := time.Parse("2006-01-02", fmt.Sprintf("2019-03-%0d", 10-i))
		dates = append(dates, ti)
	}
	return []*app.Rate{
		&app.Rate{ID: 1, Date: &dates[0], RateValue: 1.1, CurrencyID: 1},
		&app.Rate{ID: 2, Date: &dates[1], RateValue: 1.2, CurrencyID: 1},
		&app.Rate{ID: 3, Date: &dates[2], RateValue: 1.3, CurrencyID: 1},
		&app.Rate{ID: 4, Date: &dates[3], RateValue: 1.4, CurrencyID: 1},
		&app.Rate{ID: 5, Date: &dates[4], RateValue: 1.5, CurrencyID: 1},
		&app.Rate{ID: 6, Date: &dates[5], RateValue: 1.6, CurrencyID: 1},
		&app.Rate{ID: 7, Date: &dates[6], RateValue: 1.7, CurrencyID: 1},
		&app.Rate{ID: 8, Date: &dates[7], RateValue: 1.8, CurrencyID: 1},
		&app.Rate{ID: 9, Date: &dates[8], RateValue: 1.9, CurrencyID: 1},
		&app.Rate{ID: 10, Date: &dates[9], RateValue: 2.0, CurrencyID: 1},
	}, nil
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

func assertUint(t *testing.T, got, want uint) {
	t.Helper()
	if got != want {
		t.Errorf("got '%d' want '%d'", got, want)
	}
}
func TestTracks(t *testing.T) {
	rateRepo := &RateRepoMock{false, false, false}
	currencyRepo := &CurrencyRepoMock{false, false, false, false, false}
	trackService := CreateService(rateRepo, currencyRepo)

	trackService.Tracks("2019-03-14")
	assertBool(t, currencyRepo.FetchTrackedFn, true)
	assertBool(t, rateRepo.FetchBetweenDateFn, true)
}

func TestCalculateAvg(t *testing.T) {
	rateRepo := &RateRepoMock{false, false, false}
	currencyRepo := &CurrencyRepoMock{false, false, false, false, false}
	trackService := CreateService(rateRepo, currencyRepo)

	rates := []app.Rate{
		app.Rate{ID: 1, RateValue: 1},
		app.Rate{ID: 2, RateValue: 4},
	}
	got := trackService.calculateAvg(rates)
	want := float32(1+4) / float32(2)
	assertFloat(t, got, want)
}
