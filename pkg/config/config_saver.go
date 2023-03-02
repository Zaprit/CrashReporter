package config

import (
	"github.com/pelletier/go-toml/v2"
	"os"
)

// SaveConfig writes LoadedConfig to the file at Path in the TOML format
func SaveConfig() error {

	configData, err := toml.Marshal(LoadedConfig)
	if err != nil {
		return err
	}

	err = os.WriteFile(Path, configData, 0600)
	if err != nil {
		return err
	}

	return nil
}
