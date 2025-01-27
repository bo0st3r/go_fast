package config

import (
	"fmt"
	"os"
)

type Config struct {
	DatabaseDSN string
	Port        string
}

func Load() Config {
	username := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	database := os.Getenv("POSTGRES_DB")
	databaseDSN := fmt.Sprintf("postgres://%s:%s@localhost:5432/%s?sslmode=disable", username, password, database)
	port := os.Getenv("APP_PORT")
	return Config{DatabaseDSN: databaseDSN, Port: port}
}
