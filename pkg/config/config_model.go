// Package config is the package that manages the config
package config

// LoadedConfig is the currently LoadedConfig, if used before LoadConfig is called all the fields will be blank
var LoadedConfig Config

// Path is the name of the config file
var Path = "config.toml"

type DatabaseConfig struct {
	Type     string `comment:"The type of database to use, valid options are mysql, postgres, sqlite. required"`
	Hostname string `comment:"The hostname for your database, required if using postgres/mysql/mariadb"`
	Port     uint16 `comment:"The database port, not required for sqlite"`

	Username string `comment:"The database username, not required for sqlite"`
	Password string `comment:"The database password, not required for sqlite"`

	Database string `comment:"The database to use, or the path to an SQLite database, required"`
}

// Config is the schema for all the config in CrashReporter
type Config struct {
	MigrateOld         bool     `comment:"Set true to import CrashHelperCore reports on startup\nWill load them from a db folder"`
	DiscordWebhook     string   `comment:"The webhook for posting reports in discord, leave blank to disable"`
	ListenAddress      string   `comment:"The address/port to listen on use \":8080\" to listen on all interfaces on port 8080, required"`
	AdminUsers         []string `comment:"An array of all admin users, at least on is required to administer the instance"`
	OAuth2ClientID     string   `comment:"The client ID of your GitHub application, required for OAuth2 login"`
	OAuth2ClientSecret string   `comment:"The client secret of your GitHub application, required for OAuth2 login"`
	PublicURL          string   `comment:"The public URL of your instance, used for cookies and discord links, not required but recommended"`
	Debug              bool     `comment:"Enable debug mode, this will log in a verbose manor, mainly intended for developers"`
	LighthouseURL      string   `comment:"The URL for your ProjectLighthouse instance, required for user lookup, leave blank to disable"`

	Database DatabaseConfig `comment:"The configuration for the database"`
}

var DefaultConfig = Config{
	ListenAddress:      ":8080",
	AdminUsers:         []string{"<YOUR GITHUB USERNAME>"},
	OAuth2ClientID:     "<YOUR GITHUB CLIENT ID>",
	OAuth2ClientSecret: "<YOUR GITHUB CLIENT ID>",

	PublicURL:     "https://crashreporter.example.com",
	Debug:         false,
	LighthouseURL: "https://beacon.lbpunion.com",

	Database: DatabaseConfig{
		Type:     "sqlite",
		Hostname: "",
		Port:     0,
		Username: "",
		Password: "",
		Database: "crash_reporter.sqlite",
	},
}
