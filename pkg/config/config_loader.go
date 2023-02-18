package config

import (
    "github.com/pelletier/go-toml/v2"
    "os"
)

func LoadConfig(path string) {
    data, err := os.ReadFile(path)
    if err != nil {
        
    }
    toml.Unmarshal()
}