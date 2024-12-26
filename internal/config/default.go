package config

import "github.com/spf13/viper"

type (
	Config struct {
		Port              string `mapstructure:"PORT"`
		OpenWeatherAPIKey string `mapstructure:"OPENWEATHER_API_KEY"`
		WeatherAPIKey     string `mapstructure:"WEATHER_API_KEY"`
		RedisURL          string `mapstructure:"REDIS_URL"`
		CacheDuration     int    `mapstructure:"CACHE_DURATION"`
		RateLimit         int    `mapstructure:"RATE_LIMIT"`
	}
)

func Load() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	config := &Config{}
	if err := viper.Unmarshal(config); err != nil {
		return nil, err
	}

	return config, nil
}
