package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	RequestsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "weather_requests_total",
			Help: "Total number of weather requests",
		},
		[]string{"status"},
	)

	CacheHits = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "weather_cache_hits_total",
			Help: "Total number of cache hits",
		},
	)

	CacheMisses = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "weather_cache_misses_total",
			Help: "Total number of cache misses",
		},
	)

	APILatency = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "weather_api_latency_seconds",
			Help:    "API request latency in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"provider"},
	)
)
