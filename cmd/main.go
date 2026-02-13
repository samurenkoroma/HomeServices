package main

import (
	"samurenkoroma/services/configs"
	"samurenkoroma/services/internal/delivery/rest"
	"samurenkoroma/services/internal/infrastructure/repo"
	"samurenkoroma/services/internal_old/app"
	"samurenkoroma/services/internal_old/auth"
	"samurenkoroma/services/pkg/db"
)

func main() {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	application := app.NewApplication(conf, db)

	//Репозитории
	userRepo := repo.NewUserRepo(db)
	bookRepo := repo.NewBookRepo(db)

	//Сервисы
	authService := auth.NewAuthService(userRepo)
	auth.NewAuthHandler(application.App, auth.AuthHandlerDeps{
		AuthService: authService,
		Config:      conf.Auth,
	})
	rest.NewSupplierHandler(application)

	rest.BookRouter(application.App, bookRepo, conf)

	//home.NewHomeHandler(application)
	// pages.NewPageHandler(application)
	// weather.NewWeatherHandler(application)
	// finance.NewFinanceHandler(application)

	application.Run()
}
