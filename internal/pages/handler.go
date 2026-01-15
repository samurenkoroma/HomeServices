package pages

import (
	"samurenkoroma/services/internal/app"
	"samurenkoroma/services/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

type PageHandler struct {
	router fiber.Router
}

func NewPageHandler(app *app.Polevod) {
	h := &PageHandler{
		router: app.App,
	}
	h.router.Get("/", h.home)
	h.router.Get("/land-plots", middleware.JWTProtected(app.Config.Auth), h.landPlots)
	h.router.Get("/market", h.market)
	h.router.Get("/finance", h.finance)
	h.router.Get("/tools", h.tools)
	h.router.Get("/work-log", h.worlLog)

}

func (h *PageHandler) home(c *fiber.Ctx) error {
	// component := views.Main()
	return nil
}

func (h *PageHandler) landPlots(c *fiber.Ctx) error {
	user := c.Locals("email")
	return c.JSON(user)
}

func (h *PageHandler) finance(c *fiber.Ctx) error {
	// component := pages.Finance()
	return nil
}

func (h *PageHandler) market(c *fiber.Ctx) error {
	// component := pages.Market()
	return nil
}

func (h *PageHandler) tools(c *fiber.Ctx) error {
	// component := pages.Tools()
	return nil
}
func (h *PageHandler) worlLog(c *fiber.Ctx) error {
	// component := pages.WorkLog()
	return nil //tadapter.Render(c, component)
}
