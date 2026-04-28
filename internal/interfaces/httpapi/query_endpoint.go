package httpapi

import (
	"encoding/json"
	"net/http"
	"samurenkoroma/services/internal/application/query"
	"samurenkoroma/services/internal/application/query/dto"
	"samurenkoroma/services/pkg/response"
)

func QueryEndpoint(router query.Router) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var payload dto.QueryPayload

		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			http.Error(w, "invalid request payload", http.StatusBadRequest)
			return
		}

		result, err := router.Dispatch(
			r.Context(),
			payload.Query,
			payload.Data,
		)
		if err != nil {
			resp := response.FromError(err)
			statusCode := getStatusCodeForError(resp.Error.Code)
			resp.WriteJSON(w, statusCode)
			return
		}

		response.WriteSuccess(w, result)
	}
}
