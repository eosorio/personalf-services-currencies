package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/eosorio/personalf-services-currencies/handlers"
)

func main() {
	http.HandleFunc("/currencies", handlers.CurrenciesRouter)
	err := http.ListenAndServe("0.0.0.0:11114", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
