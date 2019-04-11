package model

// Currency is a struct for currency
// It consist of attributes ID, From, To and Rate
type Currency struct {
	ID    uint   `json:"id"`
	From  string `json:"from"`
	To    string `json:"to"`
	Rates []Rate `json:"rates,omitempty"`
}
