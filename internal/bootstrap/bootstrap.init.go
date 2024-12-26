package bootstrap

import (
	"log"
	"net/http"
	"time"

	"github.com/marceljsh/aleph-weather/internal/api"
	"github.com/marceljsh/aleph-weather/internal/cache"
	"github.com/marceljsh/aleph-weather/internal/config"
	"github.com/marceljsh/aleph-weather/internal/providers"
	"github.com/marceljsh/aleph-weather/internal/ratelimit"
	"github.com/marceljsh/aleph-weather/internal/service"
)

type (
	app struct {
		server *http.Server
	}
)

func App() *app {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	redisCache, err := cache.NewRedis(cfg.RedisURL)
	if err != nil {
		log.Fatalf("Failed to initialize cache: %v", err)
	}

	weatherProviders := []providers.WeatherProvider{
		providers.NewOWM(cfg.OpenWeatherAPIKey),
		providers.NewWAPI(cfg.WeatherAPIKey),
	}

	rateLimiter := ratelimit.NewTokenBucket(cfg.RateLimit)

	weatherService := service.NewWeatherService(
		weatherProviders,
		redisCache,
		rateLimiter,
		time.Duration(cfg.CacheDuration)*time.Minute,
	)

	handler := api.NewHandler(weatherService)
	router := api.NewRouter(handler)

	server := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: router,
	}

	return &app{
		server: server,
	}
}
