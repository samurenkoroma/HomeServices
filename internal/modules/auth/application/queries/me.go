package queries

import (
	"context"
	"fmt"
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/modules/auth/application/commands"
	"samurenkoroma/services/internal/modules/auth/application/dto"
	"samurenkoroma/services/internal/modules/auth/domain"
	"samurenkoroma/services/internal/modules/auth/infrastructure/persistence/postgres"
	"samurenkoroma/services/pkg/response"
)

type MeResponseQuery struct {
}

// MeResponse ответ с информацией о текущем пользователе
type MeResponse struct {
	User         commands.User               `json:"user"`
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
	var data MeResponse
	err = uow.Execute(ctx, postgres.NewPostgresAuthProvider, func(provider repository.RepositoryProvider) error {
		authProvider, ok := provider.(*postgres.PostgresAuthProvider)
		if !ok {
			return fmt.Errorf("expected FarmProvider, got %T", provider)
		}

		userRepo := authProvider.Users()
		orgRepo := authProvider.Organizations()

		// Получаем пользователя
		user, err := userRepo.FindByID(ctx, userID)
		if err != nil {
			return err
		}

		orgs, err := orgRepo.ListByUser(ctx, user.ID)
		if err != nil {
			return err
		}
		var organizations []*dto.UserOrganizationInfo
		for _, o := range orgs {
			organizations = append(organizations, &dto.UserOrganizationInfo{
				OrganizationID:   o.ID,
				OrganizationName: o.Name,
				Role:             "Владелец",
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

		data = MeResponse{
			User: commands.User{
				Id:    user.ID,
				Name:  user.Username,
				Email: user.Email,
				Role:  user.Role.String(),
			},
			Organization: organizations,
			CurrentOrg:   currentOrg,
		}
		return nil

	})

	return response.Success(data), nil
}
