package cache

import (
	"context"
	"time"

	"github.com/marceljsh/aleph-weather/internal/models"
	"github.com/redis/go-redis/v9"
)

type (
	Cache interface {
		Get(ctx context.Context, key string) (*models.WeatherData, error)
		Set(ctx context.Context, key string, value *models.WeatherData, expiration time.Duration) error
		IncrementStats(ctx context.Context, key string) error
		GetStats(ctx context.Context) (*models.Stats, error)
	}

	redisCache struct {
		client *redis.Client
	}
)

func NewRedis(url string) (Cache, error) {
	opts, err := redis.ParseURL(url)
	if err != nil {
		return nil, err
	}

	client := redis.NewClient(opts)
	return &redisCache{client: client}, nil
}
