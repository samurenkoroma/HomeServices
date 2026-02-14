package main

import (
	"samurenkoroma/services/configs"
	"samurenkoroma/services/internal/delivery/rest"
	"samurenkoroma/services/internal/infrastructure/repo"
	"samurenkoroma/services/internal_old/app"
	"samurenkoroma/services/pkg/db"
	auth2 "samurenkoroma/services/services/auth"
)

func main() {
	conf := configs.LoadConfig()
	database := db.NewDb(conf)
	application := app.NewApplication(conf, database)

	//Репозитории
	userRepo := repo.NewUserRepo(database)
	bookRepo := repo.NewBookRepo(database)

	//Сервисы
	authService := auth2.NewAuthService(userRepo)
	auth2.NewAuthHandler(application.App, auth2.AuthHandlerDeps{
		AuthService: authService,
		Config:      conf.Auth,
	})
	rest.NewSupplierHandler(application)
	rest.NewStoreHouseHandler(application)
	rest.BookRouter(application.App, bookRepo, conf)

	//home.NewHomeHandler(application)
	// pages.NewPageHandler(application)
	// weather.NewWeatherHandler(application)
	// finance.NewFinanceHandler(application)

	application.Run()
}
