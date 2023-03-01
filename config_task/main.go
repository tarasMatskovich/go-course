package main

import (
	"config_task/config"
	"fmt"
)

func main() {
	configType := "json"

	fmt.Println("Config type: ", configType)

	switch configType {
	case "json":
		ParseJsonConfig()
	case "yaml":
		ParseYamlConfig()
	case "env":
		ParseEnvConfig()
	case "flag":
		ParseFlagConfig()
	default:
		ParseJsonConfig()
	}
}

func ParseJsonConfig() {
	config := config.NewJsonConfig("./config/config.json")

	config.PrintConfig()
}

func ParseYamlConfig() {
	config := config.NewYamlConfig("./config/config.yaml")

	config.PrintConfig()
}

func ParseEnvConfig() {
	config := config.NewEnvConfig("./config/config.env")

	config.PrintConfig()
}

func ParseFlagConfig() {
	config := config.NewFlagConfig()

	config.PrintConfig()
}
