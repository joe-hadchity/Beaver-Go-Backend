package config

import (
	"fmt"
	"os"
	"strings"

)

type Config struct {
        DatabaseURL string
        Port        string
}

func LoadConfig() (*Config, error) { // Return *Config and error

	dbURL := os.Getenv("DATABASE_URL")
	// Force SSL mode if not present
	if !strings.Contains(dbURL, "sslmode=") {
		dbURL += "?sslmode=require"
	}
	port := os.Getenv("PORT")

	if dbURL == "" {
			return nil, fmt.Errorf("DATABASE_URL environment variable is not set") // Return error
	}
	if port == "" {
			port = "8080" // Default if PORT is not set
	}

	return &Config{
			DatabaseURL: dbURL,
			Port:        port,
	}, nil // Return nil error if all is well
}