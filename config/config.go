package config

import (
	"os"
	"strings"
	"time"
)

// Config holds application configuration values.
type Config struct {
	// Port the server listens on (default "8080").
	Port string

	// HTTP server timeouts
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
	// AllowedOrigins is a list of allowed CORS origins (comma-separated in ALLOWED_ORIGINS).
	AllowedOrigins []string

	// AllowCredentials controls whether Access-Control-Allow-Credentials is set.
	AllowCredentials bool
}

// Load reads configuration from environment variables and returns a Config with sensible defaults.
func Load() Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// parse allowed origins from env (comma-separated)
	var allowed []string
	if v := os.Getenv("ALLOWED_ORIGINS"); v != "" {
		for _, s := range strings.Split(v, ",") {
			s = strings.TrimSpace(s)
			if s != "" {
				allowed = append(allowed, s)
			}
		}
	} else {
		// sensible defaults for local dev + known dev route
		allowed = []string{
			"http://localhost:8080",
			"https://efv-api-17db4f-dev.apps.silver.devops.gov.bc.ca",
		}
	}

	allowCred := false
	if os.Getenv("ALLOW_CREDENTIALS") == "true" {
		allowCred = true
	}

	// Keep defaults consistent with previous main.go values
	return Config{
		Port:             port,
		ReadTimeout:      5 * time.Second,
		WriteTimeout:     10 * time.Second,
		IdleTimeout:      120 * time.Second,
		AllowedOrigins:   allowed,
		AllowCredentials: allowCred,
	}
}

// MustLoad is a convenience that returns a pointer to loaded Config.
func MustLoad() *Config {
	cfg := Load()
	return &cfg
}
