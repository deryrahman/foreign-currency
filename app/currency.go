package app

// Currency is a struct for currency
// It consist of attributes ID, From, To and Rate
type Currency struct {
	ID    uint   `json:"id"`
	From  string `json:"from"`
	To    string `json:"to"`
	Rates []Rate `json:"rates"`
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
	Store(*Currency) error
}
