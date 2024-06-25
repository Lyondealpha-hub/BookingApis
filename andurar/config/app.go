package config

import (
	"encoding/json"

	"github.com/gofiber/fiber/v3"
)

type appConfig struct {
	AppName      string
	ServerHeader string
	JSONEncoder  any
	JSONDecoder  any
}

func AppConfig() fiber.Config {

	return fiber.Config{
		AppName:      "Andurar Hotel Management Apis",
		ServerHeader: "Andurar Hotel Management Apis",
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
	}

}
