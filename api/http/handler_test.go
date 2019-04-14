package customhttp

import (
	"bytes"
	"encoding/json"
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
func (service *TrackServiceMock) CreateTrack(*app.TrackRequest) error {
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

func TestPostRates(t *testing.T) {
	rateService := &RateServiceMock{false, false}
	trackService := &TrackServiceMock{false, false, false}
	h := CreateHTTPHandler(rateService, trackService)

	s, _ := json.Marshal(&app.RateRequest{
		Date:      "2019-01-02",
		From:      "USD",
		To:        "SGD",
		RateValue: 0.5,
	})
	b := bytes.NewBuffer(s)

	request, _ := http.NewRequest(http.MethodPost, "/rates", b)
	response := httptest.NewRecorder()

	h.PostRates(response, request)
	assertBool(t, rateService.CreateRateFn, true)
}

func TestGetTracks(t *testing.T) {
	rateService := &RateServiceMock{false, false}
	trackService := &TrackServiceMock{false, false, false}
	h := CreateHTTPHandler(rateService, trackService)

	request, _ := http.NewRequest(http.MethodGet, "/tracks?date=2019-10-11", nil)
	response := httptest.NewRecorder()

	h.GetTracks(response, request)
	assertBool(t, trackService.TracksFn, true)
}

func TestPostTracks(t *testing.T) {
	rateService := &RateServiceMock{false, false}
	trackService := &TrackServiceMock{false, false, false}
	h := CreateHTTPHandler(rateService, trackService)
	s, _ := json.Marshal(&app.TrackRequest{
		From: "USD",
		To:   "SGD",
	})
	b := bytes.NewBuffer(s)

	request, _ := http.NewRequest(http.MethodPost, "/tracks", b)
	response := httptest.NewRecorder()

	h.PostTracks(response, request)
	assertBool(t, trackService.CreateTrackFn, true)
}

func TestDeleteTracks(t *testing.T) {
	rateService := &RateServiceMock{false, false}
	trackService := &TrackServiceMock{false, false, false}
	h := CreateHTTPHandler(rateService, trackService)

	request, _ := http.NewRequest(http.MethodDelete, "/tracks?from=USD&to=SGD", nil)
	response := httptest.NewRecorder()

	h.DeleteTracks(response, request)
	assertBool(t, trackService.DeleteTrackFn, true)
}
