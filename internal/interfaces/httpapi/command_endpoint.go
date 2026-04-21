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

		handlerCmd, err := router.ResolveCommandPayload(payload.Command, payload.Data)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(ErrResponse{
				Error: err.Error(),
			})
			return
		}
		if err := router.Dispatch(r.Context(), payload.Command, handlerCmd); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(ErrResponse{
				Error: err.Error(),
			})
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

type ErrResponse struct {
	Error string `json:"error"`
}
