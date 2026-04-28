package organization

import (
	"context"
	"errors"
	"fmt"
	"samurenkoroma/services/internal/application/command"
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/modules/auth/application/dto"
	"samurenkoroma/services/internal/modules/auth/domain"
	"samurenkoroma/services/internal/modules/auth/infrastructure/jwt"
	"samurenkoroma/services/internal/modules/auth/infrastructure/persistence/postgres"
)

type SwitchOrganizationCmd struct {
	OrganizationID string `json:"organization_id"`
}

type switchOrganizationResult struct {
	TokenPair  *jwt.TokenPair           `json:"tokenPair"`
	CurrentOrg dto.UserOrganizationInfo `json:"currentOrg"`
}

func (h *OrganizationHandler) Switch(ctx context.Context, cmd any) (any, error) {
	c, ok := cmd.(*SwitchOrganizationCmd)
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
		user, err := userRepo.FindByID(ctx, userID)
		if err != nil {
			return nil, err
		}

		// Проверяем членство в организации
		membership, err := membershipRepo.FindByUserAndOrganization(ctx, userID, c.OrganizationID)
		if err != nil {
			return nil, errors.New("you don't have access to this organization")
		}

		if !membership.IsActive {
			return nil, errors.New("membership is not active")
		}

		// Получаем информацию об организации
		org, err := orgRepo.FindByID(ctx, c.OrganizationID)
		if err != nil {
			return nil, err
		}

		// Обновляем текущую организацию в профиле
		user.SetCurrentOrganization(org.ID)
		if err := userRepo.Update(ctx, user); err != nil {
			return nil, err
		}
		uow.RegisterAggregate(user)
		// Генерируем новые токены с новой организацией
		tokenPair, err := h.jwtService.GenerateTokenPair(
			user.ID,
			user.Username,
			user.Email,
			string(user.Role),
			c.OrganizationID,
			string(membership.Role),
		)
		if err != nil {
			return nil, err
		}

		return switchOrganizationResult{
			TokenPair: tokenPair,
			CurrentOrg: dto.UserOrganizationInfo{
				OrganizationID:   org.ID,
				OrganizationName: org.Name,
				Role:             string(membership.Role),
				RoleName:         membership.GetRoleName(),
			},
		}, nil
	})

}
