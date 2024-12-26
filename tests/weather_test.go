package tests

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/marceljsh/aleph-weather/internal/models"
	"github.com/marceljsh/aleph-weather/internal/providers"
	"github.com/marceljsh/aleph-weather/internal/service"
)

type (
	MockWeatherProvider struct {
		mock.Mock
	}

	MockCache struct {
		mock.Mock
	}
)

func (m *MockWeatherProvider) GetWeather(ctx context.Context, city string) (*models.WeatherData, error) {
	args := m.Called(ctx, city)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.WeatherData), args.Error(1)
}

func (m *MockWeatherProvider) Name() string {
	return "MockProvider"
}

func (m *MockCache) Get(ctx context.Context, key string) (*models.WeatherData, error) {
	args := m.Called(ctx, key)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.WeatherData), args.Error(1)
}

func (m *MockCache) Set(ctx context.Context, key string, value *models.WeatherData, expiration time.Duration) error {
	args := m.Called(ctx, key, value, expiration)
	return args.Error(0)
}

func (m *MockCache) IncrementStats(ctx context.Context, key string) error {
	args := m.Called(ctx, key)
	return args.Error(0)
}

func (m *MockCache) GetStats(ctx context.Context) (*models.Stats, error) {
	args := m.Called(ctx)
	return args.Get(0).(*models.Stats), args.Error(1)
}

func TestWeatherService_GetWeather(t *testing.T) {
	mockProvider := new(MockWeatherProvider)
	mockCache := new(MockCache)

	weatherData := &models.WeatherData{
		City:        "London",
		Temperature: 20.5,
		Humidity:    65,
		Condition:   "Cloudy",
		Source:      "MockProvider",
		Timestamp:   time.Now(),
	}

	tests := []struct {
		name          string
		city          string
		setupMocks    func()
		expectedData  *models.WeatherData
		expectedError bool
	}{
		{
			name: "Cache Hit",
			city: "London",
			setupMocks: func() {
				mockCache.On("Get", mock.Anything, "weather:London").
					Return(weatherData, nil)
				mockCache.On("IncrementStats", mock.Anything, "stats:cache_hits").
					Return(nil)
			},
			expectedData:  weatherData,
			expectedError: false,
		},
		{
			name: "Cache Miss - Provider Success",
			city: "London",
			setupMocks: func() {
				mockCache.On("Get", mock.Anything, "weather:London").
					Return(nil, errors.New("not found"))
				mockCache.On("IncrementStats", mock.Anything, "stats:cache_misses").
					Return(nil)
				mockCache.On("IncrementStats", mock.Anything, "stats:api_calls").
					Return(nil)
				mockProvider.On("GetWeather", mock.Anything, "London").
					Return(weatherData, nil)
				mockCache.On("Set", mock.Anything, "weather:London", weatherData, mock.Anything).
					Return(nil)
			},
			expectedData:  weatherData,
			expectedError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMocks()

			service := service.NewWeatherService(
				[]providers.WeatherProvider{mockProvider},
				mockCache,
				nil,
				30*time.Minute,
			)

			result, err := service.GetWeather(context.Background(), tt.city)

			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedData, result)
			}

			mockProvider.AssertExpectations(t)
			mockCache.AssertExpectations(t)
		})
	}
}
