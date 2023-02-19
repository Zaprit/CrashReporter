package config

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
