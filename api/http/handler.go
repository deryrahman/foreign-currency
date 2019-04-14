package customhttp

import (
	"encoding/json"
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
	queries := r.URL.Query()
	from := queries.Get("from")
	to := queries.Get("to")
	currencyResponse, err := h.RateService.CurrencyRates(from, to, 7)
	if err != nil {
		if err == app.ErrNotFound {
			w.WriteHeader(http.StatusNotFound)
		} else if err == app.ErrExist {
			w.WriteHeader(http.StatusConflict)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		errorResponse := app.ErrorResponse{ErrMsg: err.Error()}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}
	json.NewEncoder(w).Encode(currencyResponse)
}

// PostRates is a method to create daily rates
// It will read and parse request body as json and marshaling into RateRequest model
func (h *HTTPHandler) PostRates(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	rateRequest := app.RateRequest{}
	err := decoder.Decode(&rateRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorResponse := app.ErrorResponse{ErrMsg: err.Error()}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}
	err = h.RateService.CreateRate(&rateRequest)
	if err != nil {
		if err == app.ErrNotFound {
			w.WriteHeader(http.StatusNotFound)
		} else if err == app.ErrExist {
			w.WriteHeader(http.StatusConflict)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		errorResponse := app.ErrorResponse{ErrMsg: err.Error()}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}
}

// GetTracks is a method to get all tracks
// It receive query "date" with format YYYY-MM-DD
func (h *HTTPHandler) GetTracks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	queries := r.URL.Query()
	date := queries.Get("date")
	trackResponse, err := h.TrackService.Tracks(date)
	if err != nil {
		if err == app.ErrNotFound {
			w.WriteHeader(http.StatusNotFound)
		} else if err == app.ErrExist {
			w.WriteHeader(http.StatusConflict)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		errorResponse := app.ErrorResponse{ErrMsg: err.Error()}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}
	json.NewEncoder(w).Encode(trackResponse)
}

// PostTracks is a method to invoke currency rate to be tracked
// It will read request body as json. Json parameter are "from" and "to"
func (h *HTTPHandler) PostTracks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	trackRequest := app.TrackRequest{}
	err := decoder.Decode(&trackRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorResponse := app.ErrorResponse{ErrMsg: err.Error()}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}
	err = h.TrackService.CreateTrack(&trackRequest)
	if err != nil {
		if err == app.ErrNotFound {
			w.WriteHeader(http.StatusNotFound)
		} else if err == app.ErrExist {
			w.WriteHeader(http.StatusConflict)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		errorResponse := app.ErrorResponse{ErrMsg: err.Error()}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}
}

// DeleteTracks is a method to remove currency rate to be tracked
// It receive query "from" and "to"
func (h *HTTPHandler) DeleteTracks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	queries := r.URL.Query()
	from := queries.Get("from")
	to := queries.Get("to")
	err := h.TrackService.DeleteTrack(from, to)
	if err != nil {
		if err == app.ErrNotFound {
			w.WriteHeader(http.StatusNotFound)
		} else if err == app.ErrExist {
			w.WriteHeader(http.StatusConflict)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		errorResponse := app.ErrorResponse{ErrMsg: err.Error()}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}
}
