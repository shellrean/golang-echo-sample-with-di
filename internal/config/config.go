package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/samber/do/v2"
)

type Database struct {
	Dsn string
}

type Config struct {
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
	}, nil
}
