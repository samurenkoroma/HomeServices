package commands

import (
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/modules/auth/infrastructure/jwt"
)

// AuthHandler обрабатывает auth запросы (без CQRS)
type AuthHandler struct {
	uowFactory repository.Factory
	jwtService *jwt.Service
}

// NewAuthHandler создает новый AuthHandler
func NewAuthHandler(uowFactory repository.Factory, jwtService *jwt.Service) *AuthHandler {
	return &AuthHandler{
		uowFactory: uowFactory,
		jwtService: jwtService,
	}
}
