package config

import (
	"sync"
)

const APIKey = "fb38bdca9f60ed39e028849b" // Replace with your actual API key

type ConversionResponse struct {
	BaseCode         string  `json:"base_code"`
	Target           string  `json:"target_code"`
	ConversionResult float64 `json:"conversion_result"`
	ConversionRate   float64 `json:"conversion_rate"`
	LastUpdate       string  `json:"time_last_update_utc"`
	ExpirationTime   string
}
type Cache struct {
	Data  map[string]ConversionResponse
	Mutex *sync.RWMutex
}

var Memory *Cache

func NewCache() {
	Memory = &Cache{
		Data: make(map[string]ConversionResponse),
		Mutex: new(sync.RWMutex),
	}
}
