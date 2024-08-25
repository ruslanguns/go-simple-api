package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Port int
	// Add other configuration fields as needed
}

func Load() (*Config, error) {
	port, err := strconv.Atoi(getEnv("PORT", "8080"))
	if err != nil {
		return nil, fmt.Errorf("invalid port: %w", err)
	}

	return &Config{
		Port: port,
	}, nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
