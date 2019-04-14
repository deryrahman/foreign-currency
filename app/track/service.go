package track

import "github.com/deryrahman/foreign-currency/app"

// Service is struct for implementation of track service
type Service struct {
	TrackRepo    app.TrackRepository
	RateRepo     app.RateRepository
	CurrencyRepo app.CurrencyRepository
}

// CreateService is a constructor for create track service
func CreateService(trackRepo app.TrackRepository, rateRepo app.RateRepository, currencyRepo app.CurrencyRepository) *Service {
	return &Service{
		TrackRepo:    trackRepo,
		RateRepo:     rateRepo,
		CurrencyRepo: currencyRepo,
	}
}

// Tracks is a method that receive date as a string with format YYYY-MM-DD
// and will return a TrackResponse object
// TrackResponse consist of ID, From, To, RateValue, and Avg from the last 7 days before date
// If it don't have sufficient data, throw an error
func (trackService *Service) Tracks(date string) ([]*app.TrackResponse, error) {
	return nil, nil
}

// CreateTrack is a method that receive parameter "from" and "to" currency symbol
// If "to" is less than "from" lexicographically, then save it with revert true, false otherwise
func (trackService *Service) CreateTrack(from, to string) error {
	return nil
}

// DeleteTrack is a method to delete a track by it's id
func (trackService *Service) DeleteTrack(uint) (*app.Track, error) {
	return nil, nil
}
