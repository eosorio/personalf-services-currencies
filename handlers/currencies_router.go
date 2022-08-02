package handlers

import (
	"net/http"
	"strconv"
)

// CurrenciesRouter handles the currencies route
func CurrenciesRouter(w http.ResponseWriter, r *http.Request) {

	var (
		sCurrencyID   []string
		sCurrencyType []string

		currencyID   int64
		currencyType int64
		err          error
	)

	// Get parameters passed in the URL
	query := r.URL.Query()
	for key, value := range query {
		switch key {
		case "currencyId":
			sCurrencyID = value
		case "currencyType":
			sCurrencyType = value
		}
	}

	// Only one currency ID is accepted. Anything different than 1 -> 0. Returns all values
	if len(sCurrencyID) == 1 {
		if currencyID, err = strconv.ParseInt(sCurrencyID[0], 10, 64); err != nil {
			postError(w, http.StatusMethodNotAllowed)
			return
		}
	}

	// Only one currency type ID is accepted. Anything different than 1 -> 0. Returns all values
	if len(sCurrencyType) == 1 {
		if currencyType, err = strconv.ParseInt(sCurrencyType[0], 10, 64); err != nil {
			postError(w, http.StatusMethodNotAllowed)
			return
		}
	}

	switch r.Method {
	case http.MethodGet:
		currenciesGetInfo(w, r, currencyID, currencyType)
		return
	case http.MethodPost:
		return
	default:
		postError(w, http.StatusMethodNotAllowed)
		return
	}

	switch r.Method {
	case http.MethodPut:
	case http.MethodPatch:
	case http.MethodDelete:
		postError(w, http.StatusMethodNotAllowed)
		return
	}
}
