//package currencyconverter
package main
import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

const APIKey = "fb38bdca9f60ed39e028849b" // Replace with your actual API key

type ConversionResponse struct {
    BaseCode  string             `json:"base_code"`
    Target    string             `json:"target_code"`
    Value     float64            `json:"conversion_result"`
    LastUpdate string             `json:"time_last_update_utc"`
}

func Currencyconverter(w http.ResponseWriter, r *http.Request) {
    // Specify the base currency, target currency, and amount for conversion
    baseCurrency := "USD"
    targetCurrency := "EUR"
    amount := 100.0

    // Construct the API endpoint URL
	endpoint := fmt.Sprintf("https://v6.exchangerate-api.com/v6/%s/pair/%s/%s/%f",APIKey,baseCurrency,targetCurrency,amount)

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
    var conversionResp ConversionResponse
    if err := json.Unmarshal(body, &conversionResp); err != nil {
        fmt.Println("Error decoding JSON:", err)
        return
    }

    // Print conversion result
    fmt.Printf("%f %s equals %f %s (Last updated: %s)\n", amount, conversionResp.BaseCode, conversionResp.Value, conversionResp.Target, conversionResp.LastUpdate)

    w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode()
}

