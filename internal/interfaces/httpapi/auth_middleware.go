package httpapi

import (
	"context"
	"net/http"
	"samurenkoroma/services/internal/modules/auth/domain"
	"samurenkoroma/services/internal/modules/auth/infrastructure/jwt"
	"samurenkoroma/services/pkg/response"
	"strings"
)

type AuthMiddleware struct {
	jwtService *jwt.Service
}

func NewAuthMiddleware(jwtService *jwt.Service) *AuthMiddleware {
	return &AuthMiddleware{jwtService: jwtService}
}

// Authenticate проверяет токен и добавляет все данные в контекст
func (m *AuthMiddleware) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			response.WriteUnauthorized(w, "missing authorization header")
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			response.WriteUnauthorized(w, "invalid authorization format")
			return
		}

		token := parts[1]

		claims, err := m.jwtService.ValidateToken(token)
		if err != nil {
			response.WriteUnauthorized(w, err.Error())
			return
		}

		// Добавляем ВСЕ данные из токена в контекст
		ctx := r.Context()
		ctx = context.WithValue(ctx, "user_id", claims.UserID)
		ctx = context.WithValue(ctx, "username", claims.Username)
		ctx = context.WithValue(ctx, "email", claims.Email)
		ctx = context.WithValue(ctx, "organization_id", claims.OrganizationID)
		ctx = context.WithValue(ctx, "org_role", domain.OrganizationRole(claims.OrgRole))
		ctx = context.WithValue(ctx, "role", domain.Role(claims.Role))

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetOrganizationID извлекает ID текущей организации из контекста
func GetOrganizationID(ctx context.Context) string {
	if orgID, ok := ctx.Value("organization_id").(string); ok {
		return orgID
	}
	return ""
}

// GetOrgRole извлекает роль в текущей организации из контекста
func GetOrgRole(ctx context.Context) domain.OrganizationRole {
	if role, ok := ctx.Value("org_role").(domain.OrganizationRole); ok {
		return role
	}
	return ""
}
