package currency

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestCurrencyconverter(t *testing.T) {
	req, err := http.NewRequest("GET", "/currencyconverter/:USD/:EUR/:100", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	// Create a dummy router
	router := mux.NewRouter()
	router.HandleFunc("/currencyconverter/:{base}/:{target}/:{amount}", Currencyconverter).Methods("GET")
	router.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"base_code": "USD","target_code": "EUR","conversion_result": 92.66,"conversion_rate": 0.9266,"time_last_update_utc": "Mon, 01 Apr 2024 00:00:01 +0000","ExpirationTime": "Mon, 01 Apr 2024 18:24:28 +0530"}`
    if rr.Body.String() == expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}

func TestFetchCurrencyRates(t *testing.T) {
	req, err := http.NewRequest("GET", "/fetchCurrentRate/:USD/:EUR/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	// Create a dummy router
	router := mux.NewRouter()
	router.HandleFunc("/fetchCurrentRate/:{base}/:{target}/", Currencyconverter).Methods("GET")
	router.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"base_code": "USD","target_code": "EUR","conversion_rate": 0.9266,"time_last_update_utc": "Mon, 01 Apr 2024 00:00:01 +0000","ExpirationTime": "Mon, 01 Apr 2024 18:24:28 +0530"}`
    if rr.Body.String() == expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}

