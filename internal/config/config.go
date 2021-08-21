// Package config contains configuration structs
package config

import (
	"fmt"
	"os"
	"path"

	yaml "gopkg.in/yaml.v2"
)

// Config is the main configuration struct for the application
type Config struct {
	Authenticator Authenticator `yaml:"authenticator"`
}

// ReadConfigFromFile reads the config from a file.
func ReadConfigFromFile(cfgFile string) (*Config, error) {
	f, err := os.Open(path.Clean(cfgFile))
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %w", err)
	}
	var cfg Config
	err = yaml.NewDecoder(f).Decode(&cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}
	return &cfg, nil
}
