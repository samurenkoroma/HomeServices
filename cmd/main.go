package main

import (
	"fmt"
	"net/http"
	"samurenkoroma/services/configs"
	"samurenkoroma/services/internal/field/application/command"
	"samurenkoroma/services/internal/field/application/query"
	"samurenkoroma/services/internal/field/infrastructure/postgres"
	myRoute "samurenkoroma/services/internal/field/interfaces/http"
)

func main() {
	conf := configs.LoadConfig()
	//database := db.NewDb(conf)
	//application := app.NewApplication(conf, database)

	conn, err := postgres.NewDatabase(conf.Db.Dsn)
	if err != nil {
		fmt.Println(err)
	}
	uow, err := postgres.NewUow(conn)
	if err != nil {
		return
	}

	createCmd := &command.CreateLandUnitHandler{Uow: uow}
	getQuery := query.NewGetLandUnitHandler(conn)

	mux := http.NewServeMux()
	handler := myRoute.NewLandUnitHTTPHandler(createCmd, getQuery)

	handler.RegisterRoutes(mux)
	http.ListenAndServe(":8080", mux)
	//Репозитории
	//userRepo := repo.NewUserRepo(database)
	//bookRepo := repo.NewBookRepo(database)

	//Сервисы
	//authService := auth2.NewAuthService(userRepo)
	//auth2.NewAuthHandler(application.App, auth2.AuthHandlerDeps{
	//	AuthService: authService,
	//	Config:      conf.Auth,
	//})
	//rest.NewSupplierHandler(application)
	//rest.NewStoreHouseHandler(application)
	//rest.BookRouter(application.App, bookRepo, conf)
	//rest.NewPlantHandler(application.App)
	//home.NewHomeHandler(application)
	// pages.NewPageHandler(application)
	// weather.NewWeatherHandler(application)
	// finance.NewFinanceHandler(application)

	//application.Run()
}
