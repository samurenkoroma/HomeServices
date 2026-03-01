package main

import (
	"fmt"
	"net/http"
	"samurenkoroma/services/configs"
	"samurenkoroma/services/internal/growing/application/command"
	"samurenkoroma/services/internal/growing/infrastructure/postgres"
	"samurenkoroma/services/internal/growing/interfaces/httpapi"
	"samurenkoroma/services/internal/infrastructure/eventbus"
)

func main() {
	conf := configs.LoadConfig()
	//database := db.NewDb(conf)
	//application := app.NewApplication(conf, database)

	conn, err := postgres.NewDatabase(conf.Db.Dsn)
	if err != nil {
		fmt.Println(err)
	}

	eventBus := eventbus.NewInMemoryEventBus()
	uowFactory := postgres.NewPgUnitOfWorkFactory(conn, eventBus)

	//createCmd := command.NewCreateLandUnitHandler(uowFactory)
	//getQuery := query.NewGetLandUnitHandler(conn)

	mux := http.NewServeMux()
	//handler := httpapi.NewLandUnitHTTPHandler(createCmd, getQuery)

	//handler.RegisterRoutes(mux)

	router := httpapi.NewCommandRouter()
	router.Register(
		"CreateField",
		&command.CreateFacilityHandler{UowFactory: uowFactory},
		command.DecodeCreateField,
	)
	commandHandler := httpapi.CommandEndpoint(router)

	mux.Handle("/api/command", httpapi.Chain(
		commandHandler,
		httpapi.TransactionMiddleware(uowFactory),
	))

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
