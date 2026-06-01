package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Db   DBconfig
	Auth AuthConfig
}

type DBconfig struct {
	Dsn string
}

type AuthConfig struct {
	Secret string
}

//TODO refactor this

func LoadConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Print("Error loading .env file. Now using default config")
	}
	return &Config{
		DBconfig{
			os.Getenv("DSN"),
		},
		AuthConfig{
			os.Getenv("TOKEN"),
		},
	}
}
