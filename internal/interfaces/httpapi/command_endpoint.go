package httpapi

import (
	"encoding/json"
	"net/http"
	"samurenkoroma/services/internal/application/command"
	"samurenkoroma/services/internal/application/command/dto"
	"samurenkoroma/services/pkg/response"
)

func CommandEndpoint(router command.Router) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var payload dto.CommandPayload

		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			response.WriteValidationError(w, "invalid request payload: "+err.Error())
			return
		}
		if payload.Command == "" {
			response.WriteValidationError(w, "command name is required")
			return
		}

		handlerCmd, err := router.ResolveCommandPayload(payload.Command, payload.Data)
		if err != nil {
			response.WriteError(w, http.StatusBadRequest, response.CodeBadRequest,
				"failed to decode command: "+err.Error())
			return
		}

		result, err := router.Dispatch(r.Context(), payload.Command, handlerCmd)
		if err != nil {
			// Используем стандартный ответ с ошибкой
			resp := response.FromError(err)
			statusCode := getStatusCodeForError(resp.Error.Code)
			resp.WriteJSON(w, statusCode)
			return
		}

		// Если команда вернула результат (например, LoginResponse)
		if result != nil {
			response.WriteSuccess(w, result)
		} else {
			// Успешное выполнение без данных
			response.WriteSuccess(w, map[string]string{"status": "ok"})
		}
	}
}

// getStatusCodeForError возвращает HTTP статус код по коду ошибки
func getStatusCodeForError(errorCode string) int {
	switch errorCode {
	case response.CodeBadRequest, response.CodeValidation:
		return http.StatusBadRequest
	case response.CodeUnauthorized:
		return http.StatusUnauthorized
	case response.CodeForbidden:
		return http.StatusForbidden
	case response.CodeNotFound:
		return http.StatusNotFound
	case response.CodeConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
