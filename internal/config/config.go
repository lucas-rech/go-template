package config

import (
	"log/slog"

	"github.com/caarlos0/env/v11"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

var validate = validator.New()

var Version = "dev"

type Envs struct {
	AppName    string `env:"APP_NAME" envDefault:"go-template"`
	AppEnv     string `env:"APP_ENV" envDefault:"dev"`
	ServerPort string `env:"SERVER_PORT" envDefault:"8080" validate:"numeric"`

	LogLevel  slog.Level `env:"LOG_LEVEL" envDefault:"info"`
	LogFormat string     `env:"LOG_FORMAT" envDefault:"json"`
}

func Load() (*Envs, error) {
	_ = godotenv.Load()

	cfg := &Envs{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	if err := validate.Struct(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
