// Package config provides application configuration.
package config

import (
	"errors"
	"fmt"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

// Config holds the application configuration.
type Config struct {
	Env               string        `env:"ENV,notEmpty,required"`
	PORT              int           `env:"PORT,notEmpty,required"`
	Domain            string        `env:"DOMAIN,notEmpty,required"`
	ReadHeaderTimeout time.Duration `env:"READ_HEADER_TIMEOUT" envDefault:"5s"`
	ReadTimeout       time.Duration `env:"READ_TIMEOUT" envDefault:"10s"`
	WriteTimeout      time.Duration `env:"WRITE_TIMEOUT" envDefault:"30s"`
	IdleTimeout       time.Duration `env:"IDLE_TIMEOUT" envDefault:"60s"`
	ShutdownTimeout   time.Duration `env:"SHUTDOWN_TIMEOUT" envDefault:"10s"`
}

// Load loads the configuration from environment variables.
func Load() (*Config, []string, error) {
	var notices []string
	err := godotenv.Load()
	if err != nil {
		notices = append(notices, "No .env file found, rely solely on system environement variables")
	}

	var cfg Config

	if err := env.ParseWithOptions(&cfg, env.Options{Prefix: "APP_"}); err != nil {
		return &Config{}, notices, fmt.Errorf("failed to parse environment variables: %w", err)
	}

	if err := cfg.validate(); err != nil {
		return &Config{}, notices, fmt.Errorf("invalid configuration: %w", err)
	}

	return &cfg, notices, nil
}

func (c *Config) validate() error {
	switch c.Env {
	case "dev", "test", "prod":
	default:
		return errors.New("APP_ENV must be one of: dev, test, prod")
	}

	if c.PORT < 1024 || c.PORT > 65535 {
		return fmt.Errorf("APP_PORT %d is invalid; must be between 1024 and 65535", c.PORT)
	}

	return nil
}

// IsProduction returns true if the environment is production.
func (c *Config) IsProduction() bool {
	return c.Env == "prod"
}
