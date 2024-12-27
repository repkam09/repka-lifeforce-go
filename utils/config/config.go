package config

import (
	"os"
	"sync"
)

type Config struct {
	WEATHER_API_KEY string
	LAST_FM_API_KEY string
	DEBUG_MODE      string
}

var (
	configInstance *Config
	once           sync.Once
)

func GetConfig() *Config {
	once.Do(func() {
		configInstance = &Config{
			WEATHER_API_KEY: getRequiredEnv("WEATHER_API_KEY"),
			LAST_FM_API_KEY: getRequiredEnv("LAST_FM_API_KEY"),
			DEBUG_MODE:      getEnv("DEBUG_MODE", "false"),
		}
	})
	return configInstance
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getRequiredEnv(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	panic("Required environment variable not set: " + key)
}
