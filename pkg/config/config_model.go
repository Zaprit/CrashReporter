// Package config is the package that manages the config
package config

// LoadedConfig is the currently LoadedConfig, if used before LoadConfig is called all the fields will be blank
var LoadedConfig Config

// Path is the name of the config file
var Path = "config.toml"

// Config is the schema for all the config in CrashReporter
type Config struct {
	DiscordWebhook     string
	ListenAddress      string
	AdminUsers         []string
	OAuth2URL          string
	OAuth2ClientID     string
	OAuth2ClientSecret string
	DBFile             string
	PublicURL          string
	Debug              bool
	LighthouseURL      string
}
