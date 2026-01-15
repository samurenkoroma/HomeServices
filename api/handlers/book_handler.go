package handlers

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"path"
	"samurenkoroma/services/configs"
	"samurenkoroma/services/pkg/entities"
	"samurenkoroma/services/pkg/payloads"
	"samurenkoroma/services/pkg/repositories"
	"samurenkoroma/services/pkg/response"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type BookHandler struct {
	cfg  *configs.ServerConfig
	repo repositories.BookRepository
}

func NewBookhandler(repo repositories.BookRepository, cfg *configs.Config) *BookHandler {
	return &BookHandler{
		repo: repo,
		cfg:  &cfg.Server,
	}
}

func (h *BookHandler) Create() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var payload payloads.BookRequest
		if err := ctx.BodyParser(&payload); err != nil {
			log.Println(err, ctx.Request())
			return response.ERROR(ctx, errors.New("wrong data"), http.StatusBadRequest)
		}

		var authors = []entities.Author{}
		for _, p := range payload.Authors {
			a := entities.Author{
				Name: p,
			}

			authors = append(authors, a)
		}

		book, err := h.repo.Create(&entities.Book{
			Title:   payload.Title,
			Authors: authors,
		})

		if err != nil {
			return response.ERROR(ctx, err, http.StatusConflict)
		}

		return response.JSON(ctx, h.makeBookResponse(*book))

	}
}
func (h *BookHandler) GetList() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params = repositories.NewBookQueryParams()

		if err := ctx.QueryParser(params); err != nil {
			return response.ERROR(ctx, err, http.StatusBadRequest)
		}
		var books = h.repo.GetList(params)

		var data = []payloads.BookResponse{}
		for _, b := range books {
			data = append(data, h.makeBookResponse(b))
		}

		return response.JSON(ctx, payloads.BookListResponse{Data: data, Count: len(books)})
	}
}

func (h *BookHandler) GetOne() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id, _ := ctx.ParamsInt("id")

		book, err := h.repo.GetById(uint(id))
		if err != nil {
			return response.ERROR(ctx, err, http.StatusNotFound)
		}

		return response.JSON(ctx, h.makeBookResponse(*book))
	}
}

func (h *BookHandler) GetResource() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id, _ := ctx.ParamsInt("id")

		resource, err := h.repo.GetResourceById(uint(id))
		if err != nil {
			return response.ERROR(ctx, err, http.StatusNotFound)
		}
		ctx.Set(fiber.HeaderContentDisposition, `attachment; filename="`+strconv.Quote(resource.File)+`"`)
		return ctx.Download(path.Join(h.cfg.StorageDir, resource.File))
	}
}

func (h *BookHandler) makeBookResponse(book entities.Book) payloads.BookResponse {
	var resources = []payloads.ResourceResponse{}
	var authors = []payloads.AuthorResponse{}
	if len(book.Resources) != 0 {
		for _, r := range book.Resources {
			resources = append(resources, payloads.ResourceResponse{
				Type: uint(r.Type),
				Link: fmt.Sprintf("http://%s/resource/%d", h.cfg.ApiHost, r.ID),
			})
		}
	}

	if len(book.Authors) != 0 {
		for _, a := range book.Authors {
			authors = append(authors, payloads.AuthorResponse{
				Id:   a.ID,
				Name: a.Name,
			})
		}
	}
	return payloads.BookResponse{
		Title:     book.Title,
		Authors:   authors,
		Resources: resources,
		Id:        book.ID,
	}
}
