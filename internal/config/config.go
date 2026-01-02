package config

import "os"

type Config struct {
	AppName string
	Port    string
}

func Load() *Config {
	return &Config{
		AppName: getEnv("APP_NAME", "gocore"),
		Port:    getEnv("APP_PORT", "8080"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
