package config

import (
	"fmt"
	"log"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type EnvConfig struct {
	ListenPort string `env:"LISTEN_PORT"`
	TimeFormat string `env:"TIME_FORMAT"`
}

func (c *EnvConfig) PrintConfig() {
	fmt.Println("Listen Port: ", c.ListenPort)
	fmt.Println("TimeFormat: ", c.TimeFormat)
}

func NewEnvConfig(path string) Config {
	err := godotenv.Load(path)
	if err != nil {
		log.Fatal("Error when opening .env file: ", err)
	}

	config := EnvConfig{}
	err = env.Parse(&config)
	if err != nil {
		log.Fatal("Error during parsing .env file: ", err)
	}

	return &config
}
