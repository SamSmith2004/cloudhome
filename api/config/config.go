package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Server struct {
		Port int
	}
}

func LoadConfig() (*Config, error) {
	// err := godotenv.Load()
	// if err != nil {
	// 	return nil, fmt.Errorf("Failed to load env: %w", err)
	// }

	config := &Config{}

	// Server configuration
	portStr := getEnv("PORT", "8080")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, fmt.Errorf("invalid PORT: %w", err)
	}
	config.Server.Port = port

	return config, nil
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
