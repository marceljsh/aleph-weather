package api

import (
	"net/http"
	"time"

	"github.com/marceljsh/aleph-weather/internal/metrics"
	"github.com/marceljsh/aleph-weather/internal/service"
	"github.com/marceljsh/aleph-weather/pkg/logger"
	"github.com/marceljsh/aleph-weather/pkg/respond"
	"go.uber.org/zap"
)

type (
	Handler struct {
		weatherService service.WeatherService
	}
)

func NewHandler(weatherService service.WeatherService) *Handler {
	return &Handler{weatherService}
}

func (h *Handler) GetWeather(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	city := r.URL.Query().Get("city")
	if city == "" {
		metrics.RequestsTotal.WithLabelValues("error").Inc()
		respond.BadRequest(w, "city parameter is required")
		return
	}

	weather, err := h.weatherService.GetWeather(r.Context(), city)
	if err != nil {
		metrics.RequestsTotal.WithLabelValues("error").Inc()
		logger.Error("Failed to get weather",
			zap.String("city", city),
			zap.Error(err))
		respond.InternalErr(w, err.Error())
		return
	}

	metrics.RequestsTotal.WithLabelValues("success").Inc()
	metrics.APILatency.WithLabelValues(weather.Source).
		Observe(time.Since(start).Seconds())

	respond.Ok(w, weather)
}

func (h *Handler) GetStats(w http.ResponseWriter, r *http.Request) {
	stats, err := h.weatherService.GetStats(r.Context())
	if err != nil {
		respond.InternalErr(w, err.Error())
		return
	}

	respond.Ok(w, stats)
}

func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	status := map[string]string{
		"status":    "ok",
		"timestamp": time.Now().UTC().Format(time.RFC3339),
	}
	respond.Ok(w, status)
}
