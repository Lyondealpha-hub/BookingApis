package api

import (
	"fmt"

	"andurar.api/app/adapters/gateways"
	"andurar.api/app/adapters/presenters"
	applicaiton "andurar.api/app/application"
	"andurar.api/app/application/booking"
	requestdto "andurar.api/app/domain/request_dto"
	"andurar.api/ent"
	"github.com/gofiber/fiber/v3"
)

type bookingHandler struct {
	service gateways.BookingService
}

func NewBookingHandler(db *ent.Client) *bookingHandler {
	return &bookingHandler{service: booking.NewService(booking.NewRepository(db))}
}

func (h *bookingHandler) FetchAll() fiber.Handler {
	return func(c fiber.Ctx) error {
		result, err := h.service.FetchAll()

		if err != nil {
			return c.SendString("error")
		}

		return c.Status(fiber.StatusOK).JSON(result)
	}
}

func (h *bookingHandler) Create() fiber.Handler {
	return func(c fiber.Ctx) error {
		request := new(requestdto.BookingRequest)

		if err := applicaiton.BodyParser(c, request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ErrorResponse(fmt.Errorf("bad request")))
		}
		result, err := h.service.Create(request)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(result)
	}
}
