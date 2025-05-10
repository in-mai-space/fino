package config

type NotionConfig struct {
	ClientID          string `env:"CLIENT_ID"`
	RedirectURI       string `env:"REDIRECT_URI"`
	AuthorizationCode string `env:"AUTHORIZATION_CODE"`
}
