package customhttp

import (
	"net/http"

	"github.com/deryrahman/foreign-currency/app"
)

// HTTPHandler is a wrapper of rate service and track service
type HTTPHandler struct {
	RateService  app.RateService
	TrackService app.TrackService
}

// CreateHTTPHandler is a constructor to create HTTPHandler object
func CreateHTTPHandler(rateService app.RateService, trackService app.TrackService) *HTTPHandler {
	return &HTTPHandler{
		RateService:  rateService,
		TrackService: trackService,
	}
}

// GetRates is a method to get rates
// It receive query "from" and "to", and will retrive corresponding data of rates details
func (h *HTTPHandler) GetRates(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

// PostRates is a method to create daily rates
// It will read and parse request body as json and marshaling into RateRequest model
func (h *HTTPHandler) PostRates(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

// GetTracks is a method to get all tracks
// It receive query "date" with format YYYY-MM-DD
func (h *HTTPHandler) GetTracks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

// PostTracks is a method to invoke currency rate to be tracked
// It will read request body as json. Json parameter are "from" and "to"
func (h *HTTPHandler) PostTracks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

// DeleteTracks is a method to remove currency rate to be tracked
// It receive query "from" and "to"
func (h *HTTPHandler) DeleteTracks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}
