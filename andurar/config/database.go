package config

import (
	"os"
)

type (
	DBConfig struct {
		Driver   string
		Host     string
		Port     string
		Name     string
		Username string
		Password string
		SSL      string
	}
)

func DB() *DBConfig {
	return &DBConfig{
		Driver:   os.Getenv("DB_DRIVER"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_DATABASE"),
		SSL:      os.Getenv("DB_SSLMODE"),
	}
}
