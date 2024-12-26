package providers

import (
	"context"
	"net/http"
	"time"

	"github.com/marceljsh/aleph-weather/internal/models"
)

type (
	WeatherProvider interface {
		Name() string
		GetWeather(ctx context.Context, city string) (*models.WeatherData, error)
	}

	openWeatherMap struct {
		apiKey     string
		baseUrl    string
		httpClient *http.Client
	}

	openWeatherResponse struct {
		Main struct {
			Temperture float64 `json:"temp"`
			Humidity   int     `json:"humidity"`
		} `json:"main"`
		Weather []struct {
			Description string `json:"description"`
		} `json:"weather"`
		Name string `json:"name"`
	}

	weatherAPI struct {
		apiKey     string
		baseURL    string
		httpClient *http.Client
	}

	weatherAPIResponse struct {
		Current struct {
			TempC     float64 `json:"temp_c"`
			Humidity  int     `json:"humidity"`
			Condition struct {
				Text string `json:"text"`
			} `json:"condition"`
		} `json:"current"`
		Location struct {
			Name string `json:"name"`
		} `json:"location"`
	}
)

func NewOWM(apiKey string) WeatherProvider {
	return &openWeatherMap{
		apiKey:  apiKey,
		baseUrl: "https://api.openweathermap.org/data/2.5/weather",
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func NewWAPI(apiKey string) WeatherProvider {
	return &weatherAPI{
		apiKey:  apiKey,
		baseURL: "http://api.weatherapi.com/v1",
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}
