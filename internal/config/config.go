package config

import (
	"github.com/caarlos0/env/v8"
	"go.uber.org/fx"
)

type Config struct {
	Version    string   `env:"VERSION" envDefault:"UNDEFINED"`
	HashSecret string   `env:"HASH_SECRET,required"`
	HTTPPort   int      `env:"HTTP_PORT" envDefault:"8000"`
	Database   Database `envPrefix:"DB_"`
	Redis      Redis    `envPrefix:"REDIS_"`
}

var version = "UNDEFINED"

// New builds a configuration struct that will be used by the application.
// It will be available to all of the application's dependencies.
func New(lc fx.Lifecycle) *Config {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		panic(err)
	}
	cfg.Version = version

	return &cfg
}
