package httpapi

import (
	"encoding/json"
	"net/http"
	"samurenkoroma/services/internal/application/query"
	"samurenkoroma/services/internal/application/query/dto"
)

func QueryEndpoint(router query.QueryRouter) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var payload dto.QueryPayload

		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			http.Error(w, "invalid request payload", http.StatusBadRequest)
			return
		}

		query, err := router.Decode(payload.Query, payload.Data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		result, err := router.Dispatch(
			r.Context(),
			payload.Query,
			query,
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	}
}
