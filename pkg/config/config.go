package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

// base config struct
type Config struct {
	LogLevel string `envconfig:"LOG_LEVEL"` //LogLevel for logger can be: panic, fatal, error, warn, info, debug, trace
	DataServer
}

// dataServer config struct
type DataServer struct {
	Port string `envconfig:"DATA_SERVER_PORT"` //DataServer port
}

// GetConfig returns configuration or error
func GetConfig() (*Config, error) {
	var config Config
	if err := envconfig.Process("", &config); err != nil {
		return nil, fmt.Errorf("failed to get config: %v", err)
	}
	return &config, nil
}
