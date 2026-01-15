package routes

import (
	"samurenkoroma/services/api/handlers"
	"samurenkoroma/services/configs"
	"samurenkoroma/services/pkg/repositories"

	"github.com/gofiber/fiber/v2"
)

func BookRouter(router fiber.Router, repo repositories.BookRepository, cfg *configs.Config) {
	booksGroup := router.Group("/books")

	handler := handlers.NewBookhandler(repo, cfg)
	booksGroup.Get("", handler.GetList())
	booksGroup.Post("", handler.Create())
	booksGroup.Get("/:id", handler.GetOne())
	router.Get("resource/:id", handler.GetResource())
}
