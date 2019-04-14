package track

import (
	"time"

	"github.com/deryrahman/foreign-currency/app"
)

// Service is struct for implementation of track service
type Service struct {
	RateRepo     app.RateRepository
	CurrencyRepo app.CurrencyRepository
	DateLayout   string
}

// CreateService is a constructor for create track service
func CreateService(rateRepo app.RateRepository, currencyRepo app.CurrencyRepository) *Service {
	return &Service{
		RateRepo:     rateRepo,
		CurrencyRepo: currencyRepo,
		DateLayout:   "2006-01-02",
	}
}

// Tracks is a method that receive date as a string with format YYYY-MM-DD
// and will return a TrackResponse object
// TrackResponse consist of ID, From, To, RateValue, and Avg from the last 7 days before date
// If it don't have sufficient data, throw an error
func (trackService *Service) Tracks(date string) ([]*app.TrackResponse, error) {
	dateEnd, err := time.Parse(trackService.DateLayout, date)
	if err != nil {
		return nil, err
	}
	dateBegin := dateEnd.AddDate(0, 0, -7)
	currencies, err := trackService.CurrencyRepo.FetchTracked()
	if err != nil {
		return nil, err
	}
	result := []*app.TrackResponse{}
	for i := range currencies {
		rates, _ := trackService.RateRepo.FetchBetweenDate(currencies[i].ID, &dateBegin, &dateEnd)
		from := currencies[i].From
		to := currencies[i].To
		// TODO calculate insufficient data
		rateValue := rates[0].RateValue
		avg := float32(0) // TODO: calculate avg given rate
		if currencies[i].TrackedRev {
			tmp := from
			from = to
			to = tmp
			rateValue = 1 / rateValue
			avg = 1 / avg
		}
		result = append(result, &app.TrackResponse{
			ID:        currencies[i].ID,
			From:      from,
			To:        to,
			RateValue: rateValue,
			Avg:       avg,
		})
	}
	return result, nil
}

func (trackService *Service) calculateAvg(rates []app.Rate) float32 {
	if len(rates) == 0 {
		return -1
	}
	result := float32(0)
	for i := range rates {
		result += rates[i].RateValue
	}
	return result / float32(len(rates))
}

// CreateTrack is a method that receive parameter "from" and "to" currency symbol
// If "to" is less than "from" lexicographically, then save it with revert true, false otherwise
func (trackService *Service) CreateTrack(from, to string) error {
	return nil
}

// DeleteTrack is a method to delete a track by it's id
func (trackService *Service) DeleteTrack(uint) error {
	return nil
}
