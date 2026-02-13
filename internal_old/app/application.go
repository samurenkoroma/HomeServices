package app

import (
	"fmt"
	"log"
	"samurenkoroma/services/configs"
	"samurenkoroma/services/pkg/db"
	"samurenkoroma/services/pkg/logger"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/rs/zerolog"
)

type AuthClaims struct {
	Email string
}

type Polevod struct {
	Logger *zerolog.Logger
	App    *fiber.App
	Config *configs.Config
	Db     *db.Db
}

func NewApplication(cfg *configs.Config, db *db.Db) *Polevod {
	app := fiber.New(fiber.Config{
		BodyLimit: 1000 * 1024 * 1024,
	})
	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: logger.NewLogger(cfg.Logger),
	}))

	app.Use(recover.New())
	// app.Static("/public", "./public")
	//app.Use(middleware.IsAuthenticated(cfg.Auth))

	application := Polevod{
		Logger: logger.NewLogger(cfg.Logger),
		App:    app,
		Config: cfg,
		Db:     db,
	}

	return &application

}

func (a *Polevod) Run() {
	log.Fatal(a.App.Listen(fmt.Sprintf("0.0.0.0%s", a.Config.Server.ApiPort)))
}

func NotImplement() error {
	return fiber.NewError(fiber.StatusInternalServerError, "Not implement")
}
