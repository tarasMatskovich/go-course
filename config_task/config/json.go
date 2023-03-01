package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type JsonConfig struct {
	ListenPort string `json:"ListenPort"`
	TimeFormat string `json:"TimeFormat"`
}

func (c *JsonConfig) PrintConfig() {
	fmt.Println("Listen Port: ", c.ListenPort)
	fmt.Println("TimeFormat: ", c.TimeFormat)
}

func NewJsonConfig(path string) Config {
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	config := JsonConfig{}
	err = json.Unmarshal(content, &config)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	return &config
}
