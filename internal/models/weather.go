package models

import "time"

type (
	WeatherData struct {
		City        string    `json:"city"`
		Temperature float64   `json:"temperature"`
		Humidity    int       `json:"humidity"`
		Condition   string    `json:"condition"`
		Source      string    `json:"source"`
		Cached      bool      `json:"cached"`
		Timestamp   time.Time `json:"timestamp"`
	}

	Stats struct {
		APICalls    int64 `json:"api_calls"`
		CacheHits   int64 `json:"cache_hits"`
		CacheMisses int64 `json:"cache_misses"`
	}
)
