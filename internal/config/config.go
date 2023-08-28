// Package config provides functions for parsing configuration files for a Go web service.
package config

import (
	"github.com/caarlos0/env/v9"
)

// Config struct
type Config struct {
	PgxDBAddr string `env:"PGXCONN" envDefault:"postgres://eugen:ur2qly1ini@localhost:5432/eugene"`
}

// NewConfig creates a new Config instance
func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
