package currency

import (
	"CurrencyConverter/config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func Currencyconverter(w http.ResponseWriter, r *http.Request) {
	// Specify the base currency, target currency, and amount for conversion
	vars := mux.Vars(r)
	baseCurrency := vars["base"]
	targetCurrency := vars["target"]
	amount, _ := strconv.ParseFloat(vars["amount"], 64)

	// Construct the API endpoint URL
	endpoint := fmt.Sprintf("https://v6.exchangerate-api.com/v6/%s/pair/%s/%s/%f", config.APIKey, baseCurrency, targetCurrency, amount)

	// Make GET request to the API
	response, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer response.Body.Close()

	// Read response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// Unmarshal response JSON
	var conversionResp config.ConversionResponse
	if err := json.Unmarshal(body, &conversionResp); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	expireTime := time.Now().Add(3 * time.Hour)
	conversionResp.ExpirationTime = expireTime.Format("Mon, 02 Jan 2006 15:04:05 -0700")

	// Print conversion result
	fmt.Printf("%f %s equals %f %s (Last updated: %s)\n", amount, conversionResp.BaseCode, conversionResp.ConversionResult, conversionResp.Target, conversionResp.LastUpdate)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(conversionResp)
}

func FetchCurrencyRates(w http.ResponseWriter, r *http.Request) {
	// Fetch the base currency, target currency
	vars := mux.Vars(r)
	baseCurrency := vars["base"]
	targetCurrency := vars["target"]
	cache := config.Memory.Data
	flag := true
	if len(cache) > 0 {
		flag = false
		config.Memory.Mutex.Lock()
		defer config.Memory.Mutex.Unlock()

		value, ok := cache[baseCurrency]
		if !ok {
			flag = true
		}
		layout := "Mon, 02 Jan 2006 15:04:05 -0700"
		// Check if the cached value has expired
		expirationTime, err := time.Parse(layout, value.ExpirationTime)
		if err != nil || time.Now().After(expirationTime) {
			delete(cache, baseCurrency)
			flag = true
		} else {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(value)
		}
	}
	if flag {
		// Construct the API endpoint URL
		endpoint := fmt.Sprintf("https://v6.exchangerate-api.com/v6/%s/pair/%s/%s/", config.APIKey, baseCurrency, targetCurrency)

		// Make GET request to the API
		response, err := http.Get(endpoint)
		if err != nil {
			fmt.Println("Error making request:", err)
			return
		}
		defer response.Body.Close()

		// Read response body
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Error reading response:", err)
			return
		}

		// Unmarshal response JSON
		var conversionResp config.ConversionResponse
		if err := json.Unmarshal(body, &conversionResp); err != nil {
			fmt.Println("Error decoding JSON:", err)
			return
		}
		expireTime := time.Now().Add(3 * time.Hour)
		conversionResp.ExpirationTime = expireTime.Format("Mon, 02 Jan 2006 15:04:05 -0700")

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(conversionResp)
		config.Memory.Data[baseCurrency] = conversionResp
	}

}
