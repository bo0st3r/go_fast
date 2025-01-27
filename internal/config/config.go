package config

import (
	"fmt"
	"os"
)

type Config struct {
	DatabaseDSN string
	Port        string
}

func Load() (Config, error) {
	requiredEnvVars := map[string]string{
		"POSTGRES_USER":     os.Getenv("POSTGRES_USER"),
		"POSTGRES_PASSWORD": os.Getenv("POSTGRES_PASSWORD"),
		"POSTGRES_DB":       os.Getenv("POSTGRES_DB"),
		"APP_PORT":          os.Getenv("APP_PORT"),
	}

	// Validate all required environment variables are set
	for name, value := range requiredEnvVars {
		if value == "" {
			return Config{}, fmt.Errorf("required environment variable %s is not set", name)
		}
	}

	// Construct database DSN
	databaseDSN := fmt.Sprintf("postgres://%s:%s@localhost:5432/%s?sslmode=disable",
		requiredEnvVars["POSTGRES_USER"],
		requiredEnvVars["POSTGRES_PASSWORD"],
		requiredEnvVars["POSTGRES_DB"])

	return Config{
		DatabaseDSN: databaseDSN,
		Port:        requiredEnvVars["APP_PORT"],
	}, nil
}
