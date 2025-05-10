package config

type DatabaseConfig struct {
	Host     string `env:"HOST"`
	Port     int    `env:"PORT"`
	Username string `env:"USERNAME"`
	Password string `env:"PASSWORD"`
	Database string `env:"DATABASE"`
	SSLMode  string `env:"SSLMODE"`
}
