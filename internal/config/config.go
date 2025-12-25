package config

import "os"

type Config struct {
	Port          string
	AllowedOrigin string
	DBPath        string
	AdminToken    string
}

func Load() Config {
	return Config{
		Port:          getEnv("PORT", "8080"),
		AllowedOrigin: getEnv("ALLOWED_ORIGIN", "*"),
		DBPath:        getEnv("DB_PATH", "blog.db"),
		AdminToken:    getEnv("ADMIN_TOKEN", ""),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
