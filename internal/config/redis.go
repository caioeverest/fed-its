package config

type Redis struct {
	Addr     string `env:"ADDR" envDefault:"localhost:6379"`
	Username string `env:"USERNAME"`
	Password string `env:"PASSWORD"`
	DB       int    `env:"DB" envDefault:"0"`
}
