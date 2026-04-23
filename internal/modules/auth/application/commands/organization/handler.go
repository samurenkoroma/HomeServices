package organization

import (
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/modules/auth/infrastructure/jwt"
)

type OrganizationHandler struct {
	uowFactory repository.Factory
	jwtService *jwt.Service
}

func NewOrganizationHandler(uowFactory repository.Factory, jwtService *jwt.Service) *OrganizationHandler {
	return &OrganizationHandler{uowFactory: uowFactory, jwtService: jwtService}
}
