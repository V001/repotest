package config

import (
	"github.com/caarlos0/env/v6"
)

// ENV - переменные окружения
// FILE - файлы
// Флаги - переменные из командной стркои
// внешние сервисы - Hashicorp Vault
type Config struct {
	HTTP     HTTP
	Port     int    `env:"PORT" envDefault:"9091"`
	Database string `env:"DATABASE" envDefault:"easdasd"`
	JWTKey   string `envconfig:"JWT_KEY" default:"supersecret"`
	PgURL    string `envconfig:"PG_URL" default:"user=root password=secret dbname=auth sslmode=disable host=localhost port=5432" `
}

type HTTP struct {
	PORT   int
	URL    string
	UseSSL bool
}

func New() (*Config, error) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
