package queries

import (
	"context"
	"fmt"
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/modules/auth/application/commands/auth"
	"samurenkoroma/services/internal/modules/auth/application/dto"
	"samurenkoroma/services/internal/modules/auth/domain"
	"samurenkoroma/services/internal/modules/auth/infrastructure/persistence/postgres"
)

type MeQuery struct {
}

// MeResponse ответ с информацией о текущем пользователе
type MeResponse struct {
	User         auth.User                   `json:"user"`
	Organization []*dto.UserOrganizationInfo `json:"organizations"`
	CurrentOrg   *dto.UserOrganizationInfo   `json:"currentOrg"`
}

func (h *UserHandler) Handle(ctx context.Context, cmd any) (any, error) {
	userID, ok := ctx.Value("user_id").(string)
	if !ok {
		return nil, domain.ErrUnauthorized
	}

	uow, err := h.uowFactory.Begin(ctx)
	if err != nil {
		return nil, err
	}
	return uow.Execute(ctx, postgres.NewPostgresAuthProvider, func(provider repository.RepositoryProvider) (any, error) {
		authProvider, ok := provider.(*postgres.PostgresAuthProvider)
		if !ok {
			return nil, fmt.Errorf("expected FarmProvider, got %T", provider)
		}

		userRepo := authProvider.Users()
		membershipRepo := authProvider.Memberships()
		orgRepo := authProvider.Organizations()

		// Получаем пользователя
		user, err := userRepo.FindByID(ctx, userID)
		if err != nil {
			return nil, err
		}

		orgs, err := orgRepo.ListByUser(ctx, user.ID)
		if err != nil {
			return nil, err
		}
		var organizations []*dto.UserOrganizationInfo
		for _, o := range orgs {
			member, err2 := membershipRepo.FindByUserAndOrganization(ctx, user.ID, o.ID)
			if err2 != nil {
				return nil, err2
			}
			organizations = append(organizations, &dto.UserOrganizationInfo{
				OrganizationID:   o.ID,
				OrganizationName: o.Name,
				Role:             member.GetRoleName(),
			})
		}
		// Определяем текущую организацию
		var currentOrg *dto.UserOrganizationInfo
		currentOrgID := user.GetCurrentOrganizationID()

		if currentOrgID != "" {
			for _, org := range organizations {
				if org.OrganizationID == currentOrgID {
					currentOrg = org
					break
				}
			}
		}

		return MeResponse{
			User: auth.User{
				Id:    user.ID,
				Name:  user.Username,
				Email: user.Email,
				Role:  user.Role.String(),
			},
			Organization: organizations,
			CurrentOrg:   currentOrg,
		}, nil

	})
}
