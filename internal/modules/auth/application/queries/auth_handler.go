package queries

import (
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/modules/auth/infrastructure/jwt"
)

type UserHandler struct {
	uowFactory repository.Factory
	jwtService *jwt.Service
}

func NewUserHandler(uowFactory repository.Factory, jwtService *jwt.Service) *UserHandler {
	return &UserHandler{
		uowFactory: uowFactory,
		jwtService: jwtService,
	}
}
