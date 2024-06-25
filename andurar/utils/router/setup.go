package router

import (
	"log"

	"andurar.api/app/frameworks/http/routes/api"
	"github.com/gofiber/fiber/v3"
)

// Router interface with the Router method
type Router interface {
	Router(fiber *fiber.App)
}

// NewRouter initializes and sets up the router
func NewRouter(app *fiber.App, params ...any) {
	log.Println("Initializing new router with parameters:", params)
	apiRouter := api.NewAPIRouter(params)
	if apiRouter == nil {
		log.Fatal("Failed to initialize API router: apiRouter is nil")
	}
	setup(app, apiRouter)
}

// setup configures the given router
func setup(app *fiber.App, router Router) {
	log.Println("Setting up router")
	if router == nil {
		log.Fatal("Router is nil")
	}
	router.Router(app)
}
