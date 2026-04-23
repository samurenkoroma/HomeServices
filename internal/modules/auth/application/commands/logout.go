package commands

import (
	"net/http"
	"samurenkoroma/services/pkg/response"
)

type LogoutRequest struct {
	RefreshToken string `json:"refresh_token"`
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	// В простой реализации просто возвращаем успех
	// Для полноценной реализации нужно добавить blacklist токенов в Redis
	response.Success(map[string]string{
		"message": "logout successful",
	}).WriteJSON(w, http.StatusOK)
}
