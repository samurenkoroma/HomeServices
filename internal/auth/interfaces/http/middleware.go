package http

import (
	"net/http"
	"strings"

	"samurenkoroma/services/internal/auth/domain"
	"samurenkoroma/services/internal/auth/infrastructure/jwt"
	authctx "samurenkoroma/services/internal/common/auth"
)

// AuthMiddleware создает middleware для проверки JWT
type AuthMiddleware struct {
	jwtService *jwt.Service
}

// NewAuthMiddleware создает новый AuthMiddleware
func NewAuthMiddleware(jwtService *jwt.Service) *AuthMiddleware {
	return &AuthMiddleware{jwtService: jwtService}
}

// Authenticate проверяет наличие и валидность токена
func (m *AuthMiddleware) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Получаем токен из заголовка Authorization
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			writeError(w, http.StatusUnauthorized, "missing authorization header")
			return
		}

		// Извлекаем токен (Bearer <token>)
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			writeError(w, http.StatusUnauthorized, "invalid authorization format")
			return
		}

		token := parts[1]

		// Валидируем токен
		claims, err := m.jwtService.ValidateToken(token)
		if err != nil {
			writeError(w, http.StatusUnauthorized, err.Error())
			return
		}

		// Сохраняем информацию о пользователе в контексте
		ctx := authctx.WithUser(r.Context(), &authctx.UserInfo{
			ID:       claims.UserID,
			Username: claims.Username,
			Email:    claims.Email,
			Role:     domain.Role(claims.Role),
		})

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// RequirePermission проверяет наличие права доступа
func (m *AuthMiddleware) RequirePermission(perm domain.Permission) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user := authctx.GetUser(r.Context())
			if user == nil {
				writeError(w, http.StatusUnauthorized, "unauthorized")
				return
			}

			if !user.Role.HasPermission(perm) {
				writeError(w, http.StatusForbidden, "forbidden: insufficient permissions")
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

// RequireRole проверяет наличие роли
func (m *AuthMiddleware) RequireRole(roles ...domain.Role) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user := authctx.GetUser(r.Context())
			if user == nil {
				writeError(w, http.StatusUnauthorized, "unauthorized")
				return
			}

			for _, role := range roles {
				if user.Role == role || user.Role == domain.RoleAdmin {
					next.ServeHTTP(w, r)
					return
				}
			}

			writeError(w, http.StatusForbidden, "forbidden: insufficient role")
		})
	}
}

// OptionalAuth опциональная аутентификация (не требует токена)
func (m *AuthMiddleware) OptionalAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader != "" {
			parts := strings.Split(authHeader, " ")
			if len(parts) == 2 && parts[0] == "Bearer" {
				claims, err := m.jwtService.ValidateToken(parts[1])
				if err == nil {
					ctx := authctx.WithUser(r.Context(), &authctx.UserInfo{
						ID:       claims.UserID,
						Username: claims.Username,
						Email:    claims.Email,
						Role:     domain.Role(claims.Role),
					})
					r = r.WithContext(ctx)
				}
			}
		}
		next.ServeHTTP(w, r)
	})
}

func writeError(w http.ResponseWriter, status int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(`{"error":"` + msg + `"}`))
}
