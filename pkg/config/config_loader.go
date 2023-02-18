package config

import (
    "github.com/pelletier/go-toml/v2"
    "os"
)

var LoadedConfig Config

func LoadConfig(path string) error {
    data, err := os.ReadFile(path)
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