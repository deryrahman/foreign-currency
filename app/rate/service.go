package rate

import (
	"time"

	"github.com/deryrahman/foreign-currency/app"
)

// Service is struct for implementation of rate service
type Service struct {
	RateRepo     app.RateRepository
	CurrencyRepo app.CurrencyRepository
	DateLayout   string
}

// CreateService is a constructor for create rate service
func CreateService(rateRepo app.RateRepository, currencyRepo app.CurrencyRepository) *Service {
	return &Service{
		RateRepo:     rateRepo,
		CurrencyRepo: currencyRepo,
		DateLayout:   "2006-01-02",
	}
}

// CurrencyRates is a method to get currency with their rates
// It has parameter "from", "to", and "lastNRates"
// lastNRates < 0 will retrieve all rates, lastNRates >= 0 will retrieve recent top lastNRates
// Before call method fetch on currency repo, "from" should less than "to" lexicographically
func (rateService *Service) CurrencyRates(from, to string, lastNRates int) (*app.CurrencyResponse, error) {
	revert := from > to
	if revert {
		tmp := from
		from = to
		to = tmp
	}
	currency, err := rateService.CurrencyRepo.FetchOne(from, to, lastNRates)
	if err != nil {
		return nil, err
	}
	if revert {
		for i := range currency.Rates {
			currency.Rates[i].RateValue = 1 / currency.Rates[i].RateValue
		}
		tmp := currency.From
		currency.From = currency.To
		currency.To = tmp
	}
	currencyResponse := &app.CurrencyResponse{
		ID:    currency.ID,
		From:  currency.From,
		To:    currency.To,
		Avg:   rateService.calculateAvg(currency.Rates),
		Var:   rateService.calculateVar(currency.Rates),
		Rates: currency.Rates,
	}
	return currencyResponse, nil
}

func (rateService *Service) calculateAvg(rates []app.Rate) float32 {
	if len(rates) == 0 {
		return -1
	}
	result := float32(0)
	for i := range rates {
		result += rates[i].RateValue
	}
	return result / float32(len(rates))
}

func (rateService *Service) calculateVar(rates []app.Rate) float32 {
	if len(rates) == 0 {
		return -1
	}
	max := float32(rates[0].RateValue)
	min := float32(rates[0].RateValue)
	for _, v := range rates[1:] {
		if v.RateValue > max {
			max = v.RateValue
		}
		if v.RateValue < min {
			min = v.RateValue
		}
	}
	return max - min
}

// CreateRate is a method to create daily rate
// If currency doesn't exist yet, then create one using currency repo
// create currency, must have "from" less than "to" lexicographically
func (rateService *Service) CreateRate(rateRequest *app.RateRequest) error {
	from := rateRequest.From
	to := rateRequest.To
	rateValue := rateRequest.RateValue
	revert := from > to
	if revert {
		tmp := from
		from = to
		to = tmp
		rateValue = 1 / rateValue
	}
	currency, err := rateService.CurrencyRepo.FetchOne(from, to, 0)
	if err == app.ErrNotFound {
		currency = &app.Currency{
			From:       from,
			To:         to,
			Tracked:    false,
			TrackedRev: false,
		}
		rateService.CurrencyRepo.Store(currency)
	}
	ti, err := time.Parse(rateService.DateLayout, rateRequest.Date)
	if err != nil {
		return err
	}
	rate := app.Rate{
		Date:       &ti,
		CurrencyID: currency.ID,
		RateValue:  rateValue,
	}
	return rateService.RateRepo.Store(&rate)
}
