package customhttp

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/deryrahman/foreign-currency/app"
)

type RateServiceMock struct {
	CurrencyRatesFn bool
	CreateRateFn    bool
}

type TrackServiceMock struct {
	TracksFn      bool
	CreateTrackFn bool
	DeleteTrackFn bool
}

func (service *RateServiceMock) CurrencyRates(from, to string, lastNRates int) (*app.CurrencyResponse, error) {
	service.CurrencyRatesFn = true
	return nil, nil
}
func (service *RateServiceMock) CreateRate(*app.RateRequest) error {
	service.CreateRateFn = true
	return nil
}

func (service *TrackServiceMock) Tracks(date string) ([]*app.TrackResponse, error) {
	service.TracksFn = true
	return nil, nil
}
func (service *TrackServiceMock) CreateTrack(from, to string) error {
	service.CreateTrackFn = true
	return nil
}
func (service *TrackServiceMock) DeleteTrack(from, to string) error {
	service.DeleteTrackFn = true
	return nil
}

func assertBool(t *testing.T, got, want bool) {
	t.Helper()
	if got != want {
		t.Errorf("got '%v' want '%v'", got, want)
	}
}
func TestGetRates(t *testing.T) {
	rateService := &RateServiceMock{false, false}
	trackService := &TrackServiceMock{false, false, false}
	h := CreateHTTPHandler(rateService, trackService)

	request, _ := http.NewRequest(http.MethodGet, "/rates", nil)
	response := httptest.NewRecorder()

	h.GetRates(response, request)
	assertBool(t, rateService.CurrencyRatesFn, true)
}
