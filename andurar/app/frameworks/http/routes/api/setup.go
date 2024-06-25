package api

import (
	"log"

	"andurar.api/ent"
	"github.com/gofiber/fiber/v3"
)

type apiRouter struct {
	Adapter *ent.Client
}

// NewAPIRouter creates a new instance of apiRouter and initializes it with params
func NewAPIRouter(params []any) *apiRouter {
	log.Println("Creating new API router with parameters:", params)
	instance := new(apiRouter)
	return instance.instantiate(params)
}

// Router sets up routes for the apiRouter
func (r *apiRouter) Router(app *fiber.App) {
	if r.Adapter == nil {
		log.Fatal("Database adapter is not initialized in apiRouter")
	}
	api := app.Group("/api")
	// adminApi := app.Group("/api/admin")

	// Client Booking
	ClientBooking(api, r)
}

// instantiate initializes the apiRouter with the given parameters
func (r *apiRouter) instantiate(params []any) *apiRouter {
	for _, param := range params {
		if adapter, ok := param.(*ent.Client); ok {
			r.Adapter = adapter
			log.Println("Database adapter initialized in apiRouter")
			continue
		}
	}
	if r.Adapter == nil {
		log.Println("Warning: Database adapter was not provided in parameters")
	}
	return r
}
