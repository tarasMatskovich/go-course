package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

var (
	Configuration Config
)

type Config struct {
	Port       string `env:"LISTEN_PORT"`
	TimeFormat string `env:"TIME_FORMAT"`
}

func New(configPath string) (*Config, error) {
	err := godotenv.Load(configPath)
	if err != nil {
		return nil, err
	}

	config := Config{}
	err = env.Parse(&config)
	if err != nil {
		return nil, err
	}

	Configuration = config

	return &config, nil
}
