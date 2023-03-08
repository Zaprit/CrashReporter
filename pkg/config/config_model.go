// Package config is the package that manages the config
package config

// LoadedConfig is the currently LoadedConfig, if used before LoadConfig is called all the fields will be blank
var LoadedConfig Config

// Path is the name of the config file
var Path = "config.toml"

// Config is the schema for all the config in CrashReporter
type Config struct {
	MigrateOld         bool     `comment:"Set true to import CrashHelperCore reports on startup\nWill load them from a db folder"`
	DiscordWebhook     string   `comment:"The webhook for posting reports in discord, leave blank to disable"`
	ListenAddress      string   `comment:"The address/port to listen on use \":8080\" to listen on all interfaces on port 8080, required"`
	AdminUsers         []string `comment:"An array of all admin users, at least on is required to administer the instance"`
	OAuth2ClientID     string   `comment:"The client ID of your GitHub application, required for OAuth2 login"`
	OAuth2ClientSecret string   `comment:"The client secret of your GitHub application, required for OAuth2 login"`
	DBFile             string   `comment:"The path to your database file, required"`
	PublicURL          string   `comment:"The public URL of your instance, used for cookies and discord links, not required but recommended"`
	Debug              bool     `comment:"Enable debug mode, this will log in a verbose manor, mainly intended for developers"`
	LighthouseURL      string   `comment:"The URL for your ProjectLighthouse instance, required for user lookup, leave blank to disable"`
}

var DefaultConfig = Config{
	DiscordWebhook:     "",
	ListenAddress:      ":8080",
	AdminUsers:         []string{"<YOUR GITHUB USERNAME>"},
	OAuth2ClientID:     "<YOUR GITHUB CLIENT ID>",
	OAuth2ClientSecret: "<YOUR GITHUB CLIENT ID>",
	DBFile:             "crash_reporter.db",
	PublicURL:          "https://crashreporter.example.com",
	Debug:              false,
	LighthouseURL:      "https://beacon.lbpunion.com",
}
