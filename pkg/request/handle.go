package request

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"samurenkoroma/services/pkg/response"
)

func HandleBody[T any](w *http.ResponseWriter, r *http.Request) (*T, error) {

	body, err := decode[T](r.Body)

	if err != nil {
		response.ErrJson(*w, err.Error(), http.StatusBadRequest)
		return nil, err
	}

	err = isValid(body)
	if err != nil {
		response.ErrJson(*w, err.Error(), http.StatusBadRequest)
		return nil, err
	}
	return &body, nil
}

func HandlerRequest[T any](ctx *fiber.Ctx) (*T, error) {
	var payload T
	if err := ctx.BodyParser(&payload); err != nil {
		return &payload, err
	}
	return &payload, nil
}
