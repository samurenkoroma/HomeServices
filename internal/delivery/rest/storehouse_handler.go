package rest

import (
	"samurenkoroma/services/internal_old/app"
	"samurenkoroma/services/pkg/request"
	"samurenkoroma/services/services/storehouse"
	"samurenkoroma/services/services/storehouse/service"

	"github.com/gofiber/fiber/v2"
)

type StoreHouseHandler struct {
	router        fiber.Router
	seedService   *service.SeedService
	vendorService *service.VendorService
}

func NewStoreHouseHandler(app *app.Polevod) {
	h := StoreHouseHandler{
		router:        app.App,
		seedService:   service.NewSeedService(storehouse.NewSeedsRepository(app.Db)),
		vendorService: service.NewVendorService(storehouse.NewVendorRepository(app.Db)),
	}
	g := h.router.Group("/storehouse")
	g.Post("/seed", h.AddSeed)
	g.Get("/seed", h.ListSeed)
	g.Get("/vendor", h.ListVendor)
	g.Post("/vendor", h.AddVendor)
}

func (h StoreHouseHandler) AddSeed(ctx *fiber.Ctx) error {
	req, err := request.HandlerRequest[storehouse.CreateSeedRequest](ctx)

	if err != nil {
		return err
	}

	response, err := h.seedService.Add(req)
	if err != nil {
		return err
	}
	ctx.Status(fiber.StatusCreated)
	return ctx.JSON(fiber.Map{"data": response})
}

func (h StoreHouseHandler) ListSeed(ctx *fiber.Ctx) error {

	response, err := h.seedService.List()
	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{"data": response})
}

func (h StoreHouseHandler) AddVendor(ctx *fiber.Ctx) error {
	req, err := request.HandlerRequest[storehouse.CreateVendorRequest](ctx)

	if err != nil {
		return err
	}

	response, err := h.vendorService.Add(req)
	if err != nil {
		return err
	}
	ctx.Status(fiber.StatusCreated)
	return ctx.JSON(fiber.Map{"data": response})
}

func (h StoreHouseHandler) ListVendor(ctx *fiber.Ctx) error {

	response, err := h.vendorService.List()
	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{"data": response})
}
