package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewRouter(handler *Handler) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)

	r.Get("/weather", handler.GetWeather)
	r.Get("/stats", handler.GetStats)
	r.Get("/health", handler.HealthCheck)
	r.Handle("/metrics", promhttp.Handler())

	return r
}
