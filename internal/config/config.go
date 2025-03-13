package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	GRPCServerPort      string
	WebSocketServerPort string
}

func LoadConfig() (Config, error) {
	var config Config
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../")

	if err := viper.ReadInConfig(); err != nil {
		return config, fmt.Errorf("error reading configuration file: %v", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, fmt.Errorf("error unmarshalling configuration: %v", err)
	}

	return config, nil
}
