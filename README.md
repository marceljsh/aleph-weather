# Real-Time Weather Aggregator

A Go-based backend service that aggregates weather data from multiple sources, implements caching, and provides real-time weather information.

## Features

- Concurrent weather data fetching from multiple providers
- Redis caching with configurable duration
- Rate limiting for API requests
- Prometheus metrics
- Structured logging
- Docker support

## Requirements

- Go 1.21+
- Redis
- Docker (optional)
- API keys for:
  - OpenWeatherMap
  - WeatherAPI

## Installation

### Local Setup
1. Clone the repository:
```bash
git clone https://github.com/marceljsh/aleph-weather.git
cd weather-aggregator
```

2. Install dependencies:
```bash
go mod download
```

3. Run the application:
```bash
make run
```

### Docker Setup

1. Build and run using Docker Compose:
```bash
make docker-build
```
