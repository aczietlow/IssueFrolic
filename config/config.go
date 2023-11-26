package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Token    string
	Username string `json:"user"`
}

func LoadConfig(filepath string) (*Config, error) {
	configFile, err := os.ReadFile(filepath)

	if err != nil {
		fmt.Printf("%v\r\n", "Error reading config file.")
		return nil, err
	}
	var config Config
	if err := json.Unmarshal(configFile, &config); err != nil {
		return nil, err
	}

	return &config, nil

}
