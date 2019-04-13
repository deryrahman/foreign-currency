package app

// Track is a struct for tracking
// It consist of attributes ID, Currency, CurrencyID, and Revert
// Revert is just a special boolean if currency source and destination
// is not in lexicographical order
type Track struct {
	ID         uint     `json:"id"`
	Currency   Currency `json:"currency"`
	CurrencyID uint     `json:"currency_id"`
	Revert     bool     `json:"revert"`
}

// TrackResponse is used for data structure tracking response
// It consist of attributes ID, From, To, RateValue, and Avg
type TrackResponse struct {
	ID        uint    `json:"id"`
	From      string  `json:"from"`
	To        string  `json:"to"`
	RateValue float32 `json:"rate"`
	Avg       float32 `json:"avg"`
}

// TrackRepository is an interface for track repository layer
type TrackRepository interface {
	Fetch() ([]*Track, error)
	Store(*Track) error
	DeleteById(uint) (*Track, error)
}

// TrackService is an interface for track bussiness logic layer
type TrackService interface {
	Tracks(date string) ([]*TrackResponse, error)
	CreateTrack(from, to string) error
	DeleteTrack(uint) (*Track, error)
}
