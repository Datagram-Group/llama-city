package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	GRPCServerPort      string
	WebSocketServerPort string
}

func LoadConfig(configFilePath string) (*Config, error) {
	viper.SetConfigFile(configFilePath)

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &config, nil
}
