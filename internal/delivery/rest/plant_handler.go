package rest

import (
	"github.com/gofiber/fiber/v2"
)

func NewPlantHandler(app fiber.Router) {
	h := PlantHandler{
		router: app,
	}
	g := h.router.Group("/plants")
	g.Get("/families", h.ListFamilies)
	g.Get("/families/:id", h.GetOneFamilia)

	g.Get("/families/:fId/species", h.ListSpeciesByFamilia)
	g.Get("/families/:fId/species/:id", h.GetOneSpecies)
	g.Get("/:id", h.GetPlant)
}

type PlantHandler struct {
	router fiber.Router
}

func (h PlantHandler) GetPlant(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"route": "/plants",
		"id":    ctx.Params("id"),
	})
}

func (h PlantHandler) ListFamilies(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"data": []fiber.Map{
			{"id": 1, "name": "Паслёновые"},
			{"id": 2, "name": "Тыквенные"},
		},
	})
}

func (h PlantHandler) GetOneFamilia(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"route": "/families/:id",
		"id":    ctx.Params("id"),
	})
}

func (h PlantHandler) ListSpeciesByFamilia(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"data": []fiber.Map{
			{"id": 1, "name": "Картофель"},
			{"id": 2, "name": "Баклажан"},
			{"id": 2, "name": "Томат"},
		},
	})
}

func (h PlantHandler) GetOneSpecies(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"route": "/families/:fId/species/:id",
		"fid":   ctx.Params("fId"),
		"id":    ctx.Params("id"),
	})
}
