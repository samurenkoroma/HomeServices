package auth

// // RefreshRequest запрос на обновление токена
//
//	type RefreshRequest struct {
//		RefreshToken string `json:"refresh_token"`
//	}
//
// // RefreshResponse ответ с новыми токенами
//
//	type RefreshResponse struct {
//		TokenPair *jwt.TokenPair `json:"token_pair"`
//	}
//
// // POST /auth/refresh
//
//	func (h *AuthHandler) Refresh(w http.ResponseWriter, r *http.Request) {
//		var req RefreshRequest
//		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
//			writeError(w, http.StatusBadRequest, "invalid request body")
//			return
//		}
//
//		if req.RefreshToken == "" {
//			writeError(w, http.StatusBadRequest, "refresh_token is required")
//			return
//		}
//
//		tokenPair, err := h.jwtService.RefreshToken(req.RefreshToken)
//		if err != nil {
//			writeError(w, http.StatusUnauthorized, "invalid refresh token")
//			return
//		}
//
//		writeJSON(w, http.StatusOK, RefreshResponse{
//			TokenPair: tokenPair,
//		})
//	}
//
// LogoutRequest запрос на выход
