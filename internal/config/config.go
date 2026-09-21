package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/samber/do/v2"
)

type Database struct {
	Dsn string
}

type Server struct {
	Port string
	Host string
}

type Config struct {
	Server
	Database
}

func Load(i do.Injector) (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	return &Config{
		Database: Database{
			Dsn: os.Getenv("DATABASE_URL"),
		},
		Server: Server{
			Host: os.Getenv("SERVER_HOST"),
			Port: os.Getenv("SERVER_PORT"),
		},
	}, nil
}
