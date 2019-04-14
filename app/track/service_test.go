package track

import (
	"testing"
	"time"

	"github.com/deryrahman/foreign-currency/app"
)

type TrackRepoMock struct {
	FetchFn      bool
	StoreFn      bool
	DeleteByIDFn bool
}
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

func (repo *TrackRepoMock) Fetch() ([]*app.Track, error) {
	repo.FetchFn = true
	return nil, nil
}
func (repo *TrackRepoMock) Store(*app.Track) error {
	repo.StoreFn = true
	return nil
}
func (repo *TrackRepoMock) DeleteByID(uint) (*app.Track, error) {
	repo.DeleteByIDFn = true
	return nil, nil
}

func (repo *CurrencyRepoMock) Fetch() ([]*app.Currency, error) {
	repo.FetchFn = true
	return nil, nil
}
func (repo *CurrencyRepoMock) FetchOne(from, to string, lastNRates int) (*app.Currency, error) {
	repo.FetchOneFn = true
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
func TestTracks(t *testing.T) {
	trackRepo := &TrackRepoMock{false, false, false}
	rateRepo := &RateRepoMock{false, false, false}
	currencyRepo := &CurrencyRepoMock{false, false, false}
	trackService := CreateService(trackRepo, rateRepo, currencyRepo)

	trackService.Tracks("2019-03-14")
	assertBool(t, trackRepo.FetchFn, true)
	assertBool(t, rateRepo.FetchBetweenDateFn, true)
}
