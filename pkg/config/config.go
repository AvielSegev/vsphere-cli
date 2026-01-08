package config

import (
	"fmt"
	"os"
	"strconv"
)

// Config holds the vSphere connection configuration
type Config struct {
	Host     string
	Username string
	Password string
	Insecure bool
}

// Source tracks where a config value came from
type Source string

const (
	SourceEnv     Source = "env"
	SourceFlag    Source = "flag"
	SourceDefault Source = "default"
)

// ConfigWithSource holds config values with their sources
type ConfigWithSource struct {
	Host     string
	HostSrc  Source
	Username string
	UserSrc  Source
	Password string
	PassSrc  Source
	Insecure bool
	InsecSrc Source
}

// LoadFromEnv loads configuration from environment variables
func LoadFromEnv() (*Config, error) {
	cfg := &Config{
		Host:     os.Getenv("VCLI_HOST"),
		Username: os.Getenv("VCLI_USERNAME"),
		Password: os.Getenv("VCLI_PASSWORD"),
		Insecure: false,
	}

	if insecure := os.Getenv("VCLI_INSECURE"); insecure != "" {
		val, err := strconv.ParseBool(insecure)
		if err != nil {
			return nil, fmt.Errorf("invalid VCLI_INSECURE value: %w", err)
		}
		cfg.Insecure = val
	}

	return cfg, nil
}

// Validate checks if required configuration is present
func (c *Config) Validate() error {
	if c.Host == "" {
		return fmt.Errorf("VCLI_HOST is required")
	}
	if c.Username == "" {
		return fmt.Errorf("VCLI_USERNAME is required")
	}
	if c.Password == "" {
		return fmt.Errorf("VCLI_PASSWORD is required")
	}
	return nil
}

// MaskPassword returns password with only first 2 and last 2 chars visible
func MaskPassword(password string) string {
	if len(password) <= 4 {
		return "••••"
	}
	return password[:2] + "••••••" + password[len(password)-2:]
}
