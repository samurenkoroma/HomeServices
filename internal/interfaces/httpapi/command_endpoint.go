package httpapi

import (
	"encoding/json"
	"net/http"
	"samurenkoroma/services/internal/application/command"
	"samurenkoroma/services/internal/application/command/dto"
)

func CommandEndpoint(router command.Router) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var payload dto.CommandPayload

		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		handlerCmd, _ := router.ResolveCommandPayload(payload.Command, payload.Data)

		err := router.Dispatch(r.Context(),
			payload.Command,
			handlerCmd,
		)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
