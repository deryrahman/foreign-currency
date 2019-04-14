package app

// Currency is a struct for currency
// It consist of attributes ID, From, To, Rates, Tracked, and TrackedRev
// Tracked is just a special boolean if currency pair is tracked
// TrackedRev is just a special boolean if currency pair is tracked without lexicographical order
type Currency struct {
	ID         uint   `json:"id"`
	From       string `json:"from"`
	To         string `json:"to"`
	Rates      []Rate `json:"rates"`
	Tracked    bool   `json:"tracked"`
	TrackedRev bool   `json:"tracked_rev"`
}

// CurrencyResponse is used for data structure currency response
// It consist of attributes ID, From, To, Avg, Var, and Rates
type CurrencyResponse struct {
	ID    uint    `json:"id"`
	From  string  `json:"from"`
	To    string  `json:"to"`
	Avg   float32 `json:"avg"`
	Var   float32 `json:"var"`
	Rates []Rate  `json:"rates"`
}

// CurrencyRepository is an interface for currency repository layer
type CurrencyRepository interface {
	Fetch() ([]*Currency, error)
	FetchOne(from, to string, lastNRates int) (*Currency, error)
	Update(uint, *Currency) (*Currency, error)
	Store(*Currency) error
}
