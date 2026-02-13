package auth

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"samurenkoroma/services/configs"
	"samurenkoroma/services/pkg/jwt"
	"samurenkoroma/services/pkg/request"
)

type AuthHandlerDeps struct {
	*AuthService
	Config configs.AuthConfig
}

type AuthHandler struct {
	*AuthService
	jwt *jwt.JWT
}

func NewAuthHandler(app *fiber.App, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		AuthService: deps.AuthService,
		jwt:         jwt.NewJWT(deps.Config),
	}
	router := app.Group("/auth")
	router.Post("/login", handler.Login)
	router.Post("/register", handler.Register)
	router.Post("/refresh", handler.Refresh)
}

func (handler *AuthHandler) Login(c *fiber.Ctx) error {
	body, err := request.HandlerRequest[LoginRequest](c)
	if err != nil {
		return err
	}

	email, err := handler.AuthService.Login(body.Email, body.Password)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, err.Error())
	}
	token, err := handler.jwt.Create(jwt.JWTData{Email: email})
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}
	//TODO вынести в горутину
	if err := handler.AuthService.updateRefreshToken(email, token.Refresh); err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	data := LoginResponse{
		RefreshToken: token.Refresh,
		AccessToken:  token.Access,
	}
	return c.JSON(data)
}

func (handler *AuthHandler) Register(c *fiber.Ctx) error {
	body, err := request.HandlerRequest[RegisterRequest](c)
	if err != nil {
		return err
	}

	email, err := handler.AuthService.Register(body.Email, body.Password, body.Name)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, err.Error())
	}
	token, err := handler.jwt.Create(jwt.JWTData{Email: email})
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}
	//TODO вынести в горутину
	if err := handler.AuthService.updateRefreshToken(email, token.Refresh); err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	data := RegisterResponse{
		RefreshToken: token.Refresh,
		AccessToken:  token.Access,
	}
	return c.JSON(data)
}

func (handler *AuthHandler) Refresh(c *fiber.Ctx) error {
	refresh := c.Get("Refresh-Token")
	isValid, data := handler.jwt.ParseRefresh(refresh)
	if !isValid {
		return fiber.NewError(http.StatusBadRequest, "wrong refresh token")
	}

	token, err := handler.jwt.Create(jwt.JWTData{Email: data.Email})
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}
	//TODO вынести в горутину
	if err := handler.AuthService.updateRefreshToken(data.Email, token.Refresh); err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(RegisterResponse{
		RefreshToken: token.Refresh,
		AccessToken:  token.Access,
	})
}
