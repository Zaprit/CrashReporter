package config

import (
	"github.com/pelletier/go-toml/v2"
	"os"
)

func LoadConfig() error {
	data, err := os.ReadFile(Path)
	if err != nil {
		return err
	}
	var config Config
	err = toml.Unmarshal(data, &config)
	if err != nil {
		return err
	}
	LoadedConfig = config
	return nil
}
