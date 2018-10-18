package config

import (
	"encoding/json"
	"io/ioutil"
)

// Project configuration properties
// Engine needs to know where to locate its game data
type Config struct {
	GameDirectory string
}

var config Config

// Get
func Get() *Config {
	return &config
}

func Load() (*Config, error) {
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		return &config, err
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		return &config, err
	}

	return &config, nil
}
