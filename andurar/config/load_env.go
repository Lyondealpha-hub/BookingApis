package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var Env map[string]string

func LoadEnvVariables() {
	if os.Getenv("ENVIRONMENT") == "Development" {
		err := godotenv.Load()

		if err != nil {
			log.Fatal("Error loading environment variables", err)
		}
	}
}

func Setup() {
	if os.Getenv("ENVIRONMENT") != "production" {
		envFile := ".env"
		var err error
		Env, err = godotenv.Read(envFile)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
