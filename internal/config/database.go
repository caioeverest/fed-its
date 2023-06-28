package config

type Database struct {
	User     string `env:"USER" envDefault:"postgres"`
	Password string `env:"PASSWORD"`
	Name     string `env:"NAME" envDefault:"fedits"`
	Host     string `env:"HOST" envDefault:"localhost"`
	Port     int    `env:"PORT" envDefault:"5432"`
	SSLMode  string `env:"SSL_MODE"`
	TimeZone string `env:"TIME_ZONE" envDefault:"America/Sao_Paulo"`
}
