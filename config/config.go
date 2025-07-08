package config

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Env    string
	Server struct {
		Port int
	}
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(os.Getenv("CONFIG_PATH"))

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
