package currencies

import "errors"

/*
  CurrenciesRecordSet holds item for currencies records
  Code   : Currency Code (ISO 4217)
  Name   : Currency name (for example Canadian Dollar)
  TypeID : Currency type ID. Possible values from types ID table are digital, crypto, physical
*/

type CurrenciesRecordSet struct {
	ID     int64  `json:"id"`
	Code   string `json:"code"`
	Name   string `json:"name"`
	TypeID int64  `json:"type_id"`
}

// errors
var (
	// ErrRecordInvalid invalid record found
	ErrRecordInvalid = errors.New("Record is invalid")
)
