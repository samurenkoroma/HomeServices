package main

import (
	"samurenkoroma/services/configs"
	"samurenkoroma/services/internal/app"
	"samurenkoroma/services/internal/auth"
	"samurenkoroma/services/internal/user"
	"samurenkoroma/services/pkg/db"
	"samurenkoroma/services/pkg/repositories"
	"samurenkoroma/services/pkg/routes"
)

func main() {
	conf := configs.LoadConfig()
	database := db.NewDb(conf)
	application := app.NewApplication(conf, database)

	//Репозитории
	userRepo := user.NewUserRepo(database)
	bookRepo := repositories.NewBookRepo(database)

	//Сервисы
	authService := auth.NewAuthService(userRepo)

	auth.NewAuthHandler(application.App, auth.AuthHandlerDeps{
		AuthService: authService,
		Config:      conf.Auth,
	})

	routes.BookRouter(application.App, bookRepo, conf)

	//home.NewHomeHandler(application)
	// pages.NewPageHandler(application)
	// weather.NewWeatherHandler(application)
	// finance.NewFinanceHandler(application)

	application.Run()
}
