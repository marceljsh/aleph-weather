package service

import (
	"context"
	"fmt"
	"sync"

	"github.com/marceljsh/aleph-weather/internal/models"
	"github.com/marceljsh/aleph-weather/internal/providers"
)

func (s *weatherService) GetStats(ctx context.Context) (*models.Stats, error) {
	return s.cache.GetStats(ctx)
}

func (s *weatherService) GetWeather(ctx context.Context, city string) (*models.WeatherData, error) {
	// coba cache dulu
	cacheKey := fmt.Sprintf("weather:%s", city)
	if cached, err := s.cache.Get(ctx, cacheKey); err == nil {
		cached.Cached = true
		s.cache.IncrementStats(ctx, "stats:cache_hits")
		return cached, nil
	}
	s.cache.IncrementStats(ctx, "stats:cache_misses")

	// cek rate limit
	if err := s.rateLimiter.Wait(ctx); err != nil {
		return nil, fmt.Errorf("rate limit exceeded: %w", err)
	}

	results := make(chan *models.WeatherData, len(s.providers))
	errors := make(chan error, len(s.providers))
	var wg sync.WaitGroup

	for _, provider := range s.providers {
		wg.Add(1)
		go func(p providers.WeatherProvider) {
			defer wg.Done()

			weather, err := p.GetWeather(ctx, city)
			if err != nil {
				errors <- fmt.Errorf("%s: %w", p.Name(), err)
				return
			}
			results <- weather
		}(provider)
	}

	// tunggu semua
	go func() {
		wg.Wait()
		close(results)
		close(errors)
	}()

	var weather *models.WeatherData
	select {
	case weather = <-results:
		s.cache.IncrementStats(ctx, "stats:api_calls")
		if err := s.cache.Set(ctx, cacheKey, weather, s.cacheDuration); err != nil {
			return nil, fmt.Errorf("caching result: %w", err)
		}
		return weather, nil

	case err := <-errors:
		return nil, err

	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
