package organization

import (
	"context"
	"fmt"
	"samurenkoroma/services/internal/application/command"
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/modules/auth/application/dto"
	"samurenkoroma/services/internal/modules/auth/domain"
	"samurenkoroma/services/internal/modules/auth/infrastructure/persistence/postgres"
)

type CreateOrganizationCmd struct {
	Name string `json:"name"`
}

func (h *OrganizationHandler) Create(ctx context.Context, cmd any) (any, error) {
	c, ok := cmd.(*CreateOrganizationCmd)
	if !ok {
		return nil, command.ErrInvalidCommandType
	}

	uow, err := h.uowFactory.Begin(ctx)
	if err != nil {
		return nil, err
	}

	// Получаем текущего пользователя из контекста
	userID, ok := ctx.Value("user_id").(string)
	if !ok {
		return nil, domain.ErrUnauthorized
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
		_, err := userRepo.FindByID(ctx, userID)
		if err != nil {
			return nil, err
		}

		newOrg, err := domain.NewSimpleOrganization(c.Name)
		if err != nil {
			return nil, err
		}

		if err := orgRepo.Save(ctx, newOrg); err != nil {
			return nil, err
		}

		membership, err := domain.NewMembership(userID, newOrg.ID, domain.OrgRoleOwner)
		if err != nil {
			return nil, err
		}

		if err := membershipRepo.Save(ctx, membership); err != nil {
			return nil, err
		}

		return dto.UserOrganizationInfo{
			OrganizationID:   newOrg.ID,
			OrganizationName: newOrg.Name,
			Role:             string(membership.Role),
			RoleName:         membership.GetRoleName(),
		}, nil
	})
}
