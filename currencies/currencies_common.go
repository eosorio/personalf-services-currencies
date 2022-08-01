package currencies

import "errors"

/*
  CurrenciesRecordSet holds item for currencies records
  Code : Currency Code (ISO 4217)
  Name : Currency name (for example Canadian Dollar)
  Type : Currency type. Possible values are digital, crypto, physical
*/
type CurrenciesRecordSet struct {
	ID   int64  `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
	Type string `json:"type"`
}

// errors
var (
	// ErrRecordInvalid invalid record found
	ErrRecordInvalid = errors.New("Record is invalid")
)
