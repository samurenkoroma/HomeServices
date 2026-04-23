package commands

import (
	"encoding/json"
	"fmt"
	"net/http"
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/modules/auth/application/dto"
	"samurenkoroma/services/internal/modules/auth/domain"
	"samurenkoroma/services/internal/modules/auth/infrastructure/jwt"
	"samurenkoroma/services/internal/modules/auth/infrastructure/persistence/postgres"
	"samurenkoroma/services/pkg/response"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type User struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}
type LoginResult struct {
	TokenPair     *jwt.TokenPair             `json:"tokenPair"`
	User          User                       `json:"user"`
	Organizations []dto.UserOrganizationInfo `json:"organizations"`
	CurrentOrg    *dto.UserOrganizationInfo  `json:"currentOrg,omitempty"`
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteValidationError(w, "invalid request body")
		return
	}

	if req.Email == "" {
		response.WriteValidationError(w, "email is required")
		return
	}
	if req.Password == "" || len(req.Password) < 6 {
		response.WriteValidationError(w, "password must be at least 6 characters")
		return
	}

	uow, err := h.uowFactory.Begin(r.Context())
	if err != nil {
		response.WriteInternalError(w, "failed to start transaction")
		return
	}
	ctx := r.Context()
	err = uow.Execute(ctx, postgres.NewPostgresAuthProvider, func(provider repository.RepositoryProvider) error {
		// Приводим провайдер к нужному типу
		authProvider, ok := provider.(*postgres.PostgresAuthProvider)
		if !ok {
			response.WriteInternalError(w, err.Error())
			return fmt.Errorf("expected FarmProvider, got %T", provider)
		}

		userRepo := authProvider.Users()
		membershipRepo := authProvider.Memberships()
		orgRepo := authProvider.Organizations()

		// Ищем пользователя
		user, err := userRepo.FindByEmail(ctx, req.Email)
		if err != nil {
			response.WriteNotFound(w, err.Error())
			return domain.ErrInvalidCredentials
		}

		// Проверяем пароль
		if !user.CheckPassword(req.Password) {
			response.WriteValidationError(w, domain.ErrInvalidCredentials.Error())
			return domain.ErrInvalidCredentials
		}

		// Проверяем статус
		if !user.IsActive() {
			return domain.ErrUserInactive
		}

		// Получаем все членства пользователя
		memberships, err := membershipRepo.FindByUser(ctx, user.ID)
		if err != nil {
			return err
		}

		// Собираем информацию об организациях
		var organizations []dto.UserOrganizationInfo
		for _, m := range memberships {
			if !m.IsActive {
				continue
			}
			org, err := orgRepo.FindByID(ctx, m.OrganizationID)
			if err != nil {
				continue
			}
			organizations = append(organizations, dto.UserOrganizationInfo{
				OrganizationID:   org.ID,
				OrganizationName: org.Name,
				Role:             string(m.Role),
				RoleName:         m.GetRoleName(),
			})
		}

		// Определяем текущую организацию
		var currentOrg *dto.UserOrganizationInfo
		currentOrgID := user.GetCurrentOrganizationID()

		if currentOrgID != "" {
			for _, org := range organizations {
				if org.OrganizationID == currentOrgID {
					currentOrg = &org
					break
				}
			}
		}

		// Если нет текущей организации, но есть организации - выбираем первую
		if currentOrg == nil && len(organizations) > 0 {
			currentOrg = &organizations[0]
			// Сохраняем выбранную организацию в профиле пользователя
			user.SetCurrentOrganization(currentOrg.OrganizationID)
			userRepo.Update(ctx, user)
		}

		// Генерируем токены с текущей организацией
		var tokenPair *jwt.TokenPair
		if currentOrg != nil {
			tokenPair, err = h.jwtService.GenerateTokenPair(
				user.ID,
				user.Username,
				user.Email,
				string(user.Role),
				currentOrg.OrganizationID,
				currentOrg.Role,
			)
		} else {
			// Если нет организаций, токен без organization_id
			tokenPair, err = h.jwtService.GenerateTokenPair(
				user.ID,
				user.Username,
				user.Email,
				string(user.Role),
				"",
				"",
			)
		}
		if err != nil {
			return err
		}

		// Обновляем время последнего входа
		user.UpdateLastLogin()
		userRepo.Update(ctx, user)

		uow.RegisterAggregate(user)
		response.Success(
			LoginResult{
				TokenPair: tokenPair,
				User: User{
					Id:    user.ID,
					Name:  user.Username,
					Email: user.Email,
					Role:  user.Role.String(),
				},
				CurrentOrg: currentOrg,
			}).WriteJSON(w, http.StatusOK)

		return nil
	})
	return
}
