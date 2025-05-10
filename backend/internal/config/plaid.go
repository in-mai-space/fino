package config

type PlaidConfig struct {
	ClientID     string `env:"CLIENT_ID"`
	Environment  string `env:"ENV"`
	Products     string `env:"PRODUCTS"`
	CountryCodes string `env:"COUNTRY_CODES"`
}
