package handlers

import (
	"net/http"

	"github.com/eosorio/personalf-services-currencies/currencies"
)

func currenciesGetInfo(w http.ResponseWriter, _ *http.Request, currencyID int64, currencyType int64) {
	currenciesInfo, err := currencies.GetCurrencies(currencyID, currencyType)
	if err != nil {
		postError(w, http.StatusInternalServerError)
		return
	}
	if len(currenciesInfo) == 0 {
		postError(w, http.StatusNotFound)
	} else {
		postBodyResponse(w, http.StatusOK, jsonResponse{"currencies": currenciesInfo})
	}
}
