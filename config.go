package config

import (
	"fmt"
	"os"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type Config struct {
	envPath string
	prefix  string
}

func New(opts ...OptFunc) *Config {
	c := &Config{
		envPath: ".env",
		prefix:  "",
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func (c *Config) Load(dest interface{}) (err error) {
	if _, err := os.Stat(c.envPath); err == nil {
		if err := godotenv.Load(c.envPath); err != nil {
			return fmt.Errorf("failed to read from .env file: %w", err)
		}
	}

	envopt := env.Options{Prefix: c.prefix}
	err = env.ParseWithOptions(dest, envopt)
	if err != nil {
		return fmt.Errorf("failed to load env: %w", err)
	}

	return nil
}
