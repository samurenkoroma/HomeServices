package rest

import (
	"samurenkoroma/services/internal/infrastructure/repo"
	"samurenkoroma/services/internal_old/app"
	"samurenkoroma/services/pkg/request"
	"samurenkoroma/services/services/accountant"
	"samurenkoroma/services/services/accountant/entity"

	"github.com/gofiber/fiber/v2"
)

type SupplierHandler struct {
	router  fiber.Router
	service *accountant.Service
}

func NewSupplierHandler(app *app.Polevod) {
	h := SupplierHandler{
		router:  app.App,
		service: accountant.NewSupplierService(repo.NewCrudRepo[entity.Supplier](app.Db)),
	}
	g := h.router.Group("/finance")
	g.Post("/supplier", h.Create)
	g.Get("/supplier", h.Filter)
	g.Delete("/supplier", h.Delete)

}

func (h *SupplierHandler) Create(c *fiber.Ctx) error {
	req, err := request.HandlerRequest[accountant.CreateSupplierRequest](c)

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
