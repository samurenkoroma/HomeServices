package httpapi

import (
	"net/http"
	"samurenkoroma/services/internal/application/command"
	"samurenkoroma/services/internal/application/query"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func NewRouter(
	commandRouter command.Router,
	queryRouter query.Router,
) http.Handler {

	mux := chi.NewMux()
	mux.Use(cors.Handler(cors.Options{
		//AllowedOrigins: []string{"https://localhost"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "Access-Control-Allow-Origin"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	commandEndpoint := CommandEndpoint(commandRouter)
	queryEndpoint := QueryEndpoint(queryRouter)

	mux.Handle("/commands/", commandEndpoint)
	mux.Handle("/queries/", queryEndpoint)

	return withMiddleware(mux)
}
