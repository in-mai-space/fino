package config

type AppConfig struct {
	Port string `env:"PORT"`
	Host string `env:"HOST"`
}
