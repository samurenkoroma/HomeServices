package response

import "github.com/gofiber/fiber/v2"

type ErrResponse struct {
	Message    string
	StatusCode int
}

func JSON(ctx *fiber.Ctx, i any) error {
	return ctx.JSON(i)
}

func ERROR(ctx *fiber.Ctx, err error, statusCode int) error {
	ctx.Status(statusCode)
	return ctx.JSON(ErrResponse{
		StatusCode: statusCode,
		Message:    err.Error(),
	})
}
