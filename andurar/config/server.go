package config

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
)

type Server struct {
	HTTP *fiber.App
}

// Constructor

func NewServer() *Server {
	// appConfig := AppConfig()
	return &Server{
		HTTP: fiber.New(),
	}
}

func (s *Server) Run() {

	fmt.Println("I want to connect to run")

	s.HTTP.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello World, from Server")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "80" // Default to port 80 if PORT environment variable is not set
	}

	log.Fatal(s.HTTP.Listen(":" + port))
	// err := s.HTTP.Listen(":" + port)

}
