package rest

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"path"
	"samurenkoroma/services/configs"
	"samurenkoroma/services/internal/infrastructure/repo"
	"samurenkoroma/services/pkg/response"
	"samurenkoroma/services/services/homelib"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type BookHandler struct {
	cfg  *configs.ServerConfig
	repo repo.BookRepository
}

func NewBookHandler(repo repo.BookRepository, cfg *configs.Config) *BookHandler {
	return &BookHandler{
		repo: repo,
		cfg:  &cfg.Server,
	}
}
func BookRouter(router fiber.Router, repo repo.BookRepository, cfg *configs.Config) {
	booksGroup := router.Group("/books")

	handler := NewBookHandler(repo, cfg)
	booksGroup.Get("", handler.GetList)
	booksGroup.Post("", handler.Create)
	booksGroup.Get("/:id", handler.GetOne)
	router.Get("resource/:id", handler.GetResource)
}

func (h *BookHandler) Create(ctx *fiber.Ctx) error {
	var dto homelib.BookRequest
	if err := ctx.BodyParser(&dto); err != nil {
		log.Println(err, ctx.Request())
		return response.ERROR(ctx, errors.New("wrong data"), http.StatusBadRequest)
	}

	var authors []homelib.Author
	for _, p := range dto.Authors {
		a := homelib.Author{
			Name: p,
		}

		authors = append(authors, a)
	}

	book, err := h.repo.Create(&homelib.Book{
		Title:   dto.Title,
		Authors: authors,
	})

	if err != nil {
		return response.ERROR(ctx, err, http.StatusConflict)
	}

	return response.JSON(ctx, h.makeBookResponse(*book))

}
func (h *BookHandler) GetList(ctx *fiber.Ctx) error {
	var params = repo.NewBookQueryParams()

	if err := ctx.QueryParser(params); err != nil {
		return response.ERROR(ctx, err, http.StatusBadRequest)
	}
	var books = h.repo.GetList(params)

	var data []homelib.BookResponse
	for _, b := range books {
		data = append(data, h.makeBookResponse(b))
	}

	return response.JSON(ctx, homelib.BookListResponse{Data: data, Count: len(books)})
}

func (h *BookHandler) GetOne(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id")

	book, err := h.repo.GetById(uint(id))
	if err != nil {
		return response.ERROR(ctx, err, http.StatusNotFound)
	}

	return response.JSON(ctx, h.makeBookResponse(*book))
}

func (h *BookHandler) GetResource(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id")

	resource, err := h.repo.GetResourceById(uint(id))
	if err != nil {
		return response.ERROR(ctx, err, http.StatusNotFound)
	}
	ctx.Set(fiber.HeaderContentDisposition, `attachment; filename="`+strconv.Quote(resource.File)+`"`)
	return ctx.Download(path.Join(h.cfg.StorageDir, resource.File))
}

func (h *BookHandler) makeBookResponse(book homelib.Book) homelib.BookResponse {
	var resources []homelib.ResourceResponse
	var authors []homelib.AuthorResponse

	if len(book.Resources) != 0 {
		for _, r := range book.Resources {
			resources = append(resources, homelib.ResourceResponse{
				Type: uint(r.Type),
				Link: fmt.Sprintf("http://%s/resource/%d", h.cfg.ApiHost, r.ID),
			})
		}
	}

	if len(book.Authors) != 0 {
		for _, a := range book.Authors {
			authors = append(authors, homelib.AuthorResponse{
				Id:   a.ID,
				Name: a.Name,
			})
		}
	}
	return homelib.BookResponse{
		Title:     book.Title,
		Authors:   authors,
		Resources: resources,
		Id:        book.ID,
	}
}
