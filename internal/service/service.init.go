package service

import (
	"context"
	"time"

	"github.com/marceljsh/aleph-weather/internal/cache"
	"github.com/marceljsh/aleph-weather/internal/models"
	"github.com/marceljsh/aleph-weather/internal/providers"
	"github.com/marceljsh/aleph-weather/internal/ratelimit"
)

type (
	WeatherService interface {
		GetStats(ctx context.Context) (*models.Stats, error)
		GetWeather(ctx context.Context, city string) (*models.WeatherData, error)
	}

	weatherService struct {
		providers     []providers.WeatherProvider
		cache         cache.Cache
		rateLimiter   ratelimit.RateLimiter
		cacheDuration time.Duration
	}
)

func NewWeatherService(
	providers []providers.WeatherProvider,
	cache cache.Cache,
	rateLimiter ratelimit.RateLimiter,
	cacheDuration time.Duration,
) WeatherService {
	return &weatherService{
		providers:     providers,
		cache:         cache,
		rateLimiter:   rateLimiter,
		cacheDuration: cacheDuration,
	}
}
