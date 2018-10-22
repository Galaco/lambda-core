package config

import (
	"encoding/json"
	"io/ioutil"
)


const minWidth = 320
const minHeight = 240

// Project configuration properties
// Engine needs to know where to locate its game data
type Config struct {
	GameDirectory string
	Video struct {
		Width int
		Height int
	}
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

	validate()

	return &config, nil
}

func validate() {
	if config.Video.Width < minWidth {
		config.Video.Width = minWidth
	}

	if config.Video.Height < minHeight {
		config.Video.Height = minHeight
	}
}