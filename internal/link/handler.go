package link

import (
	"fmt"
	"net/http"
	"samurenkoroma/services/configs"
	"samurenkoroma/services/pkg/event"
	"samurenkoroma/services/pkg/middleware"
	"samurenkoroma/services/pkg/request"
	"samurenkoroma/services/pkg/response"
	"strconv"

	"gorm.io/gorm"
)

type LinkHandlerDeps struct {
	LinkRepository *LinkRepository
	Config         *configs.Config
	EventBus       *event.EventBus
}

type LinkHandler struct {
	LinkRepository *LinkRepository
	EventBus       *event.EventBus
}

func NewLinkHandler(router *http.ServeMux, deps LinkHandlerDeps) {
	handler := &LinkHandler{
		LinkRepository: deps.LinkRepository,
		EventBus:       deps.EventBus,
	}
	router.HandleFunc("POST /link", handler.Create())
	router.Handle("PATCH /link/{id}", middleware.IsAuthed(handler.Update(), deps.Config))
	router.HandleFunc("DELETE /link/{id}", handler.Delete())
	router.HandleFunc("GET /link/{hash}", handler.GoTo())
	router.HandleFunc("GET /link", handler.List())
}

func (handler *LinkHandler) List() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		cursor, err := strconv.Atoi(r.URL.Query().Get("cursor"))
		if err != nil {
			cursor = 0
		}
		limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
		if err != nil {
			limit = 10
		}

		links := handler.LinkRepository.GetList(LinkQueryParams{Limit: limit, Cursor: cursor})
		count := handler.LinkRepository.Count()

		response.Json(w, LinkListResponse{
			Data:  links,
			Count: count,
		}, http.StatusOK)
	}
}

func (handler *LinkHandler) GoTo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.PathValue("hash")
		link, err := handler.LinkRepository.GetByHash(hash)
		if err != nil {
			response.ErrJson(w, err.Error(), http.StatusNotFound)
			return
		}

		go handler.EventBus.Publish(event.Event{
			Type: event.EventLinkVisited,
			Data: link.ID,
		})
		// go handler.StatRepository.AddClick(link.ID)
		http.Redirect(w, r, link.Url, http.StatusTemporaryRedirect)
		response.Json(w, link, http.StatusOK)

	}
}
func (handler *LinkHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[LinkCreateRequest](&w, r)
		if err != nil {
			return
		}

		link := NewLink(body.Url)
		for {
			existedLink, _ := handler.LinkRepository.GetByHash(link.Hash)
			if existedLink == nil {
				break
			}
			link.GenerateHash()
		}
		createdLink, err := handler.LinkRepository.Create(link)
		if err != nil {
			response.ErrJson(w, err.Error(), http.StatusBadRequest)
			return
		}

		response.Json(w, createdLink, http.StatusCreated)
	}
}

func (handler *LinkHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email, ok := r.Context().Value(middleware.ContextEmailKey).(string)
		if ok {
			fmt.Printf("user : %s\n", email)
		}

		body, err := request.HandleBody[LinkUpdateRequest](&w, r)
		if err != nil {
			return
		}

		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)

		if err != nil {
			response.ErrJson(w, err.Error(), http.StatusBadRequest)
			return
		}

		link, err := handler.LinkRepository.Update(&Link{
			Model: gorm.Model{
				ID: uint(id),
			},
			Url:  body.Url,
			Hash: body.Hash,
		})
		if err != nil {
			response.ErrJson(w, err.Error(), http.StatusBadRequest)
			return
		}
		response.Json(w, link, http.StatusOK)
	}
}

func (handler *LinkHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)

		if err != nil {
			response.ErrJson(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := handler.LinkRepository.Exist(uint(id)); err != nil {
			response.ErrJson(w, err.Error(), http.StatusNotFound)
			return
		}
		err = handler.LinkRepository.Delete(uint(id))
		if err != nil {
			response.ErrJson(w, err.Error(), http.StatusInternalServerError)
			return
		}
		response.Json(w, nil, http.StatusNoContent)
	}
}
