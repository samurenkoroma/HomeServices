package main

import (
	"fmt"
	"net/http"

	"samurenkoroma/services/configs"
	"samurenkoroma/services/internal/link"
	"samurenkoroma/services/internal/stat"
	"samurenkoroma/services/pkg/db"
	"samurenkoroma/services/pkg/event"
	"samurenkoroma/services/pkg/middleware"
)

func main() {
	conf := configs.LoadConfig()
	database := db.NewDb(conf)
	router := http.NewServeMux()

	eventBus := event.NewEventBus()

	//Репозитории
	linkRepo := link.NewLinkRepository(database)
	// userRepo := user.NewUserRepo(database)
	statRepo := stat.NewStatRepo(database)

	//Сервисы
	// authService := auth.NewAuthService(userRepo)
	statService := stat.NewStatService(stat.StatServiceDeps{EventBus: eventBus, StatRepository: statRepo})

	//Обработчики
	// auth.NewAuthHandler(router, auth.AuthHandlerDeps{
	// 	Config:      conf,
	// 	AuthService: authService,
	// })
	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepo,
		EventBus:       eventBus,
		Config:         conf,
	})
	stat.NewStatHandler(router, stat.StatHandlerDeps{
		StatRepository: statRepo,
	})
	stack := middleware.Chain(
		middleware.Logging,
		middleware.CORS,
	)

	server := http.Server{
		Addr:    ":8081",
		Handler: stack(router),
	}

	go statService.AddClick()

	fmt.Println("Сервер запущен на порту :8081")
	server.ListenAndServe()
}
