package app

import "time"

// Rate is a struct for daily rate each currency
// It consist of attributes ID, Date, RateValue, and CurrencyID
type Rate struct {
	ID         uint       `json:"id"`
	Date       *time.Time `json:"date"`
	RateValue  float32    `json:"rate"`
	CurrencyID uint       `json:"currency_id"`
}

// RateRequest is a used for data structure rate request
// It consist of attributes Date, From, To, and RateValue
type RateRequest struct {
	Date      *time.Time `json:"date"`
	From      string     `json:"from"`
	To        string     `json:"to"`
	RateValue float32    `json:"rate"`
}

// RateRepository is an interface for rate repository layer
type RateRepository interface {
	Fetch() ([]*Rate, error)
	FetchBetweenDate(*time.Time, *time.Time) ([]*Rate, error)
	Store(*Rate) error
}

// RateService is an interface for rate bussiness logic layer
type RateService interface {
	CurrencyRates(from, to string, lastNRates int) (*CurrencyResponse, error)
	CreateRate(*RateRequest) error
}
