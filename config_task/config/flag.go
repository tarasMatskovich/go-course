package config

import (
	"flag"
	"fmt"
)

type FlagConfig struct {
	ListenPort string
	TimeFormat string
}

func (c *FlagConfig) PrintConfig() {
	fmt.Println("Listen Port: ", c.ListenPort)
	fmt.Println("TimeFormat: ", c.TimeFormat)
}

func NewFlagConfig() Config {
	ListenPort := flag.String("listen_port", "", "a string")
	TimeFormat := flag.String("time_format", "", "a string")
	flag.Parse()

	config := FlagConfig{
		ListenPort: *ListenPort,
		TimeFormat: *TimeFormat,
	}

	return &config
}
