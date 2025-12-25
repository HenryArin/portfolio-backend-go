package config

import "os"

type Config struct {
	Port          string
	AllowedOrigin string
}

func Load() Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	origin := os.Getenv("ALLOWED_ORIGIN")
	if origin == "" {
		origin = "*"
	}

	return Config{
		Port:          port,
		AllowedOrigin: origin,
	}
}
