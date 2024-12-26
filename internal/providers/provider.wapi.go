package providers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/marceljsh/aleph-weather/internal/models"
)

func (p *weatherAPI) Name() string {
	return "WeatherAPI"
}

func (p *weatherAPI) GetWeather(ctx context.Context, city string) (*models.WeatherData, error) {
	url := fmt.Sprintf("%s/current.json?key=%s&q=%s", p.baseURL, p.apiKey, city)

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

	var result weatherAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("decoding response: %w", err)
	}

	return &models.WeatherData{
		City:        result.Location.Name,
		Temperature: result.Current.TempC,
		Humidity:    result.Current.Humidity,
		Condition:   result.Current.Condition.Text,
		Source:      p.Name(),
		Timestamp:   time.Now(),
	}, nil
}
