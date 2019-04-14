package app

// TrackResponse is used for data structure tracking response
// It consist of attributes ID, From, To, RateValue, and Avg
type TrackResponse struct {
	ID        uint    `json:"id"`
	From      string  `json:"from"`
	To        string  `json:"to"`
	RateValue float32 `json:"rate"`
	Avg       float32 `json:"avg"`
}

// TrackRequest is used for data structure requesting track
// It consist of attributes From and To
type TrackRequest struct {
	From string `json:"from"`
	To   string `json:"to"`
}

// TrackService is an interface for track bussiness logic layer
type TrackService interface {
	Tracks(date string) ([]*TrackResponse, error)
	CreateTrack(*TrackRequest) error
	DeleteTrack(from, to string) error
}
