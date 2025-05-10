package config

type AuthConfig struct {
	SupabaseJWTSecretKey string `env:"JWT_SECRET_KEY"`
}
