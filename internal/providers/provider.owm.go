package providers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/marceljsh/aleph-weather/internal/models"
)

func (p *openWeatherMap) Name() string {
	return "OpenWeatherMap"
}

func (p *openWeatherMap) GetWeather(ctx context.Context, city string) (*models.WeatherData, error) {
	url := fmt.Sprintf("%s/weather?q=%s&appid=%s&units=metric", p.baseUrl, city, p.apiKey)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	resp, err := p.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var result openWeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("decoding response: %w", err)
	}

	return &models.WeatherData{
		City:        result.Name,
		Temperature: result.Main.Temperture,
		Humidity:    result.Main.Humidity,
		Condition:   result.Weather[0].Description,
		Source:      p.Name(),
		Timestamp:   time.Now(),
	}, nil
}
