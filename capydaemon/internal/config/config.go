package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Config struct {
	TLSCertFile string `json:"tls_cert_file"`
	TLSKeyFile  string `json:"tls_key_file"`
	ServerAddr  string `json:"server_addr"`
	Port        int    `json:"port"`
}

func LoadConfig(path string) (*Config, error) {
	if path == "" {
		path = "config.json"
	}

	data, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return defaultConfig(), nil
		}
		return nil, fmt.Errorf("unable to read config: %w", err)
	}

	cfg := defaultConfig()
	if err := json.Unmarshal(data, cfg); err != nil {
		return nil, fmt.Errorf("unable to parse config: %w", err)
	}

	if cfg.TLSCertFile == "" || cfg.TLSKeyFile == "" {
		return nil, fmt.Errorf("config missing TLS certificate paths")
	}

	return cfg, nil
}

func defaultConfig() *Config {
	return &Config{
		TLSCertFile: "certs/daemon.crt",
		TLSKeyFile:  "certs/daemon.key",
		ServerAddr:  "0.0.0.0",
		Port:        8443,
	}
}
