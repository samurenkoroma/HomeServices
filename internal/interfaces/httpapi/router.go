package httpapi

import (
	"net/http"
	"samurenkoroma/services/internal/application/command"
	"samurenkoroma/services/internal/application/query"
)

func NewRouter(
	commandRouter command.Router,
	queryRouter query.Router,
) http.Handler {

	mux := http.NewServeMux()

	commandEndpoint := CommandEndpoint(commandRouter)
	queryEndpoint := QueryEndpoint(queryRouter)

	mux.Handle("/commands/", commandEndpoint)
	mux.Handle("/queries/", queryEndpoint)

	return withMiddleware(mux)
}
