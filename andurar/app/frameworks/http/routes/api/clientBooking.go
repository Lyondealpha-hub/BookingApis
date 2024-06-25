package api

import (
	"andurar.api/app/frameworks/http/handlers/api"
	"github.com/gofiber/fiber/v3"
)

func ClientBooking(r fiber.Router, router *apiRouter) {

	handler := api.NewBookingHandler(router.Adapter)

	booking := r.Group("/bookings")
	booking.Get("", handler.FetchAll())
	booking.Post("", handler.Create())

}
