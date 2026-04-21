package auth

import (
	"context"

	"samurenkoroma/services/internal/auth/domain"
)

type contextKey string

const userInfoKey contextKey = "user_info"

// UserInfo информация о пользователе в контексте
type UserInfo struct {
	ID       string
	Username string
	Email    string
	Role     domain.Role
}

// WithUser добавляет пользователя в контекст
func WithUser(ctx context.Context, user *UserInfo) context.Context {
	return context.WithValue(ctx, userInfoKey, user)
}

// GetUser извлекает пользователя из контекста
func GetUser(ctx context.Context) *UserInfo {
	user, ok := ctx.Value(userInfoKey).(*UserInfo)
	if !ok {
		return nil
	}
	return user
}

// GetUserID извлекает ID пользователя из контекста
func GetUserID(ctx context.Context) string {
	user := GetUser(ctx)
	if user == nil {
		return ""
	}
	return user.ID
}

// HasRole проверяет наличие роли у пользователя в контексте
func HasRole(ctx context.Context, role domain.Role) bool {
	user := GetUser(ctx)
	if user == nil {
		return false
	}
	return user.Role == role || user.Role == domain.RoleAdmin
}

// HasPermission проверяет наличие права у пользователя в контексте
func HasPermission(ctx context.Context, perm domain.Permission) bool {
	user := GetUser(ctx)
	if user == nil {
		return false
	}
	return user.Role.HasPermission(perm)
}
