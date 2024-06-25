package web

import (
	"github.com/gofiber/fiber/v3"
)

type webRouter struct{}

func NewWebRouter(params []any) *webRouter {
	instance := &webRouter{}
	return instance
}

func (r *webRouter) Router(app *fiber.App) {

	webGroup := app.Group("")
	r.index(webGroup)
	r.monitor(webGroup)
	app.Use(
		func(c fiber.Ctx) error {
			return c.SendStatus(404) // => 404 "Not Found"
		},
	)
}

func (r *webRouter) index(router fiber.Router) {
	router.Get(
		"/", func(c fiber.Ctx) error {
			return c.SendStatus(fiber.StatusOK)
		},
	)
}

func (r *webRouter) monitor(router fiber.Router) {
	// router.Get("/monitor", monitor.New())
}
