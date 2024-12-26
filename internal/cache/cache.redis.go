package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/marceljsh/aleph-weather/internal/models"
	"github.com/marceljsh/aleph-weather/pkg/parse"
	"github.com/redis/go-redis/v9"
)

func (c *redisCache) Get(ctx context.Context, key string) (*models.WeatherData, error) {
	val, err := c.client.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var weather models.WeatherData
	if err := json.Unmarshal([]byte(val), &weather); err != nil {
		return nil, err
	}

	return &weather, nil
}

func (c *redisCache) Set(ctx context.Context, key string, value *models.WeatherData, expiration time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return c.client.Set(ctx, key, data, expiration).Err()
}

func (c *redisCache) IncrementStats(ctx context.Context, key string) error {
	return c.client.Incr(ctx, key).Err()
}

func (c *redisCache) GetStats(ctx context.Context) (*models.Stats, error) {
	pipe := c.client.Pipeline()

	apiCalls := pipe.Get(ctx, "stats:api_calls")
	cacheHits := pipe.Get(ctx, "stats:cache_hits")
	cacheMisses := pipe.Get(ctx, "stats:cache_misses")

	_, err := pipe.Exec(ctx)
	if err != nil && err != redis.Nil {
		return nil, err
	}

	stats := &models.Stats{
		APICalls:    parse.IntOrZero(apiCalls.Val()),
		CacheHits:   parse.IntOrZero(cacheHits.Val()),
		CacheMisses: parse.IntOrZero(cacheMisses.Val()),
	}

	return stats, nil
}
