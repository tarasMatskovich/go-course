package config

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type YamlConfig struct {
	ListenPort string `yaml:"ListenPort"`
	TimeFormat string `yaml:"TimeFormat"`
}

func (c *YamlConfig) PrintConfig() {
	fmt.Println("Listen Port: ", c.ListenPort)
	fmt.Println("TimeFormat: ", c.TimeFormat)
}

func NewYamlConfig(path string) Config {
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	config := YamlConfig{}
	err = yaml.Unmarshal(content, &config)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	return &config
}
