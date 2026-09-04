package configs

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Db   DBConfig
	Auth AuthConfig
}

type DBConfig struct {
	Dsn string
}

type AuthConfig struct {
	Secret string
}

func LoadConfig() (*Config, error) {
	_ = godotenv.Load()

	cfg := &Config{
		Db: DBConfig{
			Dsn: os.Getenv("DSN"),
		},
		Auth: AuthConfig{
			Secret: os.Getenv("TOKEN"),
		},
	}

	if cfg.Db.Dsn == "" {
		return nil, errors.New("DSN is required")
	}

	if cfg.Auth.Secret == "" {
		return nil, errors.New("TOKEN is required")
	}

	return cfg, nil
}
