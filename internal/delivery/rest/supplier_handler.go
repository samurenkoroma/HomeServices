package rest

import (
	"samurenkoroma/services/internal/infrastructure/payload"
	"samurenkoroma/services/internal/infrastructure/repo"
	"samurenkoroma/services/internal/infrastructure/use_case"
	"samurenkoroma/services/internal_old/app"
	"samurenkoroma/services/pkg/request"

	"github.com/gofiber/fiber/v2"
)

type SupplierHandler struct {
	router  fiber.Router
	service *use_case.Service
}

func NewSupplierHandler(app *app.Polevod) {
	h := SupplierHandler{
		router:  app.App,
		service: use_case.NewSupplierService(repo.NewSupplierRepo(app.Db)),
	}
	g := h.router.Group("/finance")
	g.Post("/supplier", h.Create)
	g.Get("/supplier", h.Filter)
	g.Delete("/supplier", h.Delete)

}

func (h *SupplierHandler) Create(c *fiber.Ctx) error {
	req, err := request.HandlerRequest[payload.CreateSupplierRequest](c)

	if err != nil {
		return err
	}

	response, err := h.service.Create(req)
	if err != nil {
		return err
	}

	c.Status(fiber.StatusCreated)
	return c.JSON(fiber.Map{"data": response})
}

func (h *SupplierHandler) Filter(c *fiber.Ctx) error {

	response, err := h.service.List()
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"data": response})
}

func (h *SupplierHandler) Delete(ctx *fiber.Ctx) error {
	ctx.Get("id")
	return nil
}
