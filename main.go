package main

import (
	"log"
	"net/http"

	"CurrencyConverter/config"
	"CurrencyConverter/currency"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	config.NewCache()
	// Define routes
	router.HandleFunc("/currencyconverter/base/{base}/target/{target}/amount/{amount}", currency.Currencyconverter).Methods("GET")
	router.HandleFunc("/fetchCurrentRate/base/{base}/target/{target}", currency.FetchCurrencyRates).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))

}
