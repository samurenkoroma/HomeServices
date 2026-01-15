package home

import (
	"samurenkoroma/services/internal/app"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type HomeHandler struct {
	router fiber.Router
	logger zerolog.Logger
}

func NewHomeHandler(app *app.Polevod) {
	h := &HomeHandler{
		router: app.App,
		logger: *app.Logger,
	}

	main := h.router.Group("/main")
	main.Get("/", h.home)
	main.Get("/error", h.error)

}

func (h *HomeHandler) home(c *fiber.Ctx) error {
	return app.NotImplement()
}

func (h *HomeHandler) error(c *fiber.Ctx) error {
	return app.NotImplement()
}
