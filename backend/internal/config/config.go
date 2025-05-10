package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	AuthConfig     `envPrefix:"SUPABASE_"`
	DatabaseConfig `envPrefix:"DB_"`
	PlaidConfig    `envPrefix:"PLAID_"`
	NotionConfig   `envPrefix:"NOTION_"`
	AppConfig      `envPrefix:"APP_"`
}

func GetConfigurations(path string) (*Config, error) {
	if path != "" {
		if err := godotenv.Load(path); err != nil {
			return nil, fmt.Errorf("Failed to load environment variables: %s", err.Error())
		}
	}

	cfg, err := env.ParseAs[Config]()
	if err != nil {
		return nil, fmt.Errorf("Failed to parse environment variables: %s", err.Error())
	}

	return &cfg, nil
}
